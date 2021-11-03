package docker_server

import (
	"context"
	"fmt"
	"strings"

	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Add
func (b *BuilderServer) Run(ctx context.Context, r *v1.RunRequest) (*v1.RunResponse, error) {
	cs, err := b.store.Get(r.Container.Id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "container: %s does not exist", r.Container.Id)
	}

	// Load the container file state to append to ready to commit
	cs.AddLine(fmt.Sprintf("RUN %s", strings.Join(r.GetCommand(), " ")))

	// Append line to the container context
	return &v1.RunResponse{}, nil
}
