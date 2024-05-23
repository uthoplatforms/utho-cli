.PHONY: publish
publish: tidy
	git tag v$(tag)
	git push origin v$(tag)

.PHONY: test
test:
	go test -v ./...

.PHONY: test/count
test/count:
	go test -v ./... | grep -c RUN

.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v