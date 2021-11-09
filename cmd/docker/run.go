// Copyright 2021 Nitric Technologies Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nitrictech/boxygen/pkg/common"
	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	docker_server "github.com/nitrictech/boxygen/pkg/server/docker"
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
	// Run the server
	srv := grpc.NewServer()

	v1.RegisterBuilderServer(srv, docker_server.New())

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
