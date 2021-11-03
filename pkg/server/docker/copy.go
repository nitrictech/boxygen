package docker_server

import (
	"context"
	"fmt"

	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Copy
func (b *BuilderServer) Copy(ctx context.Context, r *v1.CopyRequest) (*v1.CopyResponse, error) {
	cs, err := b.store.Get(r.Container.Id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "container: %s does not exist", r.Container.Id)
	}

	if r.From != "" {
		// add container state dependency as well
		cs.AddDependency(r.From)
		cs.AddLine(fmt.Sprintf("COPY --from %s %s %s", r.From, r.Source, r.Dest))
	} else {
		cs.AddLine(fmt.Sprintf("COPY %s %s", r.Source, r.Dest))
	}

	return &v1.CopyResponse{}, nil
}
