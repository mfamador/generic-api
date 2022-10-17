REPO=marcoamador
NAME=generics-api
VERSION=0.1.0

ifeq (,$(shell which richgo))
$(warning "richgo not found, consider installing it")
	GO=go
else
	GO=richgo
endif

ifeq (,$(shell which govulncheck))
$(error "govulncheck not found, consider installing it")
endif

all: docker
clean: docker-clean
build: compile
test: unit integration

run:
	@go run cmd/api/main.go -c config/pipeline.yaml

compile:
	@go build -o /dev/null -ldflags "-s -w" ./{{ .CmdDir }}
	@# compiling without binary output (remove -o /dev/null if you want to generate a binary)

deps:
	@go mod tidy -v

unit:
	@$(GO) test  ./... -short -v -covermode=count -count=1 -coverprofile=.unit-coverage-$(NAME).out -timeout 300s && \
	go tool cover -html=.unit-coverage-$(NAME).out -o /tmp/unit-coverage-$(NAME).html && \
	open /tmp/unit-coverage-$(NAME).html

integration:
	@$(GO) test ./test/... -run TestIntegration -v -covermode=count -count=1 -coverprofile=.integration-coverage-$(NAME).out -coverpkg=./... && \
	go tool cover -html=.integration-coverage-$(NAME).out -o /tmp/integration-coverage-$(NAME).html && \
	open /tmp/integration-coverage-$(NAME).html

cover:
	@$(GO) test ./... -v -covermode=count -count=1 -coverprofile=.coverage.out -coverpkg=./... -timeout 300s && \
	go tool cover -html=.coverage.out -o /tmp/coverage.html && \
	open /tmp/unit-coverage-$(NAME).html

docker:
	@docker build -f build/Dockerfile -t $(REPO)/$(NAME):$(VERSION) \
              --no-cache --build-arg git_access_token=$(GIT_ACCESS_TOKEN) .

docker-push:
	@docker push $(REPO)/$(NAME):$(VERSION)

docker-clean:
	@docker rmi $(REPO)/$(NAME):$(VERSION)

docker-run:
	@docker run -f build/Dockerfile --rm --name $(NAME) $(REPO)/$(NAME):$(VERSION) \
            --build-arg git_access_token=$(GIT_ACCESS_TOKEN) .

vuln:
	govulncheck ./...

.PHONY: all clean build test run deps integration unit docker docker-push docker-clean docker-run vuln
