# Linux Systemd Unit Checker

Find common mistakes in systemd service files without starting, enabling, or modifying a service.

[![Build and verify](https://github.com/netcodermike/linux-systemd-unit-checker/actions/workflows/verify.yml/badge.svg)](https://github.com/netcodermike/linux-systemd-unit-checker/actions/workflows/verify.yml)

**Language:** Go · **Platform:** Linux-focused; parsing also runs on other Go platforms · **App type:** Linux service configuration tool

First-release learning implementation, with explicit scope below. This is one independently runnable app, not a placeholder in a combined repository.

## Features

- Checks section/key structure and required service commands.
- Validates service types, duplicate ExecStart, dbus BusName, and oneshot restart rules.
- Models empty assignment resets.
- Line-numbered diagnostics and distinct warning/error exit codes.

## Requirements

Go 1.22+. To perform authoritative unit validation, use systemd-analyze on your target Linux host. No external Go modules.

## Quick start

```sh
git clone https://github.com/netcodermike/linux-systemd-unit-checker.git
cd linux-systemd-unit-checker
go build -o unit-check .
./unit-check examples/worker.service
# Follow with native systemd validation on your Linux host:
systemd-analyze verify examples/worker.service
```

Commands run from the repository root. Lines marked for a particular operating system apply only there. Language tools are prerequisites, not installed by the app. Close desktop windows normally; terminal/server apps stop with Ctrl+C unless their usage notes say otherwise.

## Example workflow

Check the provided worker, then copy it and add a second ExecStart under Type=simple; expect an error. Change Type to oneshot and verify duplicate commands are permitted. Remove Description to see a warning. The checker never installs or starts these fixtures.

## Architecture, configuration, and data

`main.go` parses a supported INI-like subset into section/key entry lists. Empty values reset prior entries. Semantic checks inspect the effective values and command list. `main_test.go` covers valid units, duplicate commands, reset behavior, stop-only services, and malformed input. `examples/worker.service` is a benign one-shot fixture. Exit codes: 0 no supported-subset issues, 1 warnings, 2 errors or bad invocation. The app reads one file and never executes service commands.

## Verification

Build with the quick-start command where needed, then run:

```sh
go test ./...
```

Manual acceptance: Run the provided fixture, introduce each documented invalid rule, and compare a representative real unit with systemd-analyze verify on its target host.

See [validation record](docs/VALIDATION.md) for observed checks and their limits. CI results verify only what the workflow runs; they do not establish native GUI correctness or production readiness.

## Limitations and next steps

This is a subset linter, not a systemd parser or security audit. Drop-ins, specifier expansion, quoting nuances, executable path validation, dependency cycles, unknown keys, and line continuations are not fully interpreted. Continuations get a warning and may produce misleading additional diagnostics. Checks assume a modern systemd release; compare with your host version. Always follow with systemd-analyze verify.

## Troubleshooting

Exit 1 means a warning rather than a crash; inspect diagnostics. Files over the per-line scanner limit produce an error. Use an ordinary .service unit rather than a timer/socket unit. Cross-platform parsing does not imply a non-Linux service manager exists.

## Source map

- `go.mod`
- `main.go`
- `main_test.go`
- `examples/worker.service`

## Development and license

Read [contributing](CONTRIBUTING.md) before changing the app, and [security/privacy](SECURITY.md) before supplying private data. Use synthetic fixtures for tests. No telemetry, paid API key, or account is required by the app itself. Dependencies, when present, are stated in Requirements.

[MIT license](LICENSE), copyright 2026 netcodermike.
