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
	"path"

	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Copy
func (b *BuilderServer) Copy(r *v1.CopyRequest, srv v1.Builder_CopyServer) error {
	cs, err := b.store.Get(r.Container.Id)

	if err != nil {
		return status.Errorf(codes.NotFound, "container: %s does not exist", r.Container.Id)
	}

	if r.From != "" {
		// add container state dependency as well
		if !b.store.Has(r.From) {
			return status.Errorf(codes.NotFound, "container %s does not exist", r.From)
		}

		cs.AddDependency(r.From)

		appendAndLog(fmt.Sprintf("COPY --from layer-%s %s %s", r.From, r.Source, r.Dest), cs, srv)
	} else {
		// Workspace COPY
		appendAndLog(fmt.Sprintf("COPY %s %s", path.Join(b.workspace, r.Source), r.Dest), cs, srv)
	}

	return nil
}
