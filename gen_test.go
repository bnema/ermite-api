// SPDX-License-Identifier: MIT

package api_test

import (
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"bnema.dev/ermite-api/gen/ermite"
	"bnema.dev/ermite-api/gen/ermite/ermiteconnect"
)

// TestProcedures pins the procedure paths: they are the wire contract.
func TestProcedures(t *testing.T) {
	tests := []struct{ got, want string }{
		{ermiteconnect.RepoServiceListReposProcedure, "/ermite.RepoService/ListRepos"},
		{ermiteconnect.RepoServiceGetRepoProcedure, "/ermite.RepoService/GetRepo"},
		{ermiteconnect.RepoServiceCreateRepoProcedure, "/ermite.RepoService/CreateRepo"},
		{ermiteconnect.RepoServiceUpdateRepoProcedure, "/ermite.RepoService/UpdateRepo"},
		{ermiteconnect.RepoServiceDeleteRepoProcedure, "/ermite.RepoService/DeleteRepo"},
		{ermiteconnect.GitServiceListRefsProcedure, "/ermite.GitService/ListRefs"},
		{ermiteconnect.GitServiceGetTreeProcedure, "/ermite.GitService/GetTree"},
		{ermiteconnect.GitServiceListPathsProcedure, "/ermite.GitService/ListPaths"},
		{ermiteconnect.GitServiceGetBlobProcedure, "/ermite.GitService/GetBlob"},
		{ermiteconnect.GitServiceListCommitsProcedure, "/ermite.GitService/ListCommits"},
		{ermiteconnect.GitServiceGetCommitProcedure, "/ermite.GitService/GetCommit"},
		{ermiteconnect.GitServiceResolveRevisionProcedure, "/ermite.GitService/ResolveRevision"},
		{ermiteconnect.TokenServiceCreateTokenProcedure, "/ermite.TokenService/CreateToken"},
		{ermiteconnect.TokenServiceListTokensProcedure, "/ermite.TokenService/ListTokens"},
		{ermiteconnect.TokenServiceRevokeTokenProcedure, "/ermite.TokenService/RevokeToken"},
		{ermiteconnect.EventServiceWatchProcedure, "/ermite.EventService/Watch"},
		{ermiteconnect.MergeRequestServiceCreateMergeRequestProcedure, "/ermite.MergeRequestService/CreateMergeRequest"},
		{ermiteconnect.MergeRequestServiceListMergeRequestsProcedure, "/ermite.MergeRequestService/ListMergeRequests"},
		{ermiteconnect.MergeRequestServiceGetMergeRequestProcedure, "/ermite.MergeRequestService/GetMergeRequest"},
		{ermiteconnect.MergeRequestServiceAddCommentProcedure, "/ermite.MergeRequestService/AddComment"},
		{ermiteconnect.MergeRequestServiceResolveCommentProcedure, "/ermite.MergeRequestService/ResolveComment"},
		{ermiteconnect.MergeRequestServiceCloseMergeRequestProcedure, "/ermite.MergeRequestService/CloseMergeRequest"},
		{ermiteconnect.MergeRequestServiceMergeMergeRequestProcedure, "/ermite.MergeRequestService/MergeMergeRequest"},
		{ermiteconnect.MergeRequestServiceMarkReadProcedure, "/ermite.MergeRequestService/MarkRead"},
		{ermiteconnect.SessionServiceCreateLoginLinkProcedure, "/ermite.SessionService/CreateLoginLink"},
		{ermiteconnect.SessionServiceSignOutEverywhereProcedure, "/ermite.SessionService/SignOutEverywhere"},
		{ermiteconnect.ServerServiceGetVersionProcedure, "/ermite.ServerService/GetVersion"},
		{ermiteconnect.ServerServiceGetSigningKeyProcedure, "/ermite.ServerService/GetSigningKey"},
		{ermiteconnect.RunnerServiceRegisterRunnerProcedure, "/ermite.RunnerService/RegisterRunner"},
		{ermiteconnect.RunnerServiceFetchJobProcedure, "/ermite.RunnerService/FetchJob"},
		{ermiteconnect.RunnerServiceUpdateJobProcedure, "/ermite.RunnerService/UpdateJob"},
		{ermiteconnect.RunnerServiceAppendLogProcedure, "/ermite.RunnerService/AppendLog"},
		{ermiteconnect.RunnerServiceHeartbeatProcedure, "/ermite.RunnerService/Heartbeat"},
		{ermiteconnect.RunnerServiceWatchJobProcedure, "/ermite.RunnerService/WatchJob"},
		{ermiteconnect.RunServiceListRunsProcedure, "/ermite.RunService/ListRuns"},
		{ermiteconnect.RunServiceGetRunProcedure, "/ermite.RunService/GetRun"},
		{ermiteconnect.RunServiceGetJobProcedure, "/ermite.RunService/GetJob"},
		{ermiteconnect.RunServiceGetJobLogProcedure, "/ermite.RunService/GetJobLog"},
		{ermiteconnect.RunServiceCancelRunProcedure, "/ermite.RunService/CancelRun"},
		{ermiteconnect.RunServiceRerunRunProcedure, "/ermite.RunService/RerunRun"},
		{ermiteconnect.RunServiceListChecksProcedure, "/ermite.RunService/ListChecks"},
		{ermiteconnect.RunServiceListRunnersProcedure, "/ermite.RunService/ListRunners"},
		{ermiteconnect.RunServiceRemoveRunnerProcedure, "/ermite.RunService/RemoveRunner"},
	}
	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("procedure = %q, want %q", tt.got, tt.want)
		}
	}
}

