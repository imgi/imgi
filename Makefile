#

VERSION ?= 0.0.0


docker: docker.dev docker.imgi

docker.dev: tools/docker/dev.Dockerfile
	docker build -t imgi/imgi-dev -f tools/docker/dev.Dockerfile .

docker.imgi: tools/docker/Dockerfile docker.dev
	docker build -t imgi/imgi -f tools/docker/Dockerfile .

docker.push: docker
	docker tag imgi/imgi imgi/imgi:$(VERSION)
	docker push imgi/imgi:$(VERSION)
	docker push imgi/imgi:latest

docker.run:
	docker run --rm -p 6969:6969 -v $(shell pwd)/testdata:/images imgi/imgi

.PHONY: docker.run
