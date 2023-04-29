# k8s playground

## go

Just a minimal http server printing out the hostame to see where k8s routes the request

## docker

Image from scratch with statically linked go binary

## k8s

deployment and service

Next up:

- ingress
- more services
- use [httproutes](https://gateway-api.sigs.k8s.io/v1alpha2/api-types/httproute/) ?
-- looks like this can be managed by an ingress?

Ingress can do prefix routing to different backend services
