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
	"github.com/nitrictech/boxygen/pkg/backend/dockerfile"
	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Add
func (b *BuilderServer) Run(r *v1.RunRequest, srv v1.Builder_RunServer) error {
	cs, err := b.store.Get(r.Container.Id)

	if err != nil {
		return status.Errorf(codes.NotFound, "container: %s does not exist", r.Container.Id)
	}

	cs.Run(dockerfile.RunOptions{
		Command: r.Command,
	})

	// Append line to the container context
	return nil
}
