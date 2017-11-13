# hostinfo

Simple Go microservice which returns a host info string. Accessible via curl, wget, etc.. Listens on port 9898 by default, if an int is passed on the command line that port is used. Particularly useful when demonstrating k8s service routing mesh operation (deploy several replicas of these, create a service for them then curl the service sequentially to see various hosts engaged).


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
$ kubectl run planetmongo --image rxmllc/hostinfo --port 9898 --replicas=2

$ kubectl expose deployment planetmongo --port 80 --target-port 9898

$ curl planetmongo.default.svc.cluster.local
Ming-The-Merciless 192.168.131.12

$ curl planetmongo.default.svc.cluster.local
Doctor-Zarko 10.90.11.43
```
