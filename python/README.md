# hostinfo Python

A Python port of the hostinfo test application


## Build Instructions

A [Dockerfile](./Dockerfile) is included in this repo


### Docker

Use `docker build` on the included Dockerfile with tag `docker.io/rxmllc/hostinfo/vN-python`.

For Multiarch:

Make sure you provide DockerHub credentials to your container engine with the `login` command.

Install Buildkit: `sudo apt install -y docker-buildx-plugin`
Init BuildKit: `sudo docker buildx create --use`
Build multiarch: `docker buildx build --no-cache --platform linux/arm/v7,linux/arm64/v8,linux/amd64 -t hostinfo:v1-python --push .`

Since the build is going through the Buildkit container, you will not be able to see 


### Podman

On Debian/Ubuntu using `podman`: 

`podman build --platform linux/arm/v7,linux/arm64/v8,linux/amd64 -t hostinfo:v1-python`