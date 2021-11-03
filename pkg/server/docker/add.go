package docker_server

import (
	"context"
	"fmt"

	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Add
func (b *BuilderServer) Add(ctx context.Context, r *v1.AddRequest) (*v1.AddResponse, error) {
	cs, err := b.store.Get(r.Container.Id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "container: %s does not exist", r.Container.Id)
	}

	cs.AddLine(fmt.Sprintf("ADD %s %s", r.Src, r.Dest))

	return &v1.AddResponse{}, nil
}
