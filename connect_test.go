// SPDX-License-Identifier: MIT

package api_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"bnema.dev/ermite-api/gen/ermite"
	"bnema.dev/ermite-api/gen/ermite/ermiteconnect"
	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connecthttp"
)

type migrationServer struct {
	ermiteconnect.UnimplementedServerServiceHandler
	ermiteconnect.UnimplementedRunnerServiceHandler
}

func (migrationServer) GetVersion(context.Context, *ermite.GetVersionRequest) (*ermite.GetVersionResponse, error) {
	return &ermite.GetVersionResponse{Version: "test"}, nil
}

func (migrationServer) WatchJob(_ context.Context, req *ermite.WatchJobRequest, stream ermiteconnect.RunnerServiceWatchJobServerStream) error {
	return stream.Send(&ermite.WatchJobResponse{Msg: &ermite.WatchJobResponse_Ended{Ended: &ermite.JobEnded{Error: req.JobId}}})
}

// Exercise the generated v2 registration, transport, unwrapped unary API and
// per-RPC stream types over the same HTTP prefix used by downstream servers.
func TestConnectV2HTTP(t *testing.T) {
	server := connect.NewServer()
	ermiteconnect.RegisterServerServiceHandler(server, migrationServer{})
	ermiteconnect.RegisterRunnerServiceHandler(server, migrationServer{})
	rpcMux := http.NewServeMux()
	connecthttp.Mount(rpcMux, server)
	mux := http.NewServeMux()
	mux.Handle("/-/api/", http.StripPrefix("/-/api", rpcMux))
	httpServer := httptest.NewServer(mux)
	t.Cleanup(httpServer.Close)
	client := connect.NewClient(connecthttp.NewTransport(httpServer.Client(), httpServer.URL+"/-/api"))
	res, err := ermiteconnect.NewServerServiceClient(client).GetVersion(t.Context(), &ermite.GetVersionRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if res.Version != "test" {
		t.Fatalf("version = %q", res.Version)
	}
	stream, err := ermiteconnect.NewRunnerServiceClient(client).WatchJob(t.Context(), &ermite.WatchJobRequest{JobId: "job-1"})
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()
	msg, err := stream.Receive()
	if err != nil {
		t.Fatal(err)
	}
	if msg.GetEnded().GetError() != "job-1" {
		t.Fatalf("unexpected message: %v", msg)
	}
	if _, err := stream.Receive(); !errors.Is(err, io.EOF) {
		t.Fatalf("stream end = %v, want EOF", err)
	}
}
