# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a comprehensive Go language learning repository that combines documentation (built with VitePress) and runnable code examples. The project is designed as a complete learning path from Go basics to microservices architecture.

## Repository Structure

The repository uses **Go workspaces** (`go.work`) to manage multiple modules:

- `docs/` - VitePress documentation site with markdown files
- `examples/` - Organized Go code examples matching the documentation structure
  - `basics/` - Fundamental Go syntax (variables, types, control flow, functions, etc.)
  - `standard-library/` - Standard library usage (fmt, time, http, context, etc.)
  - `web-development/` - Web frameworks (HTTP server, Gin framework, routing, middleware)
  - `database/` - Database operations (SQL, GORM)
  - `projects/` - Complete projects (TODO API, blog system, e-commerce, chat app)
  - `microservices/` - Microservice patterns (gRPC, service discovery, load balancing, API gateway, etc.)

Each top-level category in `examples/` is its own Go module referenced in `go.work`.

## Common Commands

### Documentation Development

```bash
# Start documentation development server (http://127.0.0.1:3000)
pnpm dev
# or
make docs-serve

# Build documentation for production
pnpm build
# or
make docs-build

# Preview built documentation
pnpm preview
```

### Go Development

```bash
# Format all Go code
make fmt
# Equivalent to: go work sync && go fmt ./...

# Run linting (requires golangci-lint)
make lint

# Run all tests with race detection
make test
# Equivalent to: go work sync && go test -v -race ./...

# Build all examples
make build

# Install development tools (golangci-lint, goimports)
make install-tools

# Check and tidy all module dependencies
make check-modules

# Complete initialization (tools + deps + format)
make init

# Quick check (format + lint)
make check

# Full validation (format + lint + test)
make full-check
```

### Running Individual Examples

Each example can be run directly:

```bash
# Run a specific example
cd examples/basics/01-variables-constants
go run main.go

# Run tests for a specific example
cd examples/basics/16-testing
go test -v

# Run benchmark tests
go test -bench=. -benchmem
```

### Running Tests

```bash
# Run all tests in workspace
go test -v -race ./...

# Run tests in a specific module
cd examples/basics
go test -v ./...

# Run a specific test function
go test -v -run TestAdd

# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Architecture Patterns

### Go Workspace Organization

The project uses Go 1.18+ workspace feature. All example modules are defined in `go.work`:
- Each major category (basics, web-development, microservices, etc.) is a separate module
- Run `go work sync` to synchronize workspace dependencies
- This allows independent versioning while keeping examples organized

### Documentation ↔ Code Mapping

Documentation structure in `docs/` directly mirrors the code examples in `examples/`:
- Each doc topic has a corresponding runnable example
- Examples demonstrate concepts explained in documentation
- When updating documentation, verify corresponding example code exists

### Project Examples Structure

The `examples/projects/` directory contains full-stack applications:
- Each project is self-contained with its own dependencies
- Projects use realistic patterns: Gin framework, SQLite/database, CORS, validation
- Projects serve as reference implementations for common application types

### Microservices Architecture

The `examples/microservices/` section demonstrates a complete microservices ecosystem:
- **gRPC Communication**: Client/server with interceptors, streaming, authentication
- **Service Discovery**: Consul and etcd implementations with health checks
- **Load Balancing**: Multiple algorithms (round-robin, weighted, least-connections, consistent hash)
- **API Gateway**: Request routing, middleware chain, circuit breaker pattern
- **Distributed Tracing**: OpenTelemetry integration for HTTP and gRPC
- **Configuration**: Apollo and Nacos config center integration
- **Message Queues**: RabbitMQ and Kafka producer/consumer patterns
- **E-commerce Example** (`06-ecommerce-microservices/`): Complete system with user, product, order services and gateway

## VSCode Configuration

The repository includes VSCode settings (`.vscode/settings.json`):
- Format on save enabled with `goimports`
- Auto-organize imports
- Test flags: `-v -race`
- Coverage overlay on save
- Lint tool: `golangci-lint`

## CI/CD Pipeline

GitHub Actions workflows (`.github/workflows/ci.yml`):
1. **Lint**: Runs `golangci-lint` on all code
2. **Test**: Executes tests with race detection and uploads coverage to codecov
3. **Build**: Compiles all examples and uploads binaries
4. **Docs**: Builds VitePress documentation

All workflows run on push to `main`/`develop` branches and pull requests.

## Development Workflow

### Adding New Examples

1. Create new directory under appropriate `examples/` category
2. Add corresponding documentation in `docs/`
3. If adding a new category, update `go.work` and `docs/.vitepress/config.ts` sidebar
4. Run `make check` before committing

### Modifying Microservices Examples

Many microservices examples require external dependencies (Consul, etcd, Kafka, etc.). The code demonstrates patterns but may need:
- Docker Compose setup for local testing
- Configuration adjustments for endpoints/ports
- Proto file compilation for gRPC examples

### Documentation Updates

Documentation is built with VitePress 1.6.4:
- Markdown files in `docs/` directory
- Sidebar configuration in `docs/.vitepress/sidebar.ts`
- Supports task lists via `markdown-it-task-lists` plugin
- Code syntax highlighting with line numbers enabled
- Local search configured

## Testing Conventions

From `examples/basics/16-testing/main_test.go`:
- Unit tests: `TestFunctionName(t *testing.T)`
- Table-driven tests: Use `t.Run()` with struct slices for multiple test cases
- Benchmark tests: `BenchmarkFunctionName(b *testing.B)`
- Example tests: `ExampleFunctionName()` with `// Output:` comment

## Important Notes

- **Go Version**: Minimum Go 1.21 (workspace uses 1.25.4)
- **Node.js**: Minimum v18 for VitePress documentation
- **Package Manager**: Uses `pnpm` (v10.23.0+) for documentation dependencies
- **Line Endings**: Repository enforces LF (`\n`) line endings
- **Import Organization**: Auto-organized on save via `goimports`

## Known Issues

From README.md common problems section:
- Windows `go mod tidy` permission errors: Create new module first with `go mod init test`
- Slow downloads: Configure Go proxy with `go env -w GOPROXY=https://goproxy.cn,direct`
