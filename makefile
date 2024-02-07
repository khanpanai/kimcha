run-test:
	go test ./... -v -cover -coverprofile="c.out"
	go tool cover -html="c.out"
	ii .

