# hostinfo Python

A Python port of the hostinfo test application.


## Build Instructions

A [Dockerfile](./Dockerfile) is included in this repo to create a container.


### Docker

Use `docker build` on the included Dockerfile with tag `docker.io/rxmllc/hostinfo/vN-python`.

For Multiarch:

Make sure you provide DockerHub credentials to your container engine with the `login` command.

- Install Buildkit: `sudo apt install -y docker-buildx-plugin`
- Init BuildKit: `sudo docker buildx create --use`
- Build multiarch: `docker buildx build --no-cache --platform linux/arm/v7,linux/arm64/v8,linux/amd64 -t docker.io./rxmllc/hostinfo:v1-python --push .`


### Podman

On Debian/Ubuntu using `podman`: 

`podman build --platform linux/arm/v7,linux/arm64/v8,linux/amd64 -t hostinfo:v1-python`