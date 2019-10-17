#----------------------------------------------------------------------------------
# Base
#----------------------------------------------------------------------------------

ROOTDIR := $(shell pwd)
OUTPUT_DIR ?= $(ROOTDIR)/_output
SOURCES := $(shell find . -name "*.go" | grep -v test.go | grep -v '\.\#*')
RELEASE := "true"
ifeq ($(TAGGED_VERSION),)
	TAGGED_VERSION := $(shell git describe --tags)
	# This doesn't work in CI, need to find another way...
	# TAGGED_VERSION := vdev
	# RELEASE := "false"
endif
VERSION ?= $(shell echo $(TAGGED_VERSION) | cut -c 2-)

LDFLAGS := "-X github.com/gloo/pkg/version.Version=$(VERSION)"
GCFLAGS := all="-N -l"

#----------------------------------------------------------------------------------
# Repo setup
#----------------------------------------------------------------------------------

# https://www.viget.com/articles/two-ways-to-share-git-hooks-with-your-team/
.PHONY: init
init:
	git config core.hooksPath .githooks

.PHONY: update-deps
update-deps:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/gogo/protobuf/gogoproto
	go get -u github.com/gogo/protobuf/protoc-gen-gogo
	mkdir -p $$GOPATH/src/github.com/envoyproxy
	# use a specific commit (c15f2c24fb27b136e722fa912accddd0c8db9dfa) until v0.0.15 is released, as in v0.0.14 the import paths were not yet changed
	cd $$GOPATH/src/github.com/envoyproxy && if [ ! -e protoc-gen-validate ];then git clone https://github.com/envoyproxy/protoc-gen-validate; fi && cd protoc-gen-validate && git fetch && git checkout c15f2c24fb27b136e722fa912accddd0c8db9dfa
	go get -u github.com/paulvollmer/2gobytes
	go get -v -u github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

.PHONY: pin-repos
pin-repos:
	go run pin_repos.go

.PHONY: check-format
check-format:
	NOT_FORMATTED=$$(gofmt -l ./projects/ ./pkg/ ./test/) && if [ -n "$$NOT_FORMATTED" ]; then echo These files are not formatted: $$NOT_FORMATTED; exit 1; fi

check-spelling:
	./ci/spell.sh check
#----------------------------------------------------------------------------------
# Clean
#----------------------------------------------------------------------------------

# Important to clean before pushing new releases. Dockerfiles and binaries may not update properly
.PHONY: clean
clean:
	rm -rf _output
	rm -rf _test
	rm -fr site
	git clean -f -X install

#----------------------------------------------------------------------------------
# Generate mocks
#----------------------------------------------------------------------------------

# The values in this array are used in a foreach loop to dynamically generate the
# commands in the generate-client-mocks target.
# For each value, the ":" character will be replaced with " " using the subst function,
# thus turning the string into a 3-element array. The n-th element of the array will
# then be selected via the word function
MOCK_RESOURCE_INFO := \
	gloo:artifact:ArtifactClient \
	gloo:endpoint:EndpointClient \
	gloo:proxy:ProxyClient \
	gloo:secret:SecretClient \
	gloo:settings:SettingsClient \
	gloo:upstream:UpstreamClient \
	gateway:gateway:GatewayClient \
	gateway:virtual_service:VirtualServiceClient\
	gateway:route_table:RouteTableClient\

# Use gomock (https://github.com/golang/mock) to generate mocks for our resource clients.
.PHONY: generate-client-mocks
generate-client-mocks:
	@$(foreach INFO, $(MOCK_RESOURCE_INFO), \
		echo Generating mock for $(word 3,$(subst :, , $(INFO)))...; \
		mockgen -destination=projects/$(word 1,$(subst :, , $(INFO)))/pkg/mocks/mock_$(word 2,$(subst :, , $(INFO)))_client.go \
     		-package=mocks \
     		github.com/solo-io/gloo/projects/$(word 1,$(subst :, , $(INFO)))/pkg/api/v1 \
     		$(word 3,$(subst :, , $(INFO))) \
     	;)

#----------------------------------------------------------------------------------
# Dev env
#----------------------------------------------------------------------------------
dev:
	docker build -t gloo-dev-env -f Dockerfile.dev .
	docker run -it --rm \
        -v $$PWD:/go/src/github.com/gloo \
        -w /go/src/github.com/gloo \
        -e GITHUB_TOKEN=$$GITHUB_TOKEN \
        -e KUBECONFIG=$$KUBECONFIG \
        gloo-dev-env

#----------------------------------------------------------------------------------
# glooctl
#----------------------------------------------------------------------------------
CLI_DIR=projects/gloo/cli
glooctl:
	CGO_ENABLED=0 GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $(OUTPUT_DIR)/$@ $(CLI_DIR)/cmd/main.go

glooctl-docker:
	docker build -t pinchaslev/glooctl-img:$(VERSION) -f $(CLI_DIR)/Dockerfile.dev .

glooctl-push: glooctl-docker
	docker login 
	docker push pinchaslev/glooctl-img:$(VERSION)

#----------------------------------------------------------------------------------
# Gateway
#----------------------------------------------------------------------------------
GATEWAY_DIR=projects/gateway
GATEWAY_SOURCES=$(call get_sources,$(GATEWAY_DIR))
gateway: $(GATEWAY_SOURCES)
	CGO_ENABLED=0 GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $(OUTPUT_DIR)/$@ $(GATEWAY_DIR)/cmd/main.go

