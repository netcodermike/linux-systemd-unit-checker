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

## Observed CI result

Passed on November 6, 2026: [target-platform build/check](https://github.com/netcodermike/linux-systemd-unit-checker/actions/runs/42198886718).

Verified source commit: `7da647a5c7008ad08a328516c698bd3bd72348ea`. This evidence identifies the exact tested source; a later documentation-only commit does not change that source.

The commands listed above passed. Native GUI/manual acceptance is not claimed by this CI result.
