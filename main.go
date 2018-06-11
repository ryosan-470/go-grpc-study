package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "github.com/ryosan-470/go-grpc-study/proto"
)

var (
	endpoint = flag.String("alive_endpoint", "0.0.0.0:9090", "endpoint of aliveService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterAliveServiceHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":8080", mux)
}
func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
