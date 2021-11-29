package dockerfile

import "fmt"

type AddOptions struct {
	Src  string
	Dest string
}

func (c *containerStateImpl) Add(opts AddOptions) {
	c.addLine(fmt.Sprintf("ADD %s %s", opts.Src, opts.Dest))
}
