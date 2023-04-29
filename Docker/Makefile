# vim: noet:sw=2:ts=2:ai

.PHONY: build push 

TAG=docker.io/wollfieh/ng

build:
	docker build -t $(TAG) .

push: build
	docker push $(TAG)

run: build
	docker run --rm -p 3333:3333 $(TAG)

