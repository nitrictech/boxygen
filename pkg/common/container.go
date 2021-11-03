package common

import (
	"github.com/containers/buildah"
	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
)

// Get a builder for the given container
// This assumes that the builder already exists
func BuilderForContainer(c *v1.Container) (*buildah.Builder, error) {
	store, err := GetDefaultStorage()

	if err != nil {
		return nil, err
	}

	return buildah.OpenBuilder(store, c.GetId())
}
