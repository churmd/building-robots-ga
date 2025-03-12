generate:
	go generate -v ./...

format:
	go fmt ./...

test: clean_test_cache
	go test ./...

clean_test_cache:
	go clean -testcache

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

update_deps:
	go get -t -u ./...