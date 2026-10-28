// SPDX-License-Identifier: MIT

package api_test

import (
	"bnema.dev/ermite-api/gen/ermite"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestRepoLanguagesRoundTrip(t *testing.T) {
	original := &ermite.Repo{Name: "demo", Languages: []*ermite.RepoLanguage{
		{Name: "Go", Bytes: 3072, Percent: 75},
		{Name: "Shell", Bytes: 1024, Percent: 25},
	}}
	wire, err := proto.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}
	decoded := new(ermite.Repo)
	if err := proto.Unmarshal(wire, decoded); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(original, decoded) {
		t.Fatalf("languages lost: %v", decoded)
	}
	if got := decoded.GetLanguages()[1]; got.GetName() != "Shell" || got.GetBytes() != 1024 || got.GetPercent() != 25 {
		t.Fatalf("unexpected language: %v", got)
	}
}
