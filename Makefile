HUB_USER = f763180872
NAME = mini-repository

.PHONY: clean

build: clean
	go mod tidy
	go build -o bin/MiniRepos src/main.go
	chmod a+x MiniRepos
run: build
	./MiniRepos
docker: clean
	docker build --no-cache -t $(NAME) .
init: clean
	check=`docker buildx ls | grep ^xBuilder`; \
	if [ "$$check" == "" ];then \
	  docker buildx create --name xBuilder --driver docker-container; \
	  docker buildx use xBuilder; \
	fi
push: init
	cp Dockerfile DockerfileX
	sed -i 's/FROM /FROM --platform=$$TARGETPLATFORM /g' DockerfileX
	docker buildx build --platform linux/arm,linux/arm64,linux/amd64 --no-cache -t $(HUB_USER)/$(NAME) -f DockerfileX . --push
	rm -rf DockerfileX
clean:
	-docker images | egrep "<none>" | awk '{print $$3}' | xargs docker rmi
	-docker ps -a | grep "\"buildkitd\"" | awk '{print $$1}' | xargs docker rm -f
	-rm -rf MiniRepos go go.sum