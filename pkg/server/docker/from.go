// Copyright 2021 Nitric Technologies Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docker_server

import (
	"context"

	"github.com/nitrictech/boxygen/pkg/backend/dockerfile"
	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// From
func (b *BuilderServer) From(ctx context.Context, r *v1.FromRequest) (*v1.FromResponse, error) {
	con, err := b.store.NewContainer(dockerfile.NewContainerOpts{
		From:   r.Image,
		As:     r.As,
		Ignore: r.Ignore,
	})

	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "Failed to add container to state store")
	}

	// Return references to the newly created container state
	return &v1.FromResponse{
		Container: &v1.Container{
			Id: con.Name(),
		},
	}, nil
}
