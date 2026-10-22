# Services

| Service | RPCs |
|---|---|
| `RepoService` | `ListRepos`, `GetRepo`, `CreateRepo`, `UpdateRepo` (field mask, merge settings), `DeleteRepo` |
| `GitService` | `ListRefs`, `GetTree`, `ListPaths`, `GetBlob` (streamed), `ListCommits`, `GetCommit` (structured diff), `ResolveRevision` |
| `TokenService` | `CreateToken` (secret shown once), `ListTokens`, `RevokeToken` |
| `MergeRequestService` | `CreateMergeRequest`, `ListMergeRequests`, `GetMergeRequest` (commits, diff, comments), `AddComment`, `EditComment`, `DeleteComment`, `ResolveComment`, `CloseMergeRequest`, `MarkReady`, `MergeMergeRequest` (modes fast-forward, squash, merge commit), `MarkRead`, `EnableAutoMerge`, `CancelAutoMerge` (merge by itself once a check passed on the head and all its runs are complete; setting `allow_auto_merge`, `MergeRequest.auto_merge`, events `MergeRequestAutoMergeEnabled` and `MergeRequestAutoMergeCancelled`, `MergeRequestMerged.auto`; feature `merge_requests.auto_merge`). Open merge requests carry `mergeability` (up to date, behind, conflict, or checking) and `Watch` publishes `MergeRequestMergeability` each time a result is settled for new branch tips (feature `merge_requests.mergeability`) |
| `IssueService` | `CreateIssue`, `ListIssues`, `GetIssue` (comments, merge requests naming the issue, history events, `closed_by_merge_request`), `AddIssueComment`, `CloseIssue`, `ResolveIssue`, `ReopenIssue`; owner only: `ListPendingIssues`, `ApproveIssue`, `RejectIssue`. States `OPEN`, `CLOSED`, `RESOLVED`; `ResolveIssue` needs `issue:write`. Features `issues`, `issues.links`, `issues.resolve`. Issues share numbers with merge requests |
| `SessionService` | `CreateLoginLink`, `SignOutEverywhere` (owner only) |
| `EventService` | `Watch` (streamed, resumable). Topics `home`, `repo:<name>` and `owner` (owner rights only). Events 17–21: `issue_opened`, `issue_closed`, `issue_resolved` (event 31), `issue_reopened`, `issue_commented`, and `issue_submitted` (only on `owner`, no content). Event 22: `merge_request_mergeability` |
| `ServerService` | `GetVersion` (version and features), `GetSigningKey`, `RotateOIDCKey` (owner only) |
| `TrustedKeyService` | `ListTrustedKeys`, `AddTrustedKey` and `RemoveTrustedKey` (owner only). `Commit.verification` carries the result of the check (feature `commit_signatures`) |
| `RunService` | `ListRuns`, `GetRun`, `GetJob`, `GetJobLog` (streamed, follow), `ListAnnotations`, `GetStepSummaries`, `CancelRun`, `RerunRun`, `DispatchWorkflow`, `ListWorkflows`, `DisableWorkflow`, `EnableWorkflow`, `ListChecks`, `ListRunners`, `RemoveRunner` |
| `RunnerService` | `RegisterRunner`, `FetchJob` (long poll), `UpdateJob`, `AppendLog` (offset-idempotent), `Heartbeat`, `WatchJob` (cancel requests, streamed), `CompleteJob` (idempotent), `ResolveActions`, `UploadSummary` |
| `SecretService` | `SetSecret`, `ListSecrets` (names only), `DeleteSecret`. Scopes: instance, repository, environment |
| `VariableService` | `SetVariable`, `GetVariable`, `ListVariables`, `DeleteVariable`. Same scopes |
| `EnvironmentService` | `CreateEnvironment`, `ListEnvironments`, `DeleteEnvironment` |
| `ArtifactService` | `ListArtifacts`, `DownloadArtifact` (streamed, resumable), `DeleteArtifact` |
| `JobArtifactService` | `CreateJobArtifact`, `UploadJobArtifactBlock` (client stream), `CommitJobArtifact`, `FinalizeJobArtifact`, `ListJobArtifacts`, `DownloadJobArtifact` (streamed, resumable), `DeleteJobArtifact`. Scoped to the run of the job token |

