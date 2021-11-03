package common

import (
	"github.com/containers/storage"
	"github.com/containers/storage/pkg/unshare"
)

func GetDefaultStorage() (storage.Store, error) {
	storeOpts, err := storage.DefaultStoreOptions(unshare.IsRootless(), unshare.GetRootlessUID())

	if err != nil {
		return nil, err
	}

	return storage.GetStore(storeOpts)
}
