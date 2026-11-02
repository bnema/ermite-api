# Connect Go v2 migration

The generated Go API uses `connectrpc.com/connect/v2` v2.0.2. Protobuf messages,
service names, procedure constants and wire protocols are unchanged. Existing
wire clients need no migration. Go consumers (including ermite and ermite-runner)
must migrate their Connect imports and call sites together with this module.

This migration follows the upstream [v2 migration guide](https://github.com/connectrpc/connect-go/blob/v2.0.2/docs/v2-migration.md)
and [v2 guide](https://github.com/connectrpc/connect-go/blob/v2.0.2/docs/v2-guide.md).

## Clients

All generated `New<Service>Client` constructors now take a single
`*connect.Client`, not an HTTP client, URL and client options:

```go
client := connect.NewClient(
    connecthttp.NewTransport(httpClient, baseURL+"/-/api"),
)
runner := ermiteconnect.NewRunnerServiceClient(client)
response, err := runner.FetchJob(ctx, &ermite.FetchJobRequest{})
```

Import `connectrpc.com/connect/v2` and
`connectrpc.com/connect/v2/connecthttp`. Multiple service clients can share the
core client. HTTP options such as `WithGRPC`, compression and size limits move to
`connecthttp.NewTransport`; interceptors go to `connect.NewClient`.

Unary methods now take `*ermite.<Request>` and return `(*ermite.<Response>, error)`.
Remove `connect.NewRequest`, `connect.NewResponse`, and `.Msg` access. Headers and
trailers move from request/response wrappers to `connect.NewClientContext` and
its `CallInfo`; handlers use `connect.CallInfoForServerContext(ctx)`.

## Servers

`New<Service>Handler(svc, opts...) (string, http.Handler)` is removed. Use
`Register<Service>Handler(server, svc)` (no return values) with a
`server := connect.NewServer(...)`, then `connecthttp.Mount(mux, server, ...)`.
For the forge prefix, mount on an inner mux and route `/-/api/` through
`http.StripPrefix("/-/api", innerMux)`. See `TestConnectV2HTTP` for a working
client/server example. HTTP options go to `Mount`, interceptors to `NewServer`.

The `<Service>Handler` and `Unimplemented<Service>Handler` names remain, but
unary implementation signatures use bare messages as described above.

## Streaming

Generic stream types are replaced by generated per-RPC value types named
`<Service><Method>ClientStream` and `<Service><Method>ServerStream`:

- Server streaming: `EventServiceWatch`, `GitServiceGetBlob`,
  `RunServiceGetJobLog`, `RunnerServiceWatchJob`,
  `ReleaseServiceDownloadReleaseAsset`, `JobArtifactServiceDownloadJobArtifact`,
  `ArtifactServiceDownloadArtifact`.
- Client streaming: `ReleaseServiceUploadReleaseAsset`,
  `JobArtifactServiceUploadJobArtifactBlock`.

For example, `WatchJob(ctx, *ermite.WatchJobRequest)` returns
`(RunnerServiceWatchJobClientStream, error)`; its handler takes
`(context.Context, *ermite.WatchJobRequest, RunnerServiceWatchJobServerStream)`
and returns `error`. Streams are values, not pointers to generic Connect types.
Client `Receive()` returns `(*Message, error)` instead of a boolean followed by
`Msg()`/`Err()`; handle `errors.Is(err, io.EOF)` and defer `Close()` when abandoning
a server stream. Server `Send(*Message)` remains an error-returning operation.
Client-streaming constructors now return `(stream, error)`, so check the error
before `Send`/`CloseAndReceive`. Their handlers receive a generated server-stream
value; `Receive()` returns `(*Message, error)` and completion is `io.EOF`; the
handler returns the bare response message.

## Errors and middleware

The repository's `JobEnded` detail test now uses
`connectproto.NewErrorDetail`, `connect.NewError(code, "message").WithDetail(d)`
and `connectproto.UnmarshalErrorDetail`. `NewError` accepts a string rather than
an error; `WithDetail` returns a clone rather than mutating with `AddDetail`.
Downstream error detail producers and readers must make the same changes.

Downstream middleware also needs the v2 API: the old interceptor interface and
`WithInterceptors` options are replaced by `ClientInterceptor`/`ServerInterceptor`
functions around streams. Metadata is on context-carried `CallInfo`, not wrappers;
HTTP-specific information is exposed through `connecthttp`. These facilities are
not implemented by this schema-only repository.

## Generation

The tool module path is `connectrpc.com/connect/v2/cmd/protoc-gen-connect-go`;
the binary name remains `protoc-gen-connect-go`. `buf.gen.yaml` needs no change:
`go tool protoc-gen-connect-go` and `paths=source_relative` still work. v2 always
generates bare-message signatures (the old optional `simple` flag is rejected;
this repository did not use it). No proto or generated message package paths
were changed.
