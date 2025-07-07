CGO_ENABLED=0
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
all: test
all: vet
all: package
all: package_race

test: vet
test: base_test
test: staticcheck
test: shadow

base_test:
	go test ./... -v

vet:
	go vet ./...

staticcheck: staticcheck_bin
	bin/staticcheck ./...

staticcheck_bin:
	GOBIN=${ROOT_DIR}/bin go install honnef.co/go/tools/cmd/staticcheck@latest

shadow: shadow_bin
	bin/shadow ./...

shadow_bin:
	GOBIN=${ROOT_DIR}/bin go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest

package: go_cache_client

package_race: go_cache_client_race

go_cache_client:
	@go build -a -tags "netgo osusergo" -ldflags "-extldflags '-static' -s -w" -o ./bin/go_cache_client ./cmd/go_cache_client/

go_cache_client_race:
	@go build -a -tags "netgo osusergo" -ldflags "-extldflags '-static' -s -w" --race -o ./bin/go_cache_client_race ./cmd/go_cache_client/
