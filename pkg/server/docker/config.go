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
func (b *BuilderServer) Config(ctx context.Context, r *v1.ConfigRequest) (*v1.ConfigResponse, error) {
	cs, err := b.store.Get(r.Container.Id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "container: %s does not exist", r.Container.Id)
	}

	for k, v := range r.Env {
		cs.AddLine(fmt.Sprintf("ENV %s=%s", k, v))
	}

	if len(r.Entrypoint) > 0 {
		cs.AddLine(fmt.Sprintf("ENTRYPOINT %s", strings.Join(r.Entrypoint, " ")))
	}

	if len(r.Cmd) > 0 {
		cs.AddLine(fmt.Sprintf("CMD %s", strings.Join(r.Cmd, " ")))
	}

	return &v1.ConfigResponse{}, nil
}
