package buildah_server

import (
	"context"

	"github.com/nitrictech/boxygen/pkg/common"
	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (b *BuilderServer) Config(ctx context.Context, r *v1.ConfigRequest) (*v1.ConfigResponse, error) {
	// Rehydrate from an existing builder
	//builder, err := buildah.OpenBuilder(nil, r.Container.GetId())

	builder, err := common.BuilderForContainer(r.Container)

	for k, v := range r.Env {
		builder.SetEnv(k, v)
	}

	// Set the command
	builder.SetCmd(r.Cmd)
	// Set the entrpoint
	builder.SetEntrypoint(r.Entrypoint)
	// Set the working directory
	builder.SetWorkDir(r.WorkingDir)

	//if err != nil {
	//	return nil, status.Error(codes.NotFound, err.Error())
	//}

	// Save the builders
	builder.Save()

	// Now we can load and update the containers config

	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}
