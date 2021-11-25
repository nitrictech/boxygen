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
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverWriter struct {
	srv v1.Builder_CommitServer
}

func (sw *serverWriter) Write(b []byte) (int, error) {
	logStr := string(b)

	if err := sw.srv.Send(&v1.OutputResponse{
		Log: []string{logStr},
	}); err != nil {
		return 0, err
	}

	return len(b), nil
}

// Add
func (b *BuilderServer) Commit(r *v1.CommitRequest, srv v1.Builder_CommitServer) error {
	c, err := b.store.Get(r.Container.Id)

	if err != nil {
		return status.Errorf(codes.NotFound, "container state not found for: %s", r.Container.Id)
	}

	wr := &serverWriter{srv}

	cl, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		status.Errorf(codes.NotFound, "container: %s does not exist", r.Container.Id)
	}

	lines, err := b.store.Compile(r.Container.Id, nil)
	if err != nil {
		status.Errorf(codes.Internal, "error creating container descriptor: %s", err.Error())
	}

	// Create a temporary file
	file, _ := ioutil.TempFile(b.workspace, fmt.Sprintf("%s.*.dockerfile", r.Tag))
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

	rc, err := archive.TarWithOptions(b.workspace, &archive.TarOptions{})

	if err != nil {
		return status.Errorf(codes.Internal, "error tarballing workspace: %s", err.Error())
	}

	resp, err := cl.ImageBuild(context.TODO(), rc, types.ImageBuildOptions{
		Version:    types.BuilderBuildKit,
		Dockerfile: filepath.Base(file.Name()),
		Tags:       []string{r.Tag},
	})

	if err != nil {
		return status.Errorf(codes.Internal, "error building image: %s", err.Error())
	}

	defer resp.Body.Close()

	termFs, isTerm := term.GetFdInfo(os.Stderr)
	jsonmessage.DisplayJSONMessagesStream(resp.Body, wr, termFs, isTerm, nil)

	return nil
}
