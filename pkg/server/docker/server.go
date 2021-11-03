package docker_server

import (
	pb "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
)

type BuilderServer struct {
	store ContainerStateStore
	pb.UnimplementedBuilderServer
}

func New() pb.BuilderServer {
	return &BuilderServer{
		store: &containerStateStoreImpl{
			store: make(map[string]ContainerState),
		},
	}
}
