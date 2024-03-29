help:
	@echo "infosrv -- view at :9999"
	@echo ""
	@echo "build    -- build local executable"
	@echo "image    -- build container image"
	@echo "toREG    -- push image to docker hub"
	@echo "drun     -- run local docker image"
	@echo "irun     -- run local docker image with shell"
	@echo "** k8s helpers **"
	@echo "kdeploy   -- deploy as deployment and loadbalanced service"
	@echo "krestart     -- k8s restart"
	@echo "kdelete  -- delete deployment & service"
	@echo "kmkpod    -- make a k8s pod"
	@echo "kdelpod   -- delete k8s pod"
	@echo "klogs     -- show logs for deployment (each pod)"

-include .loccfg.mk
REGISTRY ?=larryrau/

build:
	go install infosrv.go

image:
	GOOS=linux GOARCH=amd64 go build infosrv.go
	docker build -t $(REGISTRY)infosrv -f Dockerfile .
	docker tag $(REGISTRY)infosrv infosrv

drun:
	docker run --rm --name infosrv -it -p 9999:9999 $(REGISTRY)infosrv

irun:
	docker run -it -p 9999:9999 -u root --entrypoint /bin/sh $(REGISTRY)infosrv

#push to registry
toreg: image
	docker push $(REGISTRY)infosrv

#
# k8s helpers
#

kloc:
	kubectl apply -k k8s/local

kust:
	kustomize build k8s/local | kubectl apply -f -

kustrm:
	kustomize build k8s/local | kubectl delete -f -

kustdump:
	kustomize build k8s/local

kdeploynovol:
	kubectl apply -f service.yaml
	cat deploy-novol.yaml |(REGISTRY=$(REGISTRY)  envsubst) >> dp
	kubectl apply -f dp
	@rm dp

kdeploy: kk
	kubectl apply -f pvolclaim.yaml
	kubectl apply -f service.yaml
	kubectl apply -f kk
	@rm kk

krestart:
	kubectl rollout restart deployment/infosrv

kdelete:
	kubectl delete service/infosrv
	kubectl delete deployment/infosrv
	kubectl delete pvc/data-vol

kmkpod:
	kubectl run infosrv --image=$(REGISTRY)infosrv --port=9999
	kubectl get pods

kdelpod:
	kubectl delete pod/infosrv

klogs:
	kubectl logs --tail=100 -l name=infosrv

kk:
	cat deployment.yaml |(REGISTRY=$(REGISTRY)  envsubst) >> kk