// TestEnumNumbers pins enum numbers: they are part of the wire contract.
func TestEnumNumbers(t *testing.T) {
	if got := ermite.Scope_SCOPE_TOKEN_ADMIN; got != 9 {
		t.Errorf("SCOPE_TOKEN_ADMIN = %d, want 9", got)
	}
	if got := ermite.FileStatus_FILE_STATUS_TYPE_CHANGED; got != 6 {
		t.Errorf("FILE_STATUS_TYPE_CHANGED = %d, want 6", got)
	}
	if got := ermite.Visibility_VISIBILITY_PRIVATE; got != 2 {
		t.Errorf("VISIBILITY_PRIVATE = %d, want 2", got)
	}
	if got := ermite.MergeMode_MERGE_MODE_SQUASH; got != 2 {
		t.Errorf("MERGE_MODE_SQUASH = %d, want 2", got)
	}
	if got := ermite.MergeRequestState_MERGE_REQUEST_STATE_CLOSED; got != 3 {
		t.Errorf("MERGE_REQUEST_STATE_CLOSED = %d, want 3", got)
	}
	if got := ermite.DiffSide_DIFF_SIDE_NEW; got != 2 {
		t.Errorf("DIFF_SIDE_NEW = %d, want 2", got)
	}
	if got := ermite.CiStatus_CI_STATUS_COMPLETED; got != 3 {
		t.Errorf("RUN_STATUS_COMPLETED = %d, want 3", got)
	}
	if got := ermite.CiConclusion_CI_CONCLUSION_ACTION_REQUIRED; got != 7 {
		t.Errorf("CONCLUSION_ACTION_REQUIRED = %d, want 7", got)
	}
	if got := ermite.StepState_STEP_STATE_COMPLETED; got != 3 {
		t.Errorf("STEP_STATE_COMPLETED = %d, want 3", got)
	}
	if got := ermite.RunnerIsolation_RUNNER_ISOLATION_CONTAINER_PRIVILEGED; got != 3 {
		t.Errorf("RUNNER_ISOLATION_CONTAINER_PRIVILEGED = %d, want 3", got)
	}
	if got := ermite.CancelReason_CANCEL_REASON_TIMEOUT; got != 3 {
		t.Errorf("CANCEL_REASON_TIMEOUT = %d, want 3", got)
	}
}

// TestEventRoundTrip checks that an event with a oneof kind survives the wire.
func TestEventRoundTrip(t *testing.T) {
	in := &ermite.WatchResponse{Event: &ermite.Event{
		Id: "42",
		Kind: &ermite.Event_Pushed{Pushed: &ermite.Pushed{
			Repo: "ermite",
			RefUpdates: []*ermite.RefUpdate{
				{Ref: "refs/heads/main", OldId: "a", NewId: "b"},
			},
		}},
	}}
	b, err := proto.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	var out ermite.WatchResponse
	if err := proto.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(in, &out) {
		t.Errorf("round trip mismatch: %v != %v", &out, in)
	}
	if out.GetEvent().GetPushed().GetRefUpdates()[0].GetNewId() != "b" {
		t.Errorf("unexpected ref update: %v", &out)
	}
}

