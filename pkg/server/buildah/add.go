package buildah_server

import (
	"context"

	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Add
func (b *BuilderServer) Add(ctx context.Context, r *v1.AddRequest) (*v1.AddResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplmented")
}
