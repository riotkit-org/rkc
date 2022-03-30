coverage:
	go test -v ./... -covermode=count -coverprofile=coverage.out

build:
	mkdir -p .build
	CGO_ENABLED=0 GO111MODULE=on go build -tags=nomsgpack -o ./.build/rkc
	chmod +x ./.build/rkc

lint:
	export GO111MODULE=on; \
	golangci-lint run \
		--verbose \
		--build-tags build