gateway-docker:
	docker build -t pinchaslev/gateway-img:$(VERSION) -f $(GATEWAY_DIR)/Dockerfile.gateway.dev .

gateway-push: gateway-docker
	docker login 
	docker push pinchaslev/gateway-img:$(VERSION)

.PHONY: gateway

#----------------------------------------------------------------------------------
# Discovery
#----------------------------------------------------------------------------------
DISCOVERY_DIR=projects/discovery
DISCOVERY_SOURCES=$(call get_sources,$(DISCOVERY_DIR))
discovery:
	CGO_ENABLED=0 GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(DISCOVERY_DIR)/cmd/main.go

discovery-docker:
	docker build -t pinchaslev/discovery-img:$(VERSION) -f $(DISCOVERY_DIR)/Dockerfile.discovery.dev .

discovery-push: discovery-docker
	docker login 
	docker push pinchaslev/discovery-img:$(VERSION)
   
.PHONY: discovery

#----------------------------------------------------------------------------------
# Gloo
#----------------------------------------------------------------------------------
GLOO_DIR=projects/gloo
GLOO_SOURCES=$(call get_sources,$(GLOO_DIR))
gloo:
	CGO_ENABLED=0 GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(GLOO_DIR)/cmd/main.go

gloo-docker:
	docker build -t pinchaslev/gloo-img:$(VERSION) -f $(GLOO_DIR)/Dockerfile.gloo.dev .

gloo-push: gloo-docker
	docker login 
	docker push pinchaslev/gloo-img:$(VERSION)
   
.PHONY: gloo

#----------------------------------------------------------------------------------
# Envoy init
#----------------------------------------------------------------------------------
ENVOYINIT_DIR=projects/envoyinit
ENVOYINIT_SOURCES=$(call get_sources,$(ENVOYINIT_DIR))
envoyinit:
	CGO_ENABLED=0 GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(ENVOYINIT_DIR)/cmd/main.go

envoyinit-docker:
	docker build -t pinchaslev/envoy-wrapper-img:$(VERSION) -f $(ENVOYINIT_DIR)/Dockerfile.envoy.dev .

envoyinit-push: envoyinit-docker
	docker login 
	docker push pinchaslev/envoy-wrapper-img:$(VERSION)
   
.PHONY: envoyinit


#----------------------------------------------------------------------------------
# Build All
#----------------------------------------------------------------------------------
.PHONY: build
build: gloo glooctl gateway discovery envoyinit

#----------------------------------------------------------------------------------
# Deployment Manifests / Helm
#----------------------------------------------------------------------------------

HELM_SYNC_DIR := $(OUTPUT_DIR)/helm
HELM_DIR := install/helm
INSTALL_NAMESPACE ?= gloo-system

.PHONY: manifest
manifest: prepare-helm install/gloo-gateway.yaml update-helm-chart

# creates Chart.yaml, values.yaml, values-knative.yaml, values-ingress.yaml. See install/helm/gloo/README.md for more info.
prepare-helm:
	go run install/helm/gloo/generate.go $(VERSION)

update-helm-chart:
	mkdir -p $(HELM_SYNC_DIR)/charts
	helm package --destination $(HELM_SYNC_DIR)/charts $(HELM_DIR)/gloo
	helm repo index $(HELM_SYNC_DIR)

HELMFLAGS ?= --namespace $(INSTALL_NAMESPACE) --set namespace.create=true

MANIFEST_OUTPUT = > /dev/null
ifneq ($(BUILD_ID),)
MANIFEST_OUTPUT =
endif

install/gloo-gateway.yaml: prepare-helm
	helm template install/helm/gloo $(HELMFLAGS) | tee $@ $(OUTPUT_YAML) $(MANIFEST_OUTPUT)


.PHONY: render-yaml
render-yaml: install/gloo-gateway.yaml

.PHONY: save-helm
save-helm:
ifeq ($(RELEASE),"true")
	gsutil -m rsync -r './_output/helm' gs://solo-public-helm/
endif

.PHONY: fetch-helm
fetch-helm:
	gsutil -m rsync -r gs://solo-public-helm/ './_output/helm'

#----------------------------------------------------------------------------------
# Release
#----------------------------------------------------------------------------------


#----------------------------------------------------------------------------------
# Docker
#----------------------------------------------------------------------------------
#
#---------
#--------- Push
#---------

DOCKER_IMAGES :=
ifeq ($(RELEASE),"true")
	DOCKER_IMAGES := docker
endif

.PHONY: docker docker-push
docker: discovery-docker gateway-docker gloo-docker envoyinit-docker

# Depends on DOCKER_IMAGES, which is set to docker if RELEASE is "true", otherwise empty (making this a no-op).
# This prevents executing the dependent targets if RELEASE is not true, while still enabling `make docker`
# to be used for local testing.
# docker-push is intended to be run by CI
docker-push: $(DOCKER_IMAGES)
ifeq ($(RELEASE),"true")
	docker push pinchaslev/gateway-img:$(VERSION) && \
	docker push pinchaslev/discovery-img:$(VERSION) && \
	docker push pinchaslev/gloo-img:$(VERSION) && \
	docker push pinchaslev/envoy-wrapper-img:$(VERSION)
endif