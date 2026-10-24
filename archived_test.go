// SPDX-License-Identifier: MIT

package api_test

import (
	"bnema.dev/ermite-api/gen/ermite"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestArchivedPresence(t *testing.T) {
	for _, value := range []*bool{nil, proto.Bool(false), proto.Bool(true)} {
		original := &ermite.UpdateRepoRequest{Repo: "demo", Archived: value}
		wire, err := proto.Marshal(original)
		if err != nil {
			t.Fatal(err)
		}
		decoded := new(ermite.UpdateRepoRequest)
		if err := proto.Unmarshal(wire, decoded); err != nil {
			t.Fatal(err)
		}
		if !proto.Equal(original, decoded) {
			t.Fatalf("presence lost: %v", decoded)
		}
	}
}
