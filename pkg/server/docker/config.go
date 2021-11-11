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
	"fmt"
	"strings"

	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Add
func (b *BuilderServer) Config(r *v1.ConfigRequest, srv v1.Builder_ConfigServer) error {
	cs, err := b.store.Get(r.Container.Id)

	if err != nil {
		return status.Errorf(codes.NotFound, "container: %s does not exist", r.Container.Id)
	}

	if r.WorkingDir != "" {
		appendAndLog(fmt.Sprintf("WORKDIR %s", r.WorkingDir), cs, srv)
	}

	if r.User != "" {
		appendAndLog(fmt.Sprintf("USER %s", r.User), cs, srv)
	}

	for _, volume := range r.Volumes {
		appendAndLog(fmt.Sprintf("VOLUME %s", volume), cs, srv)
	}

	for _, port := range r.Ports {
		appendAndLog(fmt.Sprintf("EXPOSE %d", port), cs, srv)
	}

	for k, v := range r.Env {
		appendAndLog(fmt.Sprintf("ENV %s=%s", k, v), cs, srv)
	}

	if len(r.Entrypoint) > 0 {
		appendAndLog(fmt.Sprintf("ENTRYPOINT [\"%s\"]", strings.Join(r.Entrypoint, "\", \"")), cs, srv)
	}

	if len(r.Cmd) > 0 {
		appendAndLog(fmt.Sprintf("CMD [\"%s\"]", strings.Join(r.Cmd, "\", \"")), cs, srv)
	}

	return nil
}
