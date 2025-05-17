gen/swagger: ## Generate Swagger structures
	rm -rf ./internal/port/generated
	@mkdir -p ./internal/port/generated

	ogen --target ./internal/port/generated/api --clean ./api/openapi.yaml

	make gen/swagger/stubs

gen/swagger/stubs: ## Generate stubs for swagger
	gostub "./internal/port/generated/api/oas_server_gen.go" \
		--filename="service.go" \
		--per-method \
		--per-method-filename="{{ .Method.Name.Snake.Value }}.go" \
		--out="./internal/port/handlers/" \
		--package=handlers \
		--method-body=panic \
		--type-name="Service" \
		--skip-exists

gen/go: ## Generate go/mock structures
	go generate ./...

gen: ## Generate go/mock, gRPC and Swagger structures
	make gen/go
	make gen/swagger
