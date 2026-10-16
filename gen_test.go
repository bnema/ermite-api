// SPDX-License-Identifier: MIT

package api_test

import (
	"errors"
	"testing"

	"connectrpc.com/connect"

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
		{ermiteconnect.RunnerServiceCompleteJobProcedure, "/ermite.RunnerService/CompleteJob"},
		{ermiteconnect.RunnerServiceResolveActionsProcedure, "/ermite.RunnerService/ResolveActions"},
		{ermiteconnect.RunnerServiceUploadSummaryProcedure, "/ermite.RunnerService/UploadSummary"},
		{ermiteconnect.RunServiceListRunsProcedure, "/ermite.RunService/ListRuns"},
		{ermiteconnect.RunServiceGetRunProcedure, "/ermite.RunService/GetRun"},
		{ermiteconnect.RunServiceGetJobProcedure, "/ermite.RunService/GetJob"},
		{ermiteconnect.RunServiceGetJobLogProcedure, "/ermite.RunService/GetJobLog"},
		{ermiteconnect.RunServiceListAnnotationsProcedure, "/ermite.RunService/ListAnnotations"},
		{ermiteconnect.RunServiceGetStepSummariesProcedure, "/ermite.RunService/GetStepSummaries"},
		{ermiteconnect.RunServiceCancelRunProcedure, "/ermite.RunService/CancelRun"},
		{ermiteconnect.RunServiceRerunRunProcedure, "/ermite.RunService/RerunRun"},
		{ermiteconnect.RunServiceDispatchWorkflowProcedure, "/ermite.RunService/DispatchWorkflow"},
		{ermiteconnect.RunServiceListChecksProcedure, "/ermite.RunService/ListChecks"},
		{ermiteconnect.RunServiceListRunnersProcedure, "/ermite.RunService/ListRunners"},
		{ermiteconnect.RunServiceRemoveRunnerProcedure, "/ermite.RunService/RemoveRunner"},
		{ermiteconnect.SecretServiceSetSecretProcedure, "/ermite.SecretService/SetSecret"},
		{ermiteconnect.SecretServiceListSecretsProcedure, "/ermite.SecretService/ListSecrets"},
		{ermiteconnect.SecretServiceDeleteSecretProcedure, "/ermite.SecretService/DeleteSecret"},
		{ermiteconnect.VariableServiceSetVariableProcedure, "/ermite.VariableService/SetVariable"},
		{ermiteconnect.VariableServiceGetVariableProcedure, "/ermite.VariableService/GetVariable"},
		{ermiteconnect.VariableServiceListVariablesProcedure, "/ermite.VariableService/ListVariables"},
		{ermiteconnect.VariableServiceDeleteVariableProcedure, "/ermite.VariableService/DeleteVariable"},
		{ermiteconnect.EnvironmentServiceCreateEnvironmentProcedure, "/ermite.EnvironmentService/CreateEnvironment"},
		{ermiteconnect.EnvironmentServiceListEnvironmentsProcedure, "/ermite.EnvironmentService/ListEnvironments"},
		{ermiteconnect.EnvironmentServiceDeleteEnvironmentProcedure, "/ermite.EnvironmentService/DeleteEnvironment"},
		{ermiteconnect.JobArtifactServiceCreateJobArtifactProcedure, "/ermite.JobArtifactService/CreateJobArtifact"},
		{ermiteconnect.JobArtifactServiceUploadJobArtifactBlockProcedure, "/ermite.JobArtifactService/UploadJobArtifactBlock"},
		{ermiteconnect.JobArtifactServiceCommitJobArtifactProcedure, "/ermite.JobArtifactService/CommitJobArtifact"},
		{ermiteconnect.JobArtifactServiceFinalizeJobArtifactProcedure, "/ermite.JobArtifactService/FinalizeJobArtifact"},
		{ermiteconnect.JobArtifactServiceListJobArtifactsProcedure, "/ermite.JobArtifactService/ListJobArtifacts"},
		{ermiteconnect.JobArtifactServiceDownloadJobArtifactProcedure, "/ermite.JobArtifactService/DownloadJobArtifact"},
		{ermiteconnect.JobArtifactServiceDeleteJobArtifactProcedure, "/ermite.JobArtifactService/DeleteJobArtifact"},
		{ermiteconnect.ArtifactServiceListArtifactsProcedure, "/ermite.ArtifactService/ListArtifacts"},
		{ermiteconnect.ArtifactServiceDownloadArtifactProcedure, "/ermite.ArtifactService/DownloadArtifact"},
		{ermiteconnect.ArtifactServiceDeleteArtifactProcedure, "/ermite.ArtifactService/DeleteArtifact"},
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
	enums := []struct {
		name string
		got  int32
		want int32
	}{
		{"CI_STATUS_UNSPECIFIED", int32(ermite.CiStatus_CI_STATUS_UNSPECIFIED), 0},
		{"CI_STATUS_QUEUED", int32(ermite.CiStatus_CI_STATUS_QUEUED), 1},
		{"CI_STATUS_IN_PROGRESS", int32(ermite.CiStatus_CI_STATUS_IN_PROGRESS), 2},
		{"CI_STATUS_COMPLETED", int32(ermite.CiStatus_CI_STATUS_COMPLETED), 3},
		{"CI_STATUS_PENDING", int32(ermite.CiStatus_CI_STATUS_PENDING), 4},
		{"CI_CONCLUSION_UNSPECIFIED", int32(ermite.CiConclusion_CI_CONCLUSION_UNSPECIFIED), 0},
		{"CI_CONCLUSION_SUCCESS", int32(ermite.CiConclusion_CI_CONCLUSION_SUCCESS), 1},
		{"CI_CONCLUSION_FAILURE", int32(ermite.CiConclusion_CI_CONCLUSION_FAILURE), 2},
		{"CI_CONCLUSION_CANCELLED", int32(ermite.CiConclusion_CI_CONCLUSION_CANCELLED), 3},
		{"CI_CONCLUSION_SKIPPED", int32(ermite.CiConclusion_CI_CONCLUSION_SKIPPED), 4},
		{"CI_CONCLUSION_TIMED_OUT", int32(ermite.CiConclusion_CI_CONCLUSION_TIMED_OUT), 5},
		{"CI_CONCLUSION_NEUTRAL", int32(ermite.CiConclusion_CI_CONCLUSION_NEUTRAL), 6},
		{"CI_CONCLUSION_ACTION_REQUIRED", int32(ermite.CiConclusion_CI_CONCLUSION_ACTION_REQUIRED), 7},
		{"STEP_STATE_UNSPECIFIED", int32(ermite.StepState_STEP_STATE_UNSPECIFIED), 0},
		{"STEP_STATE_PENDING", int32(ermite.StepState_STEP_STATE_PENDING), 1},
		{"STEP_STATE_IN_PROGRESS", int32(ermite.StepState_STEP_STATE_IN_PROGRESS), 2},
		{"STEP_STATE_COMPLETED", int32(ermite.StepState_STEP_STATE_COMPLETED), 3},
		{"RUNNER_ISOLATION_UNSPECIFIED", int32(ermite.RunnerIsolation_RUNNER_ISOLATION_UNSPECIFIED), 0},
		{"RUNNER_ISOLATION_VM", int32(ermite.RunnerIsolation_RUNNER_ISOLATION_VM), 1},
		{"RUNNER_ISOLATION_CONTAINER_SYSBOX", int32(ermite.RunnerIsolation_RUNNER_ISOLATION_CONTAINER_SYSBOX), 2},
		{"RUNNER_ISOLATION_CONTAINER_PRIVILEGED", int32(ermite.RunnerIsolation_RUNNER_ISOLATION_CONTAINER_PRIVILEGED), 3},
		{"CANCEL_REASON_UNSPECIFIED", int32(ermite.CancelReason_CANCEL_REASON_UNSPECIFIED), 0},
		{"CANCEL_REASON_USER", int32(ermite.CancelReason_CANCEL_REASON_USER), 1},
		{"CANCEL_REASON_FAIL_FAST", int32(ermite.CancelReason_CANCEL_REASON_FAIL_FAST), 2},
		{"CANCEL_REASON_TIMEOUT", int32(ermite.CancelReason_CANCEL_REASON_TIMEOUT), 3},
		{"ANNOTATION_LEVEL_UNSPECIFIED", int32(ermite.AnnotationLevel_ANNOTATION_LEVEL_UNSPECIFIED), 0},
		{"ANNOTATION_LEVEL_NOTICE", int32(ermite.AnnotationLevel_ANNOTATION_LEVEL_NOTICE), 1},
		{"ANNOTATION_LEVEL_WARNING", int32(ermite.AnnotationLevel_ANNOTATION_LEVEL_WARNING), 2},
		{"ANNOTATION_LEVEL_FAILURE", int32(ermite.AnnotationLevel_ANNOTATION_LEVEL_FAILURE), 3},
		{"ACTION_SOURCE_UNSPECIFIED", int32(ermite.ActionSource_ACTION_SOURCE_UNSPECIFIED), 0},
		{"ACTION_SOURCE_GITHUB", int32(ermite.ActionSource_ACTION_SOURCE_GITHUB), 1},
		{"ACTION_SOURCE_ERMITE", int32(ermite.ActionSource_ACTION_SOURCE_ERMITE), 2},
	}
	for _, e := range enums {
		if e.got != e.want {
			t.Errorf("%s = %d, want %d", e.name, e.got, e.want)
		}
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

// TestWatchJobRoundTrip checks the three kinds of WatchJob message and the
// JobEnded detail.
func TestWatchJobRoundTrip(t *testing.T) {
	msgs := []*ermite.WatchJobResponse{
		{Msg: &ermite.WatchJobResponse_Keepalive{Keepalive: &ermite.Keepalive{}}},
		{Msg: &ermite.WatchJobResponse_Cancel{Cancel: &ermite.CancelRequest{
			Reason: ermite.CancelReason_CANCEL_REASON_FAIL_FAST, Message: "m",
		}}},
		{Msg: &ermite.WatchJobResponse_Ended{Ended: &ermite.JobEnded{
			Conclusion: ermite.CiConclusion_CI_CONCLUSION_CANCELLED, Error: "e",
		}}},
	}
	for _, in := range msgs {
		b, err := proto.Marshal(in)
		if err != nil {
			t.Fatal(err)
		}
		var out ermite.WatchJobResponse
		if err := proto.Unmarshal(b, &out); err != nil {
			t.Fatal(err)
		}
		if !proto.Equal(in, &out) {
			t.Errorf("round trip mismatch: %v != %v", &out, in)
		}
	}
	// The detail travels in a Connect error.
	d, err := connect.NewErrorDetail(&ermite.JobEnded{Conclusion: ermite.CiConclusion_CI_CONCLUSION_TIMED_OUT})
	if err != nil {
		t.Fatal(err)
	}
	cerr := connect.NewError(connect.CodeFailedPrecondition, errors.New("job ended"))
	cerr.AddDetail(d)
	v, err := cerr.Details()[0].Value()
	if err != nil {
		t.Fatal(err)
	}
	if got := v.(*ermite.JobEnded).GetConclusion(); got != ermite.CiConclusion_CI_CONCLUSION_TIMED_OUT {
		t.Errorf("detail conclusion = %v", got)
	}
}

// TestCompleteJobRoundTrip checks the final report of a job and its answer.
func TestCompleteJobRoundTrip(t *testing.T) {
	in := &ermite.CompleteJobRequest{
		JobId:          "j1",
		Conclusion:     ermite.CiConclusion_CI_CONCLUSION_CANCELLED,
		Outputs:        map[string]string{"a": "b"},
		RunnerShutdown: true,
		Steps:          []*ermite.StepUpdate{{Number: 2, State: ermite.StepState_STEP_STATE_COMPLETED}},
	}
	b, err := proto.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	var out ermite.CompleteJobRequest
	if err := proto.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(in, &out) {
		t.Errorf("round trip mismatch: %v != %v", &out, in)
	}
	r := &ermite.CompleteJobResponse{Applied: false, Conclusion: ermite.CiConclusion_CI_CONCLUSION_TIMED_OUT}
	if r.GetApplied() || r.GetConclusion() != ermite.CiConclusion_CI_CONCLUSION_TIMED_OUT {
		t.Errorf("unexpected response %v", r)
	}
}

// TestRunNameField pins the number of Run.name, the GitHub name of the run
// that sits next to workflow_name, and that it survives the wire.
func TestRunNameField(t *testing.T) {
	fd := (&ermite.Run{}).ProtoReflect().Descriptor().Fields().ByName("name")
	if fd == nil || fd.Number() != 19 {
		t.Fatalf("Run.name = %v, want field 19", fd)
	}
	in := &ermite.Run{WorkflowName: "Names Workflow", Name: "run push on branch by someone", Title: "run push on branch by someone"}
	b, err := proto.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	var out ermite.Run
	if err := proto.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if out.GetName() != in.GetName() || out.GetWorkflowName() != "Names Workflow" {
		t.Errorf("round trip: name %q, workflow_name %q", out.GetName(), out.GetWorkflowName())
	}
}

// TestCI2FieldNumbers pins the numbers of the fields added for CI-2.
func TestCI2FieldNumbers(t *testing.T) {
	tests := []struct {
		msg    proto.Message
		fields map[string]int
	}{
		{&ermite.Job{}, map[string]int{"github_token": 15}},
		{&ermite.UpdateJobRequest{}, map[string]int{"annotations": 3}},
		{&ermite.CompleteJobRequest{}, map[string]int{"annotations": 6}},
		{&ermite.StepUpdate{}, map[string]int{"outcome": 7}},
		{&ermite.Step{}, map[string]int{"outcome": 9}},
		{&ermite.Annotation{}, map[string]int{
			"level": 1, "title": 2, "message": 3, "file": 4, "start_line": 5,
			"end_line": 6, "start_column": 7, "end_column": 8, "step_number": 9,
		}},
		{&ermite.ActionRef{}, map[string]int{"uses": 1}},
		{&ermite.ResolveActionsRequest{}, map[string]int{"job_id": 1, "refs": 2}},
		{&ermite.ResolveActionsResponse{}, map[string]int{"actions": 1}},
		{&ermite.ResolvedAction{}, map[string]int{
			"ref": 1, "name_with_owner": 2, "path": 3, "resolved_sha": 4,
			"tarball_url": 5, "token": 6, "source": 7, "error": 8,
		}},
		{&ermite.UploadSummaryRequest{}, map[string]int{"job_id": 1, "step_number": 2, "markdown": 3}},
		{&ermite.ListAnnotationsRequest{}, map[string]int{"repo": 1, "run_id": 2, "attempt": 3, "job_id": 4}},
		{&ermite.ListAnnotationsResponse{}, map[string]int{"annotations": 1, "truncated": 2}},
		{&ermite.JobAnnotation{}, map[string]int{"job_id": 1, "annotation": 2}},
		{&ermite.GetStepSummariesRequest{}, map[string]int{"repo": 1, "job_id": 2, "attempt": 3}},
		{&ermite.GetStepSummariesResponse{}, map[string]int{"summaries": 1, "truncated": 2}},
		{&ermite.StepSummary{}, map[string]int{"step_number": 1, "markdown": 2}},
	}
	for _, tt := range tests {
		md := tt.msg.ProtoReflect().Descriptor()
		for name, num := range tt.fields {
			if fd := md.Fields().ByName(protoreflect.Name(name)); fd == nil || int(fd.Number()) != num {
				t.Errorf("%s.%s: %v, want number %d", md.Name(), name, fd, num)
			}
		}
	}
}

// TestCI2RoundTrip checks that the CI-2 messages survive the wire.
func TestCI2RoundTrip(t *testing.T) {
	msgs := []proto.Message{
		&ermite.Job{JobId: "j1", GithubToken: "ermitegh_x"},
		&ermite.UpdateJobRequest{
			JobId: "j1",
			Steps: []*ermite.StepUpdate{{
				Number: 2, State: ermite.StepState_STEP_STATE_COMPLETED,
				Conclusion: ermite.CiConclusion_CI_CONCLUSION_SUCCESS,
				Outcome:    ermite.CiConclusion_CI_CONCLUSION_FAILURE,
			}},
			Annotations: []*ermite.Annotation{{
				Level: ermite.AnnotationLevel_ANNOTATION_LEVEL_WARNING, Title: "t", Message: "m",
				File: "a.go", StartLine: 1, EndLine: 2, StartColumn: 3, EndColumn: 4, StepNumber: 2,
			}},
		},
		&ermite.CompleteJobRequest{
			JobId:       "j1",
			Conclusion:  ermite.CiConclusion_CI_CONCLUSION_SUCCESS,
			Annotations: []*ermite.Annotation{{Level: ermite.AnnotationLevel_ANNOTATION_LEVEL_FAILURE, Message: "m", StepNumber: 3}},
		},
		&ermite.ResolveActionsRequest{JobId: "j1", Refs: []*ermite.ActionRef{{Uses: "actions/checkout@v4"}, {Uses: "o/r/p@main"}}},
		&ermite.ResolveActionsResponse{Actions: []*ermite.ResolvedAction{
			{
				Ref: "o/r/p@main", NameWithOwner: "o/r", Path: "p", ResolvedSha: "0123456789012345678901234567890123456789",
				TarballUrl: "https://example.test/o/r.tar.gz", Token: "t", Source: ermite.ActionSource_ACTION_SOURCE_ERMITE,
			},
			{Ref: "x/y@z", Error: "Unable to resolve action `x/y@z`, unable to find version `z`"},
		}},
		&ermite.UploadSummaryRequest{JobId: "j1", StepNumber: 2, Markdown: "# hi"},
		&ermite.ListAnnotationsResponse{Annotations: []*ermite.JobAnnotation{{JobId: "j1", Annotation: &ermite.Annotation{Message: "m"}}}, Truncated: true},
		&ermite.GetStepSummariesResponse{Summaries: []*ermite.StepSummary{{StepNumber: 2, Markdown: "# hi"}}, Truncated: true},
		&ermite.RunJob{Id: "j1", Steps: []*ermite.Step{{Number: 2, Outcome: ermite.CiConclusion_CI_CONCLUSION_FAILURE}}},
	}
	for _, in := range msgs {
		b, err := proto.Marshal(in)
		if err != nil {
			t.Fatal(err)
		}
		out := in.ProtoReflect().New().Interface()
		if err := proto.Unmarshal(b, out); err != nil {
			t.Fatal(err)
		}
		if !proto.Equal(in, out) {
			t.Errorf("%T: round trip mismatch: %v != %v", in, out, in)
		}
	}
}

// TestCI3FieldNumbers pins the numbers of the fields added for CI-3.
func TestCI3FieldNumbers(t *testing.T) {
	tests := []struct {
		msg    proto.Message
		fields map[string]int
	}{
		{&ermite.Job{}, map[string]int{
			"secrets": 16, "vars": 17, "environment": 18, "permissions": 19, "cache": 20, "trusted": 21,
		}},
		{&ermite.JobPermissions{}, map[string]int{"levels": 1}},
		{&ermite.JobCache{}, map[string]int{"repository": 1, "write_ref": 2, "read_refs": 3}},
		{&ermite.Run{}, map[string]int{"trusted": 20}},
		{&ermite.RunJob{}, map[string]int{"concurrency_group": 20}},
		{&ermite.DispatchWorkflowRequest{}, map[string]int{"repo": 1, "workflow_path": 2, "ref": 3, "inputs": 4}},
		{&ermite.DispatchWorkflowResponse{}, map[string]int{"run": 1}},
		{&ermite.SettingsScope{}, map[string]int{"instance": 1, "repo": 2, "environment": 3}},
		{&ermite.EnvironmentRef{}, map[string]int{"repo": 1, "name": 2}},
		{&ermite.Secret{}, map[string]int{"name": 1, "updated_at": 2}},
		{&ermite.Variable{}, map[string]int{"name": 1, "value": 2, "updated_at": 3}},
		{&ermite.Environment{}, map[string]int{"repo": 1, "name": 2, "created_at": 3}},
		{&ermite.SetSecretRequest{}, map[string]int{"scope": 1, "name": 2, "value": 3}},
		{&ermite.ListSecretsRequest{}, map[string]int{"scope": 1, "page_size": 2, "page_token": 3}},
		{&ermite.ListSecretsResponse{}, map[string]int{"secrets": 1, "next_page_token": 2}},
		{&ermite.ListVariablesResponse{}, map[string]int{"variables": 1, "next_page_token": 2}},
		{&ermite.ListEnvironmentsRequest{}, map[string]int{"repo": 1, "page_size": 2, "page_token": 3}},
		{&ermite.ListEnvironmentsResponse{}, map[string]int{"environments": 1, "next_page_token": 2}},
	}
	for _, tt := range tests {
		md := tt.msg.ProtoReflect().Descriptor()
		for name, num := range tt.fields {
			if fd := md.Fields().ByName(protoreflect.Name(name)); fd == nil || int(fd.Number()) != num {
				t.Errorf("%s.%s: %v, want number %d", md.Name(), name, fd, num)
			}
		}
	}
}

// TestCI3RoundTrip checks that the CI-3 messages survive the wire.
func TestCI3RoundTrip(t *testing.T) {
	msgs := []proto.Message{
		&ermite.Job{
			JobId: "j1", Secrets: map[string]string{"TOKEN": "s3cret"}, Vars: map[string]string{"REGION": "eu"},
			Environment: "staging", Trusted: true,
			Permissions: &ermite.JobPermissions{Levels: map[string]string{"contents": "write", "packages": "none"}},
			Cache: &ermite.JobCache{
				Repository: "ermite", WriteRef: "refs/pull/7/merge",
				ReadRefs: []string{"refs/pull/7/merge", "refs/heads/main"},
			},
		},
		&ermite.Run{Id: 9, Event: "workflow_dispatch", Trusted: true},
		&ermite.RunJob{Id: "j1", Status: ermite.CiStatus_CI_STATUS_PENDING, ConcurrencyGroup: "deploy"},
		&ermite.DispatchWorkflowRequest{Repo: "ermite", WorkflowPath: "ci.yml", Ref: "main", Inputs: map[string]string{"debug": "true"}},
		&ermite.DispatchWorkflowResponse{Run: &ermite.Run{Id: 9, Event: "workflow_dispatch"}},
		&ermite.SetSecretRequest{Scope: &ermite.SettingsScope{Where: &ermite.SettingsScope_Instance{Instance: true}}, Name: "A", Value: "v"},
		&ermite.SetSecretRequest{Scope: &ermite.SettingsScope{Where: &ermite.SettingsScope_Repo{Repo: "ermite"}}, Name: "A", Value: "v"},
		&ermite.SetVariableRequest{
			Scope: &ermite.SettingsScope{Where: &ermite.SettingsScope_Environment{Environment: &ermite.EnvironmentRef{Repo: "ermite", Name: "prod"}}},
			Name:  "A", Value: "v",
		},
		&ermite.ListSecretsResponse{Secrets: []*ermite.Secret{{Name: "A"}}, NextPageToken: "t"},
		&ermite.ListEnvironmentsResponse{Environments: []*ermite.Environment{{Repo: "ermite", Name: "prod"}}},
	}
	for _, in := range msgs {
		b, err := proto.Marshal(in)
		if err != nil {
			t.Fatal(err)
		}
		out := in.ProtoReflect().New().Interface()
		if err := proto.Unmarshal(b, out); err != nil {
			t.Fatal(err)
		}
		if !proto.Equal(in, out) {
			t.Errorf("%T: round trip mismatch: %v != %v", in, out, in)
		}
	}
}
