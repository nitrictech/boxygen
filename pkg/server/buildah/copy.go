package buildah_server

import (
	"context"

	"github.com/nitrictech/boxygen/pkg/common"
	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Add
func (b *BuilderServer) Copy(ctx context.Context, r *v1.CopyRequest) (*v1.CopyResponse, error) {
	builder, err := common.BuilderForContainer(r.Container)

	// Add to the root filesystem
	builder.Add()

	return nil, status.Error(codes.Unimplemented, "Unimplmented")
}
