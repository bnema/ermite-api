// SPDX-License-Identifier: MIT

package api_test

import (
	"testing"

	"bnema.dev/ermite-api/gen/ermite"
	"bnema.dev/ermite-api/gen/ermite/ermiteconnect"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestBernardFieldNumbers(t *testing.T) {
	for _, tc := range []struct {
		message proto.Message
		field   protoreflect.Name
		number  protoreflect.FieldNumber
	}{
		{&ermite.RepoDependencyUpdates{}, "auto_open", 10},
		{&ermite.RepoDependencyUpdates{}, "split_out", 11},
		{&ermite.DependencyBatch{}, "adopted", 8},
		{&ermite.DependencyBatch{}, "branch", 9},
		{&ermite.DependencyBatch{}, "title", 10},
		{&ermite.DependencyBatch{}, "queue_position", 11},
		{&ermite.DependencyBatch{}, "repro", 12},
		{&ermite.DependencyName{}, "ecosystem", 1},
		{&ermite.DependencyName{}, "name", 2},
		{&ermite.DependencyQueue{}, "running", 1},
		{&ermite.DependencyQueue{}, "limit", 2},
		{&ermite.DependencyQueue{}, "queued", 3},
		{&ermite.ListDependencyUpdatesResponse{}, "queue", 2},
		{&ermite.SetDependencyAutoOpenRequest{}, "repo", 1},
		{&ermite.SetDependencyAutoOpenRequest{}, "enabled", 2},
		{&ermite.SetDependencyAutoOpenResponse{}, "updates", 1},
		{&ermite.SplitDependencyRequest{}, "repo", 1},
		{&ermite.SplitDependencyRequest{}, "ecosystem", 2},
		{&ermite.SplitDependencyRequest{}, "name", 3},
		{&ermite.SplitDependencyRequest{}, "split", 4},
		{&ermite.SplitDependencyResponse{}, "updates", 1},
		{&ermite.Repo{}, "archived", 9},
		{&ermite.UpdateRepoRequest{}, "archived", 9},
		{&ermite.Event{}, "dependency_updates_found", 34},
		{&ermite.FetchJobRequest{}, "kinds", 4},
		{&ermite.Job{}, "deps", 24},
		{&ermite.DependencyUpdate{}, "new_name", 8},
		{&ermite.RepoDependencyUpdates{}, "batches", 9},
		{&ermite.OpenDependencyUpdatesRequest{}, "batch", 3},
		{&ermite.IgnoreDependencyUpdatesRequest{}, "batch", 3},
		{&ermite.DepsJob{}, "node_version", 3},
		{&ermite.DependencyBatch{}, "key", 1},
		{&ermite.DependencyBatch{}, "updates", 2},
		{&ermite.DependencyBatch{}, "state", 3},
		{&ermite.DependencyBatch{}, "error", 4},
		{&ermite.DependencyBatch{}, "run_id", 5},
		{&ermite.DependencyBatch{}, "merge_request_number", 6},
		{&ermite.DependencyBatch{}, "outdated", 7},
		{&ermite.DepsJobUpdate{}, "new_name", 5},
		{&ermite.DepsJobUpdate{}, "range", 6},
		{&ermite.DepsJobUpdate{}, "manager", 7},
		{&ermite.DepsJobUpdate{}, "manager_version", 8},
		{&ermite.DepsJobUpdate{}, "root", 9},
		{&ermite.CompleteJobRequest{}, "deps_report", 7},
	} {
		field := tc.message.ProtoReflect().Descriptor().Fields().ByName(tc.field)
		if field == nil || field.Number() != tc.number {
			t.Errorf("%T.%s: want field %d", tc.message, tc.field, tc.number)
		}
	}
}

func TestDepsProcedures(t *testing.T) {
	for _, tc := range []struct{ got, method string }{
		{ermiteconnect.DepsServiceSetDependencyAutoOpenProcedure, "SetDependencyAutoOpen"},
		{ermiteconnect.DepsServiceSplitDependencyProcedure, "SplitDependency"},
		{ermiteconnect.DepsServiceListDependencyUpdatesProcedure, "ListDependencyUpdates"},
		{ermiteconnect.DepsServiceCheckDependenciesProcedure, "CheckDependencies"},
		{ermiteconnect.DepsServiceOpenDependencyUpdatesProcedure, "OpenDependencyUpdates"},
		{ermiteconnect.DepsServiceIgnoreDependencyUpdatesProcedure, "IgnoreDependencyUpdates"},
		{ermiteconnect.DepsServiceSnoozeDependencyUpdatesProcedure, "SnoozeDependencyUpdates"},
	} {
		if tc.got != "/ermite.DepsService/"+tc.method {
			t.Errorf("unexpected procedure %q", tc.got)
		}
	}
}

func TestDepsWireRoundTrip(t *testing.T) {
	for _, original := range []proto.Message{
		&ermite.SetDependencyAutoOpenRequest{Repo: "demo", Enabled: true},
		&ermite.SplitDependencyRequest{Repo: "demo", Ecosystem: "npm", Name: "@scope/pkg", Split: true},
		&ermite.SplitDependencyRequest{Repo: "demo", Ecosystem: "npm", Name: "@scope/pkg", Split: false},
		&ermite.SetDependencyAutoOpenResponse{Updates: &ermite.RepoDependencyUpdates{AutoOpen: true}},
		&ermite.SplitDependencyResponse{Updates: &ermite.RepoDependencyUpdates{SplitOut: []*ermite.DependencyName{{Ecosystem: "npm", Name: "@scope/pkg"}}}},
		&ermite.ListDependencyUpdatesResponse{Queue: &ermite.DependencyQueue{Running: 1, Limit: 2, Queued: 3}},
		&ermite.DependencyBatch{Adopted: true, Branch: "bernard/npm", Title: "@scope", QueuePosition: 1, Repro: []string{"npm ci", "npm test"}, State: ermite.RepoDependencyUpdates_STATE_DRAFTED},
		&ermite.DependencyBatch{State: ermite.RepoDependencyUpdates_STATE_QUEUED},
		&ermite.IgnoreDependencyUpdatesRequest{Repo: "demo", Batch: "npm/web/react@v19"},
		&ermite.RepoDependencyUpdates{Batches: []*ermite.DependencyBatch{{Key: "gomod/sub/mod@v2", Updates: []*ermite.DependencyUpdate{{NewName: "example.com/mod/v2"}}, State: ermite.RepoDependencyUpdates_STATE_FAILED, Error: "failed", RunId: 7, MergeRequestNumber: 8, Outdated: true}}},
		&ermite.DepsJob{NodeVersion: "22", Updates: []*ermite.DepsJobUpdate{{Ecosystem: "npm", Name: "react", Version: "19.0.0", Range: "^19.0.0", Manager: "pnpm", ManagerVersion: "9.15.0", Root: "web", NewName: "new-name"}}},
		&ermite.OpenDependencyUpdatesRequest{Repo: "demo", Batch: "gomod/sub/mod@v2", ExpectedVersions: []*ermite.DependencyVersion{{Ecosystem: "gomod", Directory: "sub", Name: "example.com/mod", Version: "v1.2.3"}}},
		&ermite.Job{JobId: "deps-example", Deps: &ermite.DepsJob{GoVersion: "1.26", Updates: []*ermite.DepsJobUpdate{{Ecosystem: "gomod", Name: "example.com/mod", Version: "v1.2.3"}}}},
		&ermite.CompleteJobRequest{JobId: "deps-example", DepsReport: &ermite.DepsReport{Files: []*ermite.DepsFile{{Path: "go.mod", Content: []byte("module example.com/demo\n")}}}},
		&ermite.Event{Kind: &ermite.Event_DependencyUpdatesFound{DependencyUpdatesFound: &ermite.DependencyUpdatesFound{Repo: "demo", Count: 2}}},
	} {
		wire, err := proto.Marshal(original)
		if err != nil {
			t.Fatal(err)
		}
		decoded := original.ProtoReflect().New().Interface()
		if err := proto.Unmarshal(wire, decoded); err != nil {
			t.Fatal(err)
		}
		if !proto.Equal(original, decoded) {
			t.Errorf("round trip changed %T", original)
		}
	}
}

func TestDepsLegacyFieldsDeprecated(t *testing.T) {
	fields := (&ermite.RepoDependencyUpdates{}).ProtoReflect().Descriptor().Fields()
	for _, number := range []protoreflect.FieldNumber{2, 5, 6, 7, 8} {
		field := fields.ByNumber(number)
		if field == nil || !field.Options().(*descriptorpb.FieldOptions).GetDeprecated() {
			t.Errorf("RepoDependencyUpdates field %d must remain deprecated", number)
		}
	}
}

func TestDepsQueueStateNumbers(t *testing.T) {
	if ermite.RepoDependencyUpdates_STATE_QUEUED != 4 {
		t.Error("STATE_QUEUED must remain 4")
	}
	if ermite.RepoDependencyUpdates_STATE_DRAFTED != 5 {
		t.Error("STATE_DRAFTED must remain 5")
	}
}
