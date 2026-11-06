# Validation

Authoring date: November 6, 2026. Host: Windows. Target: Linux-focused; parsing also runs on other Go platforms.

## Local result

Go and systemd are unavailable on the authoring host. CI runs vet, race-enabled tests, compilation, and the example unit.

## Reproducible checks

The GitHub Actions workflow builds/runs the following on `ubuntu-latest`:

```sh
gofmt -w main.go main_test.go
go vet ./...
go test -race ./...
go build -o unit-check .
./unit-check examples/worker.service
```

## Manual acceptance

Run the provided fixture, introduce each documented invalid rule, and compare a representative real unit with systemd-analyze verify on its target host.

A workflow file is not evidence of a passing run. Observed CI run links and source commits are recorded after publication. Native GUI, sound, sleep/resume, and platform permissions require direct target-platform acceptance even when model tests pass.