// TestUpdateRepoMask checks that the request carries a field mask.
func TestUpdateRepoMask(t *testing.T) {
	req := &ermite.UpdateRepoRequest{
		Repo:       "ermite",
		Visibility: ermite.Visibility_VISIBILITY_PUBLIC,
		UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"visibility"}},
	}
	if got := req.GetUpdateMask().GetPaths(); len(got) != 1 || got[0] != "visibility" {
		t.Errorf("paths = %v", got)
	}
}

// TestMergeRequestEventRoundTrip checks the merge request event kinds and the
// repository settings survive the wire.
func TestMergeRequestEventRoundTrip(t *testing.T) {
	in := &ermite.WatchResponse{Event: &ermite.Event{
		Id: "43",
		Kind: &ermite.Event_CommentAdded{CommentAdded: &ermite.CommentAdded{
			Repo:   "ermite",
			Number: 7,
			Comment: &ermite.Comment{
				Id: "c1", Path: "a.go", Line: 3, Side: ermite.DiffSide_DIFF_SIDE_OLD, Text: "hi",
			},
		}},
	}}
	b, err := proto.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	var out ermite.WatchResponse
	if err := proto.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(in, &out) {
		t.Errorf("round trip mismatch: %v != %v", &out, in)
	}
	if got := out.GetEvent().GetCommentAdded().GetNumber(); got != 7 {
		t.Errorf("number = %d, want 7", got)
	}

	req := &ermite.UpdateRepoRequest{
		Repo: "ermite",
		Settings: &ermite.RepoSettings{
			AllowedMergeModes:  []ermite.MergeMode{ermite.MergeMode_MERGE_MODE_SQUASH},
			FailingChecksBlock: true,
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"settings.allowed_merge_modes", "settings.failing_checks_block"}},
	}
	if got := req.GetSettings().GetAllowedMergeModes(); len(got) != 1 || got[0] != ermite.MergeMode_MERGE_MODE_SQUASH {
		t.Errorf("modes = %v", got)
	}
}

// TestCIEventRoundTrip checks that the CI event kinds keep their field numbers
// and survive the wire.
func TestCIEventRoundTrip(t *testing.T) {
	in := &ermite.WatchResponse{Event: &ermite.Event{
		Id: "44",
		Kind: &ermite.Event_JobUpdated{JobUpdated: &ermite.JobUpdated{Job: &ermite.RunJob{
			Id:         "j1",
			RunId:      9,
			Attempt:    2,
			Status:     ermite.CiStatus_CI_STATUS_COMPLETED,
			Conclusion: ermite.CiConclusion_CI_CONCLUSION_TIMED_OUT,
			Isolation:  ermite.RunnerIsolation_RUNNER_ISOLATION_VM,
			Steps:      []*ermite.Step{{Number: 1, Name: "Set up job", LogStart: 0, LogEnd: 3}},
		}}},
	}}
	b, err := proto.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	var out ermite.WatchResponse
	if err := proto.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(in, &out) {
		t.Errorf("round trip mismatch: %v != %v", &out, in)
	}
	f := (&ermite.Event{}).ProtoReflect().Descriptor().Fields()
	for name, num := range map[string]int{"run_updated": 12, "job_updated": 13} {
		if fd := f.ByName(protoreflect.Name(name)); fd == nil || int(fd.Number()) != num {
			t.Errorf("event field %s: %v, want number %d", name, fd, num)
		}
	}
}

// TestJobRoundTrip checks the job handed to a runner survives the wire.
func TestJobRoundTrip(t *testing.T) {
	in := &ermite.FetchJobResponse{Job: &ermite.Job{
		JobId: "j1", RunId: 9, Attempt: 1, JobToken: "tok", WorkflowPath: ".github/workflows/ci.yml",
		WorkflowYaml: "on: push", JobKey: "build", Matrix: "{}", Contexts: "{}",
		Labels: []string{"self-hosted"}, TimeoutMinutes: 360, Repo: "ermite",
		Sha: "0123456789012345678901234567890123456789", Ref: "refs/heads/main",
	}}
	b, err := proto.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	var out ermite.FetchJobResponse
	if err := proto.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(in, &out) {
		t.Errorf("round trip mismatch: %v != %v", &out, in)
	}
}
