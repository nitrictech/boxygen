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

type ContainerState interface {
	Name() string
	Lines() []string

	Add(AddOptions)
	Copy(CopyOptions) error
	Config(ConfigOptions)
	Run(RunOptions)
	addDependency(name string)
	Dependencies() []string
	Ignore() []string
}

// ContainerState
type containerStateImpl struct {
	// unique name for this container state
	name string
	// container states that this container depends on
	dependsOn []string
	// lines composing this container image state
	lines []string
	// patterns to ignore when using file ops such as COPY
	ignore []string
	// a back reference parent store for this container
	store ContainerStateStore
}

func (c *containerStateImpl) Name() string {
	return c.name
}

func (c *containerStateImpl) Lines() []string {
	return c.lines
}

func (c *containerStateImpl) Ignore() []string {
	return c.ignore
}

func (c *containerStateImpl) addLine(line string) {
	if c.lines == nil {
		c.lines = make([]string, 0)
	}

	c.lines = append(c.lines, line)
}

func (c *containerStateImpl) addDependency(name string) {
	if c.dependsOn == nil {
		c.dependsOn = make([]string, 0)
	}

	c.dependsOn = append(c.dependsOn, name)
}

func (c *containerStateImpl) Dependencies() []string {
	if c.dependsOn == nil {
		c.dependsOn = make([]string, 0)
	}

	return c.dependsOn
}
