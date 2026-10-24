<h1 align="center">ermite-api</h1>

<p align="center">
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue?style=flat-square" alt="License: MIT"></a>
  <a href="https://github.com/bnema/ermite-api/commits/main"><img src="https://badgen.net/github/last-commit/bnema/ermite-api/main?icon=github" alt="Last commit"></a>
  <a href="https://github.com/bnema/ermite-api/stargazers"><img src="https://badgen.net/github/stars/bnema/ermite-api?icon=github" alt="GitHub stars"></a>
</p>

<p align="center">The public API of <a href="https://github.com/bnema/ermite">ermite</a>: protobuf definitions and generated Go code for its ConnectRPC services.</p>

> Read-only mirror. Development happens at git.bnema.dev; issues and pull requests are closed here.

---

```go
import (
	"bnema.dev/ermite-api/gen/ermite"
	"bnema.dev/ermite-api/gen/ermite/ermiteconnect"
)
```

- The server mounts the services under `/-/api/`. Errors are Connect codes.
- No version number: the contract only changes in backward-compatible ways, and CI rejects breaking changes (`buf breaking`).
- Tokens carry scopes only; presets such as `cli` or `agent` are a client concern.

Status: early development.

## Services

| Service | Purpose |
|---|---|
| `DepsService` | dependency updates (owner only) |
| `RepoService` | repositories and their settings |
| `GitService` | refs, trees, blobs, commits and diffs |
| `TokenService` | API tokens |
| `MergeRequestService` | merge requests, comments, merge |
| `IssueService` | issues, comments, owner moderation of public submissions |
| `SessionService` | web sign-in links (owner only) |
| `EventService` | live event stream, resumable |
| `ServerService` | version, features, signing key, OIDC key rotation |
| `RunService` | CI runs, jobs, logs, checks, runners |
| `RunnerService` | used by CI runners to fetch and report jobs |
| `SecretService`, `VariableService`, `EnvironmentService` | CI settings |
| `ArtifactService` | CI artifacts, for people |
| `JobArtifactService` | CI artifacts v4, for runners |
| `TrustedKeyService` | SSH and OpenPGP keys commit signatures are checked against |

Every RPC: [docs/services.md](docs/services.md).

## Layout

- `proto/ermite/`: one `.proto` per service, plus shared `types.proto`.
- `gen/ermite/`: generated messages; `gen/ermite/ermiteconnect/`: Connect clients and handlers. Committed, never edited by hand.

```sh
buf generate
```

## License

MIT. See [LICENSE](LICENSE).
