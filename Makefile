NAME = diltram/go-chat
INSTANCE = go-chat

.PHONY: default build 

default: build

build:
	docker rm -f $(INSTANCE); true
	docker build -f build/docker/Dockerfile-build -t $(NAME)-build .
	docker create --name $(INSTANCE) $(NAME)-build
	docker cp $(INSTANCE):$$GOROOT/src/github.com/$(NAME)/cmd/go-chat/go-chat $(shell pwd)/go-chat
	docker rm $(INSTANCE)
