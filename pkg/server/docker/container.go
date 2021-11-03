package docker_server

type ContainerState interface {
	Name() string
	Lines() []string
	// TODO: We will want to replace this with a more op/args model to translate better between more types of container file formats
	AddLine(line string)
	AddDependency(name string)
	Dependencies() []string
}

// ContainerState
type containerStateImpl struct {
	// unique name for this container state
	name string
	// container states that this container depends on
	dependsOn []string
	// lines composing this container image state
	lines []string
}

func (c *containerStateImpl) Name() string {
	return c.name
}

func (c *containerStateImpl) Lines() []string {
	return c.lines
}

func (c *containerStateImpl) AddLine(line string) {
	if c.lines == nil {
		c.lines = make([]string, 0)
	}

	c.lines = append(c.lines, line)
}

func (c *containerStateImpl) AddDependency(name string) {
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
