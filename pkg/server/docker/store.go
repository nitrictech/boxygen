package docker_server

import "fmt"

type ContainerStateStore interface {
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
				return nil, fmt.Errorf("discovered depdency cycle in compilation, exiting!")
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
