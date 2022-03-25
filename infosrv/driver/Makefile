help:
	@echo "infosrv -- view at :9999"
	@echo ""
	@echo "build    -- build local executable"
	@echo "image    -- build container image"
	@echo "toREG    -- push image to docker hub"
	@echo "deploy   -- deploy as deployment and loadbalanced service"
	@echo "redo     -- k8s redeploy"
	@echo "delete  -- delete deployment & service"
	@echo "mkpod    -- make a k8s pod"
	@echo "delpod   -- delete k8s pod"
	@echo "drun     -- run local docker image"
	@echo "irun     -- run local docker image with shell"
	@echo "logs     -- show logs for deployment (each pod)"

-include .loccfg.mk
REGISTRY ?=larryrau

build:
	go install infosrv.go

image:
	GOOS=linux GOARCH=amd64 go build infosrv.go
	docker build -t $(REGISTRY)infosrv -f Dockerfile .
	docker tag $(REGISTRY)infosrv infosrv

#push to registry
toreg: image
	docker push $(REGISTRY)infosrv

loc:
	kubectl apply -k k8s/local

kust:
	kustomize build k8s/local | kubectl apply -f -

kustrm:
	kustomize build k8s/local | kubectl delete -f -

kustdump:
	kustomize build k8s/local

deploynovol:
	kubectl apply -f service.yaml
	cat deploy-novol.yaml |(REGISTRY=$(REGISTRY)  envsubst) >> dp
	kubectl apply -f dp
	@rm dp

deploy: xx
	kubectl apply -f pvolclaim.yaml
	kubectl apply -f service.yaml
	kubectl apply -f xx
	@rm xx

redo:
	kubectl rollout restart deployment/infosrv

delete:
	kubectl delete service/infosrv
	kubectl delete deployment/infosrv
	kubectl delete pvc/data-vol

mkpod:
	kubectl run infosrv --image=$(REGISTRY)infosrv --port=9999
	kubectl get pods

delpod:
	kubectl delete pod/infosrv

drun:
	docker run --rm --name infosrv -it -p 9999:9999 $(REGISTRY)infosrv

irun:
	docker run -it -p 9999:9999 -u root --entrypoint /bin/sh $(REGISTRY)infosrv

logs:
	kubectl logs --tail=100 -l name=infosrv

xx:
	cat deployment.yaml |(REGISTRY=$(REGISTRY)  envsubst) >> xx
