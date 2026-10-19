# Services

| Service | RPCs |
|---|---|
| `RepoService` | `ListRepos`, `GetRepo`, `CreateRepo`, `UpdateRepo` (field mask, merge settings), `DeleteRepo` |
| `GitService` | `ListRefs`, `GetTree`, `ListPaths`, `GetBlob` (streamed), `ListCommits`, `GetCommit` (structured diff), `ResolveRevision` |
| `TokenService` | `CreateToken` (secret shown once), `ListTokens`, `RevokeToken` |
| `MergeRequestService` | `CreateMergeRequest`, `ListMergeRequests`, `GetMergeRequest` (commits, diff, comments), `AddComment`, `EditComment`, `DeleteComment`, `ResolveComment`, `CloseMergeRequest`, `MergeMergeRequest`, `MarkRead` |
| `IssueService` | `CreateIssue`, `ListIssues`, `GetIssue` (comments), `AddIssueComment`, `CloseIssue`, `ReopenIssue`; owner only: `ListPendingIssues`, `ApproveIssue`, `RejectIssue`. Feature `issues`. Issues share numbers with merge requests |
| `SessionService` | `CreateLoginLink`, `SignOutEverywhere` (owner only) |
| `EventService` | `Watch` (streamed, resumable). Topics `home`, `repo:<name>` and `owner` (owner rights only). Events 17–21: `issue_opened`, `issue_closed`, `issue_reopened`, `issue_commented`, and `issue_submitted` (only on `owner`, no content) |
| `ServerService` | `GetVersion` (version and features), `GetSigningKey`, `RotateOIDCKey` (owner only) |
| `TrustedKeyService` | `ListTrustedKeys`, `AddTrustedKey` and `RemoveTrustedKey` (owner only). `Commit.verification` carries the result of the check (feature `commit_signatures`) |
| `RunService` | `ListRuns`, `GetRun`, `GetJob`, `GetJobLog` (streamed, follow), `ListAnnotations`, `GetStepSummaries`, `CancelRun`, `RerunRun`, `DispatchWorkflow`, `ListWorkflows`, `DisableWorkflow`, `EnableWorkflow`, `ListChecks`, `ListRunners`, `RemoveRunner` |
| `RunnerService` | `RegisterRunner`, `FetchJob` (long poll), `UpdateJob`, `AppendLog` (offset-idempotent), `Heartbeat`, `WatchJob` (cancel requests, streamed), `CompleteJob` (idempotent), `ResolveActions`, `UploadSummary` |
| `SecretService` | `SetSecret`, `ListSecrets` (names only), `DeleteSecret`. Scopes: instance, repository, environment |
| `VariableService` | `SetVariable`, `GetVariable`, `ListVariables`, `DeleteVariable`. Same scopes |
| `EnvironmentService` | `CreateEnvironment`, `ListEnvironments`, `DeleteEnvironment` |
| `ArtifactService` | `ListArtifacts`, `DownloadArtifact` (streamed, resumable), `DeleteArtifact` |
| `JobArtifactService` | `CreateJobArtifact`, `UploadJobArtifactBlock` (client stream), `CommitJobArtifact`, `FinalizeJobArtifact`, `ListJobArtifacts`, `DownloadJobArtifact` (streamed, resumable), `DeleteJobArtifact`. Scoped to the run of the job token |
