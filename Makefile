bin_api:
	go build -o waitress app/api/main.go

run_api: bin_api
	./waitress