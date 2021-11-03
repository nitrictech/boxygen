package buildah_server

import (
	"context"

	"github.com/containers/buildah"
	"github.com/containers/buildah/define"
	"github.com/containers/common/pkg/config"
	"github.com/containers/common/pkg/parse"
	"github.com/containers/image/v5/types"
	"github.com/nitrictech/boxygen/pkg/common"
	v1 "github.com/nitrictech/boxygen/pkg/proto/builder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// From - Creates a new builder (working container, from an existing image)
func (b *BuilderServer) From(ctx context.Context, r *v1.FromRequest) (*v1.FromResponse, error) {
	config, err := config.Default()

	if err != nil {
		return nil, err
	}

	store, err := common.GetDefaultStorage()

	if err != nil {
		return nil, err
	}

	cap, err := config.Capabilities("", []string{"CAP_SYS_ADMIN", "CAP_SYS_CHROOT"}, []string{})

	if err != nil {
		return nil, err
	}

	devices := define.ContainerDevices{}
	for _, device := range append(config.Containers.Devices) {
		dev, err := parse.DeviceFromPath(device)
		if err != nil {
			return nil, err
		}
		devices = append(devices, dev...)
	}

	options := buildah.BuilderOptions{
		FromImage: r.Image,
		// Default is ""
		// Container:  iopts.name,
		// TODO: Make pull policy configurable in code...
		PullPolicy: define.PullIfMissing,
		// SignaturePolicyPath:   signaturePolicy,
		SystemContext: &types.SystemContext{},
		// DefaultMountsFilePath: "",
		// TODO: Define isolation options
		Isolation:        define.IsolationChroot,
		NamespaceOptions: define.NamespaceOptions{},
		ConfigureNetwork: define.NetworkDefault,
		//CNIPluginPath:    iopts.CNIPlugInPath,
		//CNIConfigDir:     iopts.CNIConfigDir,
		//IDMappingOptions: idmappingOptions,
		Capabilities: cap,
		//CommonBuildOpts:  commonOpts,
		Format: define.OCIv1ImageManifest,
		// BlobDirectory: iopts.BlobCache,
		Devices:    devices,
		DefaultEnv: config.GetDefaultEnv(),
		// MaxPullRetries:        maxPullPushRetries,
		// PullRetryDelay:        pullPushRetryDelay,
		// OciDecryptConfig: decConfig,
	}

	builder, err := buildah.NewBuilder(context.TODO(), store, options)

	if err != nil {
		// handle error...
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	// Persist the builder for loading by other unary commands
	// TODO: May implement an ops stream instead for working containers???
	builder.Save()

	return &v1.FromResponse{
		Container: &v1.Container{
			Id: builder.ContainerID,
		},
	}, nil
}
