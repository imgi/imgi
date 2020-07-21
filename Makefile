#

docker-dev: tools/docker/dev.Dockerfile
	docker build -t imgi/imgi-dev -f tools/docker/dev.Dockerfile .

.PHONY: docker
