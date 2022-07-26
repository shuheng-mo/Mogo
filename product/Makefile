

.PHONY: proto
proto:
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) -e ICODE=68272CAA3468CEED cap1573/cap-protoc -I ./ --micro_out=./ --go_out=./ ./proto/product/product.proto

.PHONY: build
build: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o product-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t product-service:latest
