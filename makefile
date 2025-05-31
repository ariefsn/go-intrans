test:
	go test $(shell go list ./... | grep -E '/(delivery|service)$$')

test.cover:
	go test -v -cover $(shell go list ./... | grep -E '/(delivery|service)$$')

test.cover.out:
	go test -v -coverprofile=coverage.out $(shell go list ./... | grep -E '/(delivery|service)$$')
	go tool cover -func=coverage.out

test.cover.html:
	go test -v -coverprofile=coverage.out $(shell go list ./... | grep -E '/(delivery|service)$$')
	go tool cover -html=coverage.out
