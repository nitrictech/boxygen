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
