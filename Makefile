#

docker: docker.dev docker.imgi

docker.dev: tools/docker/dev.Dockerfile
	docker build -t imgi/imgi-dev -f tools/docker/dev.Dockerfile .

docker.imgi: tools/docker/Dockerfile docker.dev
	docker build -t imgi/imgi -f tools/docker/Dockerfile .

.PHONY: docker
