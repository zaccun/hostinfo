# hostinfo

Simple Go microservice which returns a host info string (name and IP of the host). Accessible via curl, wget, etc. 
Listens on port 9898 by default, if an `int` is passed on the command line, that port is used instead. Particularly 
useful when demonstrating K8s service routing operation (deploy several replicas of this, create a service for the 
set, then curl the service sequentially to see various hosts respond). The hostinfo service logs inbound requests
and the client IP to stdout.

- `/python` contains a Python version of the service
- `/otel` contains a Go version of the service that logs in Open Telemetry format


## Server Examples

```
$ ./hostinfo      # listening on port 9898
```

```
$ ./hostinfo 80   # listening on port 80
```


## Client Examples

To curl the default port locally:

```
$ curl localhost:9898
Ming-The-Merciless 192.168.131.12
```

To wget an instance running on port 80 via a given IP address:

```
$ wget -O - 10.90.11.43
Doctor-Zarko 10.90.11.43
```


## Kubernetes Example

To curl a k8s service backed by hostinfo instances:

```
$ kubectl run planetmongo --image rxmllc/hostinfo --port 9898

$ kubectl expose pod planetmongo --port 80 --target-port 9898
```

Now from another pod:

```
$ kubectl run -it --rm client rxmllc/tools 

$ curl planetmongo.default.svc.cluster.local

planetmongo 10.90.11.42
```
