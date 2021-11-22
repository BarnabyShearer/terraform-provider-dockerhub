terraform-provider-dockerhub: *.go */*.go go.mod
	go build .

install: terraform-provider-dockerhub
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/BarnabyShearer/dockerhub/0.1.0/linux_amd64
	cp $+ ~/.terraform.d/plugins/registry.terraform.io/BarnabyShearer/dockerhub/0.1.0/linux_amd64
	-rm .terraform.lock.hcl
	terraform init
