// SPDX-License-Identifier: MIT

package api_test

import (
	"testing"

	"google.golang.org/protobuf/proto"
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
		{ermiteconnect.GitServiceGetBlobProcedure, "/ermite.GitService/GetBlob"},
		{ermiteconnect.GitServiceListCommitsProcedure, "/ermite.GitService/ListCommits"},
		{ermiteconnect.GitServiceGetCommitProcedure, "/ermite.GitService/GetCommit"},
		{ermiteconnect.GitServiceResolveRevisionProcedure, "/ermite.GitService/ResolveRevision"},
		{ermiteconnect.TokenServiceCreateTokenProcedure, "/ermite.TokenService/CreateToken"},
		{ermiteconnect.TokenServiceListTokensProcedure, "/ermite.TokenService/ListTokens"},
		{ermiteconnect.TokenServiceRevokeTokenProcedure, "/ermite.TokenService/RevokeToken"},
		{ermiteconnect.EventServiceWatchProcedure, "/ermite.EventService/Watch"},
		{ermiteconnect.ServerServiceGetVersionProcedure, "/ermite.ServerService/GetVersion"},
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
