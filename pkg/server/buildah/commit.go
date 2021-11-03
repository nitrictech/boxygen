package buildah_server

import (
	"context"

	"github.com/containers/buildah"
	"github.com/containers/image/v5/transports/alltransports"
	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (b *BuilderServer) Commit(ctx context.Context, r *v1.CommitRequest) (*v1.CommitResponse, error) {
	// Rehydrate from an existing builder
	builder, err := buildah.OpenBuilder(nil, r.Container.GetId())

	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	// Commit to the specified location
	// To commit to the local docker daemon use docker-daemon:Tag
	dest, err := alltransports.ParseImageName(r.Tag)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, _, _, err = builder.Commit(ctx, dest, buildah.CommitOptions{})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.CommitResponse{}, nil
}
