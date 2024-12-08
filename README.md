# Core Library

[![CI](https://github.com/open-telemetry/backend-core-lib/workflows/ci/badge.svg)](https://github.com/open-telemetry/opentelemetry-go/actions?query=workflow%3Aci+branch%3Amain)
[![codecov.io](https://codecov.io/gh/open-telemetry/opentelemetry-go/coverage.svg?branch=main)](https://app.codecov.io/gh/open-telemetry/opentelemetry-go?branch=main)
[![PkgGoDev](https://pkg.go.dev/badge/go.opentelemetry.io/otel)](https://pkg.go.dev/go.opentelemetry.io/otel)
[![Go Report Card](https://goreportcard.com/badge/go.opentelemetry.io/otel)](https://goreportcard.com/report/go.opentelemetry.io/otel)
[![Slack](https://img.shields.io/badge/slack-@cncf/otel--go-brightgreen.svg?logo=slack)](https://cloud-native.slack.com/archives/C01NPAXACKT)

Core-Library is the core [Go](https://golang.org/) library of [solomon-ai](https://www.app-release.solomon-ai.io/).
It provides a set of packages to directly measure performance and behavior of your software and send this data to observability platforms.

## Features

- **Authentication Client**: Stable authentication client implementation
- **Database Support**: Beta support for MongoDB and PostgreSQL databases
- **Instrumentation**: Stable instrumentation capabilities for monitoring and tracing
- **Message Queue**: Stable producer and consumer implementations for message queues
- **Validation**: Built-in validation rules and message handling
- **Documentation**: Comprehensive documentation and design guidelines

## Project Status

| Signal  | Status     | Project |
| ------- | ---------- | ------- |
| AuthClient  | Stable     | N/A     |
| Database - Mongo | Beta       | N/A     |
| Database - Postgres | Beta       | N/A     |
| Instrumentation    | Stable | N/A     |
| Message Queue - Consumer    | Stable | N/A     |
| Message Queue - Producer    | Stable | N/A     |

Project versioning information and stability guarantees can be found in the
[versioning documentation](./VERSIONING.md).

### Compatibility

backend-core-lib ensures compatibility with the current supported versions of
the [Go language](https://golang.org/doc/devel/release#policy):

> Each major Go release is supported until there are two newer major releases.
> For example, Go 1.5 was supported until the Go 1.7 release, and Go 1.6 was supported until the Go 1.8 release.

For versions of Go that are no longer supported upstream, backend-core-lib will
stop ensuring compatibility with these versions in the following manner:

- A minor release of backend-core-lib will be made to add support for the new
  supported release of Go.
- The following minor release of backend-core-lib will remove compatibility
  testing for the oldest (now archived upstream) version of Go. This, and
  future, releases of backend-core-lib may include features only supported by
  the currently supported versions of Go.

### Supported Environments

| OS      | Go Version | Architecture |
| ------- | ---------- | ------------ |
| Ubuntu  | 1.20       | amd64        |
| Ubuntu  | 1.19       | amd64        |
| Ubuntu  | 1.20       | 386          |
| Ubuntu  | 1.19       | 386          |
| MacOS   | 1.20       | amd64        |
| MacOS   | 1.19       | amd64        |
| Windows | 1.20       | amd64        |
| Windows | 1.19       | amd64        |
| Windows | 1.20       | 386          |
| Windows | 1.19       | 386          |

While this project should work for other systems, no compatibility guarantees
are made for those systems currently.

## Documentation

- 📚 [Documentation](https://github.com/SolomonAIEngineering/backend-core-library)
- 💬 [GitHub Discussions](https://github.com/SolomonAIEngineering/backend-core-library/discussions)
- 🐛 [Issue Tracker](https://github.com/SolomonAIEngineering/backend-core-library/issues)

## Contributing

See the [contributing documentation](CONTRIBUTING.md) for guidelines about how to contribute to this project.

---

<div align="center">
    <sub>Built with ❤️ by Solomon AI Engineering</sub>
</div>
