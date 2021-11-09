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
