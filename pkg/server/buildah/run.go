package buildah_server

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/containers/buildah"
	"github.com/containers/buildah/define"
	"github.com/nitrictech/boxygen/pkg/common"
	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Run
func (b *BuilderServer) Run(ctx context.Context, r *v1.RunRequest) (*v1.RunResponse, error) {
	// Rehydrate from an existing builder
	builder, err := common.BuilderForContainer(r.Container)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	_ = builder.Run(r.Command, buildah.RunOptions{
		Stdout:     stdout,
		Stderr:     stderr,
		WorkingDir: "/workspace/",
		Isolation:  define.IsolationChroot,
		Terminal:   buildah.WithoutTerminal,
	})

	// Workaround for bad PID parsing (need to resolve this to ensure proper execution)
	errors, _ := ioutil.ReadAll(stderr)
	errorStr := string(errors)
	if len(errorStr) > 0 && !strings.Contains(errorStr, "error parsing PID") {
		fmt.Println("Run error: ", errors)
		return nil, status.Error(codes.Internal, errorStr)
	}
	// Log errors...
	// fmt.Println("Errors Encountered: ", string(errors))

	output, err := ioutil.ReadAll(stdout)

	if err != nil {
		return nil, status.Error(codes.Internal, "Error reading stdout")
	}

	fmt.Println("Ran command ", string(output))

	builder.Save()

	return &v1.RunResponse{}, nil
}
