package docker_server

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// From
func (b *BuilderServer) From(ctx context.Context, r *v1.FromRequest) (*v1.FromResponse, error) {
	// generate a new container ID
	h := sha256.New()
	h.Write([]byte(r.Image))
	sum := h.Sum(nil)

	id := hex.EncodeToString(sum)

	// Create a new container state
	cs := &containerStateImpl{
		name:      id,
		dependsOn: make([]string, 0),
		lines:     make([]string, 0),
	}

	// The user has provided another container state as a dependency
	fromImage := r.Image
	if b.store.Has(r.Image) {
		// The extension is another container state dependency
		cs.AddDependency(r.Image)
		fromImage = fmt.Sprintf("layer-%s", r.Image)
	}

	// Add line to container state
	cs.AddLine(fmt.Sprintf("FROM %s as layer-%s", fromImage, id))

	// Add to central container state store
	if err := b.store.Put(id, cs); err != nil {
		return nil, status.Error(codes.FailedPrecondition, "Failed to add container to state store")
	}

	// Return references to the newly created container state
	return &v1.FromResponse{
		Container: &v1.Container{
			Id: id,
		},
	}, nil
}
