// SPDX-License-Identifier: MIT

package api_test

import (
	"testing"

	"bnema.dev/ermite-api/gen/ermite"
	"bnema.dev/ermite-api/gen/ermite/ermiteconnect"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestIssueProcedures(t *testing.T) {
	tests := map[string]string{
		ermiteconnect.IssueServiceCreateIssueProcedure:       "CreateIssue",
		ermiteconnect.IssueServiceListIssuesProcedure:        "ListIssues",
		ermiteconnect.IssueServiceGetIssueProcedure:          "GetIssue",
		ermiteconnect.IssueServiceAddIssueCommentProcedure:   "AddIssueComment",
		ermiteconnect.IssueServiceCloseIssueProcedure:        "CloseIssue",
		ermiteconnect.IssueServiceReopenIssueProcedure:       "ReopenIssue",
		ermiteconnect.IssueServiceResolveIssueProcedure:      "ResolveIssue",
		ermiteconnect.IssueServiceListPendingIssuesProcedure: "ListPendingIssues",
		ermiteconnect.IssueServiceApproveIssueProcedure:      "ApproveIssue",
		ermiteconnect.IssueServiceRejectIssueProcedure:       "RejectIssue",
	}
	for got, name := range tests {
		if want := "/ermite.IssueService/" + name; got != want {
			t.Errorf("procedure = %q, want %q", got, want)
		}
	}
}

func TestIssueContract(t *testing.T) {
	for _, tt := range []struct {
		msg    proto.Message
		fields map[string]int
	}{
		{&ermite.Issue{}, map[string]int{"repo": 1, "number": 2, "title": 3, "body": 4, "state": 5, "author": 6, "created_at": 7, "updated_at": 8, "comment_count": 9}},
		{&ermite.IssueComment{}, map[string]int{"id": 1, "author": 2, "text": 3, "created_at": 4}},
		{&ermite.RepoSettings{}, map[string]int{"issues": 3}},
		{&ermite.UpdateRepoRequest{}, map[string]int{"issues": 8}},
		{&ermite.Event{}, map[string]int{"issue_opened": 17, "issue_closed": 18, "issue_reopened": 19, "issue_commented": 20, "issue_submitted": 21}},
		{&ermite.IssueSubmitted{}, map[string]int{"repo": 1, "id": 2}},
		{&ermite.PendingIssue{}, map[string]int{"id": 1, "repo": 2, "title": 3, "body": 4, "author": 5, "created_at": 6}},
		{&ermite.ListPendingIssuesRequest{}, map[string]int{"repo": 1, "page_size": 2, "page_token": 3}},
		{&ermite.ListPendingIssuesResponse{}, map[string]int{"issues": 1, "next_page_token": 2}},
		{&ermite.ApproveIssueRequest{}, map[string]int{"repo": 1, "id": 2}},
		{&ermite.ApproveIssueResponse{}, map[string]int{"issue": 1}},
		{&ermite.RejectIssueRequest{}, map[string]int{"repo": 1, "id": 2, "ban": 3}},
	} {
		md := tt.msg.ProtoReflect().Descriptor()
		for name, num := range tt.fields {
			if fd := md.Fields().ByName(protoreflect.Name(name)); fd == nil || int(fd.Number()) != num {
				t.Errorf("%s.%s = %v, want %d", md.Name(), name, fd, num)
			}
		}
	}
	if ermite.Scope_SCOPE_ISSUE_WRITE != 11 {
		t.Fatal("issue:write scope number changed")
	}
	for name, want := range map[string]int32{"ISSUE_STATE_UNSPECIFIED": 0, "ISSUE_STATE_OPEN": 1, "ISSUE_STATE_CLOSED": 2, "ISSUE_STATE_PENDING": 3, "ISSUE_STATE_RESOLVED": 4} {
		if got := ermite.IssueState_value[name]; got != want {
			t.Errorf("%s = %d, want %d", name, got, want)
		}
	}
	for name, want := range map[string]int32{"ISSUES_MODE_UNSPECIFIED": 0, "ISSUES_MODE_OFF": 1, "ISSUES_MODE_OWNER": 2, "ISSUES_MODE_PUBLIC": 3} {
		if got := ermite.IssuesMode_value[name]; got != want {
			t.Errorf("%s = %d, want %d", name, got, want)
		}
	}
	mode := ermite.IssuesMode_ISSUES_MODE_OFF
	for _, in := range []proto.Message{
		&ermite.UpdateRepoRequest{Repo: "r", Issues: &mode},
		&ermite.Event{Kind: &ermite.Event_IssueOpened{IssueOpened: &ermite.IssueOpened{Issue: &ermite.Issue{Repo: "r", Number: 1}}}},
		&ermite.Event{Kind: &ermite.Event_IssueClosed{IssueClosed: &ermite.IssueClosed{Issue: &ermite.Issue{Repo: "r", Number: 1}}}},
		&ermite.Event{Kind: &ermite.Event_IssueReopened{IssueReopened: &ermite.IssueReopened{Issue: &ermite.Issue{Repo: "r", Number: 1}}}},
		&ermite.Event{Kind: &ermite.Event_IssueSubmitted{IssueSubmitted: &ermite.IssueSubmitted{Repo: "r", Id: "p1"}}},
		&ermite.PendingIssue{Id: "p1", Repo: "r", Title: "t", Author: "anon"},
		&ermite.RejectIssueRequest{Repo: "r", Id: "p1", Ban: true},
		&ermite.Event{Kind: &ermite.Event_IssueCommented{IssueCommented: &ermite.IssueCommented{Repo: "r", Number: 1, Comment: &ermite.IssueComment{Id: "c", Text: "hi"}}}},
	} {
		b, err := proto.Marshal(in)
		if err != nil {
			t.Fatal(err)
		}
		out := in.ProtoReflect().New().Interface()
		if err := proto.Unmarshal(b, out); err != nil {
			t.Fatal(err)
		}
		if !proto.Equal(in, out) {
			t.Errorf("%T: round trip mismatch", in)
		}
	}
	fd := (&ermite.UpdateRepoRequest{}).ProtoReflect().Descriptor().Fields().ByName("issues")
	if !fd.HasOptionalKeyword() {
		t.Fatal("UpdateRepoRequest.issues must be optional")
	}
}
