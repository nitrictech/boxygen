<div align="center">
	<img src="./assets/boxygen-logo.svg" alt="boxygen logo">
	<br/>
	<a href="https://hub.docker.com/r/nitrictech/boxygen-dockerfile"><img src="https://img.shields.io/docker/pulls/nitrictech/boxygen-dockerfile"/></a>
	<a href="https://codecov.io/gh/nitrictech/boxygen">
  <img src="https://codecov.io/gh/nitrictech/boxygen/branch/develop/graph/badge.svg?token=10BXB76WDM"/>
	</a>
	<img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/nitrictech/boxygen">
	<img alt="GitHub release (latest SemVer including pre-releases)" src="https://img.shields.io/github/v/release/nitrictech/boxygen?include_prereleases&label=RC%20release">
</div>

Boxygen is a container as code framework that allows you to build container images from code, allowing integration of container image builds into other tooling such as servers or CLI tooling.

## SDKs

 - [Node](https://github.com/nitrictech/boxygen-node)

> If you have any requests for language support please raise an issue. (contributions are also welcome)

## Supported Backends

| Backend    | Status            | Description                                        |
| ---------- | ----------------- | -------------------------------------------------- |
| Dockerfile | 🧪 (Experimental) | Generates dockerfiles and executes build on commit |

### Possible backends

The current Dockerfile generator backend works, there is potential for supporting additional backends though depending on demand/appetite. The current gRPC contract should be compatible with:

- Containerfile generator
- Buildah
- Buildkit

If you'd like to see an additional backend or support for other methods of execution in the current Dockerfile generator e.g. `podman` or `buildah bud` raise an issue or a PR :).
