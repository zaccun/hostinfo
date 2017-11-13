# hostinfo

Simple Go microservice which returns a host info string. Accessible via curl or wget. Listens on port 9898 by default, if an int is passed on the command line that port is used.

## Server Examples

$ ./hostinfo      # listening on port 9898

$ ./hostinfo 80   # listening on port 80

## Client Examples

$ curl localhost:9898
Ming-The-Merciless 192.168.131.12


$ wget -O - 10.90.11.43
Doctor-Zarko 10.90.11.43
