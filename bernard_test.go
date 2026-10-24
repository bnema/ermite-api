// SPDX-License-Identifier: MIT

package api_test

import (
	"testing"

	"bnema.dev/ermite-api/gen/ermite"
	"bnema.dev/ermite-api/gen/ermite/ermiteconnect"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestBernardFieldNumbers(t *testing.T) {
	for _, tc := range []struct {
		message proto.Message
		field   protoreflect.Name
		number  protoreflect.FieldNumber
	}{
		{&ermite.Repo{}, "archived", 9},
		{&ermite.UpdateRepoRequest{}, "archived", 9},
		{&ermite.Event{}, "dependency_updates_found", 34},
		{&ermite.FetchJobRequest{}, "kinds", 4},
		{&ermite.Job{}, "deps", 24},
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
		&ermite.OpenDependencyUpdatesRequest{Repo: "demo", ExpectedVersions: []*ermite.DependencyVersion{{Ecosystem: "gomod", Directory: "sub", Name: "example.com/mod", Version: "v1.2.3"}}},
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
