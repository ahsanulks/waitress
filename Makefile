bin_api:
	go build -o waitress app/api/main.go

run_api: bin_api
	./waitress

test:
	ENV=test go test -race -v -cover -coverprofile=cover.out ./...

order_test:
	go run functional_test/main.go
