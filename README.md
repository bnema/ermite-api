# ermite-api

The public API contract of [ermite](https://git.bnema.dev/bnema/ermite), a single-user Git forge: protobuf definitions for its ConnectRPC services and the Go code generated from them.

Status: early development. The contract defines fourteen services (below); the server that implements them is being built.

```go
import "bnema.dev/ermite-api/gen/ermite"
```

The contract has no version number. It only changes in backward-compatible ways (new fields, messages, RPCs and services), and CI rejects any breaking change.

## Services

Handlers are mounted under `/-/api/` by the server. Errors are Connect codes.

| Service | RPCs |
|---|---|
| `RepoService` | `ListRepos`, `GetRepo`, `CreateRepo`, `UpdateRepo` (field mask, merge settings), `DeleteRepo` |
| `GitService` | `ListRefs`, `GetTree`, `ListPaths`, `GetBlob` (streamed), `ListCommits`, `GetCommit` (structured diff), `ResolveRevision` |
| `TokenService` | `CreateToken` (secret shown once), `ListTokens`, `RevokeToken` |
| `MergeRequestService` | `CreateMergeRequest`, `ListMergeRequests`, `GetMergeRequest` (commits, diff, comments), `AddComment`, `ResolveComment`, `CloseMergeRequest`, `MergeMergeRequest`, `MarkRead` |
| `SessionService` | `CreateLoginLink`, `SignOutEverywhere` (owner only) |
| `EventService` | `Watch` (streamed, resumable) |
| `RunnerService` | `RegisterRunner`, `FetchJob` (long poll), `UpdateJob` (steps, annotations), `AppendLog` (offset-idempotent), `Heartbeat`, `WatchJob` (cancel requests, streamed), `CompleteJob` (idempotent), `ResolveActions` (where to download the `uses:` actions), `UploadSummary` (job summary of a step); for CI runners |
| `RunService` | `ListRuns`, `GetRun`, `GetJob`, `GetJobLog` (streamed, follow), `ListAnnotations`, `GetStepSummaries`, `CancelRun`, `RerunRun`, `DispatchWorkflow`, `ListChecks`, `ListRunners`, `RemoveRunner` |
| `SecretService` | `SetSecret`, `ListSecrets` (names only, never values), `DeleteSecret`; scopes instance, repository and environment |
| `VariableService` | `SetVariable`, `GetVariable`, `ListVariables`, `DeleteVariable`; same scopes |
| `EnvironmentService` | `CreateEnvironment`, `ListEnvironments`, `DeleteEnvironment` |
| `JobArtifactService` | `CreateJobArtifact`, `UploadJobArtifactBlock` (client stream), `CommitJobArtifact`, `FinalizeJobArtifact`, `ListJobArtifacts`, `DownloadJobArtifact` (streamed, resumable), `DeleteJobArtifact`; artifacts v4 for CI runners, scoped to the run of the job token |
| `ArtifactService` | `ListArtifacts`, `DownloadArtifact` (streamed, resumable), `DeleteArtifact`; artifacts of a run for people |
| `ServerService` | `GetVersion` (server version and feature list), `GetSigningKey` (commit signing key) |

Tokens carry scopes only; presets such as `cli` or `agent` are a client concern.

## Layout

- `proto/ermite/`: `.proto` files of the `ermite` package, one per service (`repo`, `git`, `merge_request`, `token`, `session`, `event`, `server`, `runner`, `run`, `ci_settings`, `artifact`) plus shared `types`. Checked by `buf lint` and `buf breaking`.
- `gen/ermite/`: generated Go messages. `gen/ermite/ermiteconnect/`: generated Connect clients and handlers. Committed, never edited by hand.

## Generate

```sh
buf generate
```

## License

MIT. See [LICENSE](LICENSE).