## MirrorService

Servers implementing push mirrors advertise `mirror` in `GetVersionResponse.features`.
No new token scope is introduced.

| RPC | Access and behavior |
|---|---|
| `GetMirrorKey` | Non-anonymous caller; generates the public SSH key on first call. Without `ERMITE_SECRETS_KEY`, returns `failed_precondition` |
| `RotateMirrorKey` | Owner rights: admin token or a token with every scope and no repository or branch limits. Replaces the key immediately and returns the previous fingerprint |
| `GetMirrorConfig`, `SetMirrorConfig` | Owner rights. Set patches only present fields; invalid fields return `invalid_argument` with field violations |
| `SetRepoMirror` | `repo:admin` on the repository. Sets an override, even if global mirroring is disabled; empty cron and absent all-branches inherit the global values |
| `RemoveRepoMirror` | `repo:admin` on the repository. Removes the override |
| `SetRepoMirrorExcluded` | `repo:admin` on the repository. Opts out of or into the global template; a URL override wins over exclusion |
| `GetMirrorStatus` | Read access to the repository. Error `detail` is returned only with `repo:admin` |
| `ListMirrorStatus` | Owner rights |
| `SyncMirror` | `repo:admin` for `repo`, or owner rights for `all=true` (with empty `repo`). Returns per-repository results, including queued and no-op indicators and a per-repository `error`; `all=true` continues past failing repositories and skips disabled ones |

Cron expressions have five fields, run in UTC and must be at least five minutes
apart. Mirrors push the default branch (or every branch with `all_branches`) and
tags. `force` lifts a pause, skips the no-op check and the 60-second manual sync
throttle; it **never** performs a Git force push. Pushes are atomic: one rejected
ref fails the whole push. After key rotation, add the new public key to remotes
and sync with `force` to recover from authentication failures. Configuration can
be saved without `ERMITE_SECRETS_KEY`, but pushes cannot run without it.

`EventService.Watch` adds `mirror_failed` (field 23, `MirrorFailed`) when a mirror
starts failing (ok or disabled to failing or paused, once per incident), and `mirror_recovered`
(field 24, `MirrorRecovered`) when a failing or paused mirror succeeds again.
These are transition events, not events for every attempt. They are delivered on
`repo:<name>` and `home`, subject to repository visibility. Neither event carries
a URL, reason or error detail (`MirrorFailed` field 4, `reason`, was removed).

## ReleaseService

Servers advertise `releases` in `GetVersionResponse.features`.

| RPC | Behavior |
|---|---|
| `CreateRelease`, `GetRelease`, `ListReleases`, `UpdateRelease`, `DeleteRelease` | Manage releases; get selects by id, tag or latest. Lists are paginated, newest first. Updates use a field mask |
| `UploadReleaseAsset` | Client stream: metadata and nonnegative exact size first, then data in chunks of at most 1 MiB. `replace` replaces a taken name |
| `UpdateReleaseAsset`, `DeleteReleaseAsset`, `ListReleaseAssets` | Rename/relabel with a field mask, delete, or list uploaded assets by name |
| `DownloadReleaseAsset` | Select by asset id, tag/name or latest/name; metadata first, then chunks of at most 1 MiB. Resume with a byte offset |
| `GetReleaseUsage` | Repository usage including reserved uploads, quota, asset size and count limits |

Published releases follow repository read access; drafts are visible only to
writers. Writes require `release:write` (`SCOPE_RELEASE_WRITE = 13`) or a job
token with `contents: write` for that repository. Branch-limited tokens cannot
write releases. Unreadable objects return `not_found`. Duplicate tags or asset
names return `already_exists`; limits return `resource_exhausted`. Publishing
requires an existing tag or `target_commitish` to create it; existing tags never
move. `make_latest` is a boolean pin, and updates use mask presence.

`EventService.Watch` adds `release_published` (27), `release_updated` (28),
`release_deleted` (29) and `release_unpublished` (30), delivered on `repo:<name>`
and `home` subject to repository visibility. Payloads carry repository, release
id and tag; all except deletion also carry draft and prerelease flags. Tag
deletion turns a release into a draft without deleting its assets.
