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
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverWriter struct {
	srv v1.Builder_CommitServer
}

func (sw *serverWriter) Write(b []byte) (int, error) {
	logStr := string(b)
	logout := strings.Split(logStr, "\n")

	if err := sw.srv.Send(&v1.OutputResponse{
		Log: logout,
	}); err != nil {
		return 0, err
	}

	return len(b), nil
}

// Add
func (b *BuilderServer) Commit(r *v1.CommitRequest, srv v1.Builder_CommitServer) error {
	c, err := b.store.Get(r.Container.Id)
	wr := &serverWriter{srv}

	if err != nil {
		status.Errorf(codes.NotFound, "container: %s does not exist", r.Container.Id)
	}

	lines, err := b.store.Compile(r.Container.Id, nil)
	if err != nil {
		status.Errorf(codes.Internal, "error creating container descriptor: %s", err.Error())
	}

	tmpDir := os.Getenv("BOXYGEN_TMP_DIR")
	if tmpDir == "" {
		tmpDir = "/tmp/"
	}

	// Create a temporary file
	file, err := ioutil.TempFile(tmpDir, fmt.Sprintf("%s.*.dockerfile", r.Tag))
	ignoreFile, err := os.Create(fmt.Sprintf("%s.dockerignore", file.Name()))

	// cleanup the temp file when we're done
	defer func() {
		file.Close()
		ignoreFile.Close()
		os.Remove(file.Name())
		os.Remove(ignoreFile.Name())
	}()

	if err != nil {
		return status.Errorf(codes.Internal, "error creating temporary file: %s", err.Error())
	}
	ignoreContent := strings.Join(c.Ignore(), "\n")
	content := strings.Join(lines, "\n")

	// Write the temporary file
	file.Write([]byte(content))
	ignoreFile.Write([]byte(ignoreContent))

	// Run docker build and pipe output
	// TODO: Add Podman support
	cmd := exec.Command("docker", "build", b.workspace, "-f", file.Name(), "-t", r.Tag)
	// Run with buildkit
	cmd.Env = append(cmd.Env, "DOCKER_BUILDKIT=1")

	cmd.Stdout = wr
	cmd.Stderr = wr

	err = cmd.Run()
	if err != nil {
		return status.Errorf(codes.Internal, "error building image: %s", err.Error())
	}

	return nil
}
