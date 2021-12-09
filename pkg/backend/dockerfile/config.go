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
	"strings"
)

type ConfigOptions struct {
	WorkingDir string
	Cmd        []string
	User       string
	Volumes    []string
	Ports      []int32
	Env        map[string]string
	Entrypoint []string
}

func (c *containerStateImpl) Config(opts ConfigOptions) {
	if opts.WorkingDir != "" {
		c.addLine(fmt.Sprintf("WORKDIR %s", opts.WorkingDir))
	}

	if opts.User != "" {
		c.addLine(fmt.Sprintf("USER %s", opts.User))
	}

	for _, volume := range opts.Volumes {
		c.addLine(fmt.Sprintf("VOLUME %s", volume))
	}

	for _, port := range opts.Ports {
		c.addLine(fmt.Sprintf("EXPOSE %d", port))
	}

	for k, v := range opts.Env {
		c.addLine(fmt.Sprintf("ENV %s=%s", k, v))
	}

	if len(opts.Entrypoint) > 0 {
		c.addLine(fmt.Sprintf("ENTRYPOINT [\"%s\"]", strings.Join(opts.Entrypoint, "\", \"")))
	}

	if len(opts.Cmd) > 0 {
		c.addLine(fmt.Sprintf("CMD [\"%s\"]", strings.Join(opts.Cmd, "\", \"")))
	}
}
