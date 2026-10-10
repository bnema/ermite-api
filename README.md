# ermite-api

The public API contract of [ermite](https://git.bnema.dev/bnema/ermite), a single-user Git forge: protobuf definitions for its ConnectRPC services and the Go code generated from them.

Status: early development. The contract defines five services (below); the server that implements them is being built.

```go
import "bnema.dev/ermite-api/gen/ermite"
```

The contract has no version number. It only changes in backward-compatible ways (new fields, messages, RPCs and services), and CI rejects any breaking change.

## Services

Handlers are mounted under `/-/api/` by the server. Errors are Connect codes.

| Service | RPCs |
|---|---|
| `RepoService` | `ListRepos`, `GetRepo`, `CreateRepo`, `UpdateRepo` (field mask), `DeleteRepo` |
| `GitService` | `ListRefs`, `GetTree`, `GetBlob` (streamed), `ListCommits`, `GetCommit` (structured diff), `ResolveRevision` |
| `TokenService` | `CreateToken` (secret shown once), `ListTokens`, `RevokeToken` |
| `EventService` | `Watch` (streamed, resumable) |
| `ServerService` | `GetVersion` (server version and feature list) |

Tokens carry scopes only; presets such as `cli` or `agent` are a client concern.

## Layout

- `proto/ermite/`: `.proto` files of the `ermite` package, one per service (`repo`, `git`, `token`, `event`, `server`) plus shared `types`. Checked by `buf lint` and `buf breaking`.
- `gen/ermite/`: generated Go messages. `gen/ermite/ermiteconnect/`: generated Connect clients and handlers. Committed, never edited by hand.

## Generate

```sh
buf generate
```

## License

MIT. See [LICENSE](LICENSE).
