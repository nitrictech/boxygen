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
func (b *BuilderServer) Config(ctx context.Context, r *v1.ConfigRequest) (*v1.ConfigResponse, error) {
	cs, err := b.store.Get(r.Container.Id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "container: %s does not exist", r.Container.Id)
	}

	if r.User != "" {
		cs.AddLine(fmt.Sprintf("USER %s", r.User))
	}

	for _, volume := range r.Volumes {
		cs.AddLine(fmt.Sprintf("VOLUME %s", volume))
	}

	for _, port := range r.Ports {
		cs.AddLine(fmt.Sprintf("EXPOSE %d", port))
	}

	for k, v := range r.Env {
		cs.AddLine(fmt.Sprintf("ENV %s=%s", k, v))
	}

	if len(r.Entrypoint) > 0 {
		cs.AddLine(fmt.Sprintf("ENTRYPOINT [\"%s\"]", strings.Join(r.Entrypoint, "\", \"")))
	}

	if len(r.Cmd) > 0 {
		cs.AddLine(fmt.Sprintf("CMD [\"%s\"]", strings.Join(r.Cmd, "\", \"")))
	}

	return &v1.ConfigResponse{}, nil
}
