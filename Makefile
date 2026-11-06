ENV_FILE = .env

ifneq (,$(wildcard $(ENV_FILE)))
    include $(ENV_FILE)
    export
endif

templ-gen:
	templ generate
dev:
	templ generate --watch --cmd="go run ." --proxy="http://localhost:$$PORT"
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
