# Contributing to Go-ExprTk

Thank you for considering contributing to Go-ExprTk! Here are some guidelines to help you get started.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/your-username/go-exprtk.git`
3. Create a branch: `git checkout -b my-feature`
4. Make your changes
5. Run tests: `go test -v -race ./...`
6. Commit and push your changes
7. Open a Pull Request

## Development Requirements

- Go 1.21 or later
- A C++ compiler (GCC, Clang, or MSVC)
- Git

## Running Tests

```bash
# Run all tests
go test -v ./...

# Run tests with race detection
go test -v -race ./...

# Run benchmarks
go test -bench=. -benchmem ./...
```

## Code Guidelines

- Follow standard Go conventions (`gofmt`, `go vet`)
- Add tests for new features
- Update documentation for API changes
- Keep backward compatibility in mind

## Reporting Issues

- Use the GitHub issue tracker
- Include Go version, OS, and compiler info
- Provide a minimal reproducible example

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
