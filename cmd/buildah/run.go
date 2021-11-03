package main

import (
	"fmt"
	"log"
	"net"

	"github.com/containers/buildah"
	"github.com/nitrictech/boxygen/pkg/common"
	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	buildah_server "github.com/nitrictech/boxygen/pkg/server/buildah"
	"google.golang.org/grpc"
)

const (
	defaultPort = "50051"
)

type envCtx struct {
	port string
}

func ctxFromEnv() envCtx {
	ctx := envCtx{}

	ctx.port = common.GetEnv("BOXYGEN_PORT", defaultPort)

	return ctx
}

func main() {
	if buildah.InitReexec() {
		return
	}

	// Run the server
	srv := grpc.NewServer()

	v1.RegisterBuilderServer(srv, buildah_server.New())

	ctx := ctxFromEnv()

	// TODO: Make port configurable
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", ctx.port))

	if err != nil {
		// server already started just return (due to use of reexec in buildah)
		log.Fatalf("Server already started!")
	}

	fmt.Printf("Listening on %s\n", ctx.port)

	// Start the boxygen builder server
	err = srv.Serve(lis)

	log.Fatalf(err.Error())
}
