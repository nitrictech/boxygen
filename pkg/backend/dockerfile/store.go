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
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type ContainerStateStore interface {
	// Create a new container in the container state store
	NewContainer(opts NewContainerOpts) (ContainerState, error)
	// Put a new container state
	Put(string, ContainerState) error
	// Get an existing container state
	Get(string) (ContainerState, error)
	// Has - returns true if container exists in this state store
	Has(string) bool
	// Compile - Compiles the given container and it's dependencies into a ContainerFile/Dockerfile
	Compile(string, []string) ([]string, error)
}

type containerStateStoreImpl struct {
	store map[string]ContainerState
}

func (cs *containerStateStoreImpl) ensureStore() {
	if cs.store == nil {
		cs.store = make(map[string]ContainerState)
	}
}

type NewContainerOpts struct {
	From   string
	As     string
	Ignore []string
}

func newContainer(opts NewContainerOpts) (ContainerState, error) {
	var id = opts.As
	if id == "" {
		h := sha256.New()
		h.Write([]byte(opts.From))
		sum := h.Sum(nil)

		id = hex.EncodeToString(sum)
	}

	// Create a new container state
	con := &containerStateImpl{
		name:      id,
		dependsOn: make([]string, 0),
		lines:     make([]string, 0),
		ignore:    opts.Ignore,
	}

	return con, nil
}

func NewContainer(opts NewContainerOpts) (ContainerState, error) {
	con, err := newContainer(opts)
	if err != nil {
		return nil, err
	}
	conImp := con.(*containerStateImpl)
	if opts.From != "" && opts.As != "" {
		conImp.addLine(fmt.Sprintf("FROM %s as %s", opts.From, opts.As))
	} else {
		conImp.addLine(fmt.Sprintf("FROM %s", opts.From))
	}
	return con, nil
}

func (cs *containerStateStoreImpl) NewContainer(opts NewContainerOpts) (ContainerState, error) {
	cs.ensureStore()
	con, err := newContainer(opts)
	if err != nil {
		return nil, err
	}
	conImp := con.(*containerStateImpl)
	conImp.store = cs

	// The user has provided another container state as a dependency
	fromImage := opts.From
	if cs.Has(opts.From) {
		// The extension is another container state dependency
		con.addDependency(opts.From)
		fromImage = fmt.Sprintf("layer-%s", opts.From)
	}
	// Add line to container state
	conImp.addLine(fmt.Sprintf("FROM %s as layer-%s", fromImage, con.Name()))

	// Add to central container state store
	if err := cs.Put(con.Name(), con); err != nil {
		return nil, fmt.Errorf("failed to add container to state store")
	}

	return con, err
}

func (cs *containerStateStoreImpl) Has(name string) bool {
	cs.ensureStore()
	return cs.store[name] != nil
}

func (cs *containerStateStoreImpl) Put(name string, state ContainerState) error {
	cs.ensureStore()
	if cs.store[name] != nil {
		return fmt.Errorf("container already exists")
	}

	cs.store[name] = state

	return nil
}

func (cs *containerStateStoreImpl) Get(name string) (ContainerState, error) {
	if cs.store[name] == nil {
		return nil, fmt.Errorf("container state does not exist")
	}

	return cs.store[name], nil
}

func (cs *containerStateStoreImpl) Compile(name string, dependents []string) ([]string, error) {
	finalLines := []string{}

	if cs.store[name] == nil {
		return nil, fmt.Errorf("container state does not exist")
	}

	var newDeps = dependents
	if dependents != nil {
		// check if the requested state is in our list of dependents, if so then raise a cycle error

		for _, n := range dependents {
			if n == name {
				// TODO: Provide more dependency resolution detail
				return nil, fmt.Errorf("discovered dependency cycle in compilation, exiting")
			}
		}

		// Add as a new dependency if we pass
		newDeps = append(newDeps, name)
	} else {
		newDeps = []string{name}
	}

	con := cs.store[name]

	// Compile all dependencies first and prepend to this file
	// TODO: Need to catch cycles and throw if we find one
	for _, c := range con.Dependencies() {
		lines, err := cs.Compile(c, newDeps)

		if err != nil {
			return nil, err
		}

		finalLines = append(finalLines, lines...)
	}

	finalLines = append(finalLines, con.Lines()...)

	return finalLines, nil
}

func NewStateStore() ContainerStateStore {
	return &containerStateStoreImpl{
		store: make(map[string]ContainerState),
	}
}
