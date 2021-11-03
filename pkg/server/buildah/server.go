package buildah_server

import (
	pb "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
)

type BuilderServer struct {
	pb.UnimplementedBuilderServer
}

func New() pb.BuilderServer {
	return &BuilderServer{}
}
