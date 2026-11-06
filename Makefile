templ-gen:
	templ generate
dev:
	templ generate --watch --cmd="go run ." --proxy="http://localhost:8080"
run: templ-gen
	go run .
fmt:
	golangci-lint fmt
lint:
	golangci-lint run
fix:
	golangci-lint run --fix
clean:
	go clean
	rm -rf tmp
