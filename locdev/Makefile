
help:
	@echo regup -- start registry up


## start a local registry
regup:
	docker run -d -p 5000:5000 --restart=always --name registry \
		-e REGISTRY_HTTP_ADDR=0.0.0.0:5000 \
		registry:2

regdown:
	docker kill registry
	docker rm registry

reginit:
	# add to /etc/hosts
