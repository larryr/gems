
name = msgrelay
PORT = 9901

help:
	@echo $(name)" -- view at :"$(PORT)
	@echo ""
	@echo "build    -- build local executable"
	@echo "image    -- build container image"
	@echo "toREG    -- push image to docker hub"
	@echo "drun     -- run local docker image"
	@echo "irun     -- run local docker image with shell"

-include .loccfg.mk
REGISTRY ?=larryrau/

build:
	go install $(name).go

image:
	GOOS=linux GOARCH=amd64 go build $(name).go
	docker build -t $(REGISTRY)$(name) -f Dockerfile .
	docker tag $(REGISTRY)$(name) $(name)

drun:
	docker run --rm --name $(name) -it -p $(PORT):$(PORT) $(REGISTRY)$(name)

irun:
	docker run -it -p $(PORT):$(PORT) -u root --entrypoint /bin/sh $(REGISTRY)$(name)

#push to registry
toreg: image
	docker push $(REGISTRY)$(name)

