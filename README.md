# ORDER_MICROSERVICE


## Project Structure (API GATEWAY PATTERN)
This project includes authenction service and order service accompanied with api gateway.

# `API-GATEWAY`
The API gateway in our microservice project acts as a centralized entry point for accessing and managing the order service and auth service. It provides a unified interface for clients to interact with these services without needing to know their specific locations.
```bash
# Navigate into the project
cd ./go-gin-clean-arch

# Install dependencies
make deps

# Generate wire_gen.go for dependency injection
# Please make sure you are export the env for GOPATH
make wire

# Run the project in Development Mode
make run
```

Additional commands:

```bash
➔ make help
build                          Compile the code, build Executable File
run                            Start application
test                           Run tests
test-coverage                  Run tests and generate coverage file
deps                           Install dependencies
deps-cleancache                Clear cache in Go module
wire                           Generate wire_gen.go
swag                           Generate swagger docs
help                           Display this help screen
```

API DOCUMENTATION-SWAGGER
### `https://eventsradar.online/swagger/index.html`
