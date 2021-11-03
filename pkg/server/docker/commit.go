package docker_server

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Add
func (b *BuilderServer) Commit(ctx context.Context, r *v1.CommitRequest) (*v1.CommitResponse, error) {
	_, err := b.store.Get(r.Container.Id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "container: %s does not exist", r.Container.Id)
	}

	lines, err := b.store.Compile(r.Container.Id, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating container descriptor: %s", err.Error())
	}

	tmpDir := os.Getenv("BOXYGEN_TMP_DIR")
	if tmpDir == "" {
		tmpDir = "/tmp/"
	}

	// Create a temporary file
	file, err := ioutil.TempFile(tmpDir, fmt.Sprintf("%s.*.dockerfile", r.Tag))
	// cleanup the temp file when we're done
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating temporary file: %s", err.Error())
	}

	content := strings.Join(lines, "\n")

	// Write the temporary file
	file.Write([]byte(content))

	wkspc := os.Getenv("BOXYGEN_WORKSPACE")

	if wkspc == "" {
		wkspc = "/workspace/"
	}

	// Run docker build and pipe output
	cmd := exec.Command("docker", "build", wkspc, "-f", file.Name(), "-t", r.Tag)
	// Run with buildkit
	cmd.Env = append(cmd.Env, "DOCKER_BUILDKIT=1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error building image: %s", err.Error())
	}

	return &v1.CommitResponse{}, nil
}
