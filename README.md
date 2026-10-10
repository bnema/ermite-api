# ermite-api

The public API contract of [ermite](https://git.bnema.dev/bnema/ermite), a single-user Git forge: protobuf definitions for its ConnectRPC services and the Go code generated from them.

Status: early development, no services defined yet.

```go
import "bnema.dev/ermite-api/gen/ermite"
```

The contract has no version number. It only changes in backward-compatible ways (new fields, messages, RPCs and services), and CI rejects any breaking change.

## Layout

- `proto/ermite/`: `.proto` files of the `ermite` package, checked by `buf lint` and `buf breaking`.
- `gen/`: generated Go code, committed. Never edit it by hand.

## Generate

```sh
buf generate
```

## License

MIT. See [LICENSE](LICENSE).
