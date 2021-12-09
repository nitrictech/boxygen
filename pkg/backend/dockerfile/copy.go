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

package dockerfile

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CopyOptions struct {
	Src  string
	Dest string
	From string
}

func (c *containerStateImpl) Copy(opts CopyOptions) error {

	if opts.From != "" {
		// add container state dependency as well
		if !c.store.Has(opts.From) {
			return status.Errorf(codes.NotFound, "container %s does not exist", opts.From)
		}

		c.addDependency(opts.From)
		c.addLine(fmt.Sprintf("COPY --from=layer-%s %s %s", opts.From, opts.Src, opts.Dest))
	} else {
		// Workspace COPY
		c.addLine(fmt.Sprintf("COPY %s %s", opts.Src, opts.Dest))
	}

	return nil
}
