// SPDX-License-Identifier: MIT

package api_test

import (
	"testing"

	"bnema.dev/ermite-api/gen/ermite"
	"bnema.dev/ermite-api/gen/ermite/ermiteconnect"
	"google.golang.org/protobuf/proto"
)

func TestEditProcedures(t *testing.T) {
	for _, tc := range []struct{ got, want string }{
		{ermiteconnect.IssueServiceEditIssueProcedure, "/ermite.IssueService/EditIssue"},
		{ermiteconnect.MergeRequestServiceEditMergeRequestProcedure, "/ermite.MergeRequestService/EditMergeRequest"},
	} {
		if tc.got != tc.want {
			t.Errorf("procedure = %q, want %q", tc.got, tc.want)
		}
	}
}

// Empty replacement text must remain distinguishable from an omitted field.
func TestEditOptionalTextRoundTrip(t *testing.T) {
	for _, message := range []proto.Message{
		&ermite.UpdateRepoRequest{Repo: "demo"},
		&ermite.UpdateRepoRequest{Repo: "demo", CiAutoCancel: proto.Bool(false)},
		&ermite.UpdateRepoRequest{Repo: "demo", CiAutoCancel: proto.Bool(true)},
		&ermite.EditIssueRequest{Repo: "demo", Number: 1},
		&ermite.EditIssueRequest{Repo: "demo", Number: 1, Title: proto.String(""), Body: proto.String("")},
		&ermite.EditIssueRequest{Repo: "demo", Number: 1, Title: proto.String("new")},
		&ermite.EditIssueRequest{Repo: "demo", Number: 1, Body: proto.String("new")},
		&ermite.EditMergeRequestRequest{Repo: "demo", Number: 1},
		&ermite.EditMergeRequestRequest{Repo: "demo", Number: 1, Title: proto.String(""), Description: proto.String("")},
		&ermite.EditMergeRequestRequest{Repo: "demo", Number: 1, Title: proto.String("new")},
		&ermite.EditMergeRequestRequest{Repo: "demo", Number: 1, Description: proto.String("new")},
	} {
		wire, err := proto.Marshal(message)
		if err != nil {
			t.Fatal(err)
		}
		decoded := message.ProtoReflect().New().Interface()
		if err := proto.Unmarshal(wire, decoded); err != nil {
			t.Fatal(err)
		}
		if !proto.Equal(message, decoded) {
			t.Errorf("presence lost: %v -> %v", message, decoded)
		}
	}
}
