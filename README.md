# gostyle-action

GitHub Action for [gostyle](https://github.com/k1LoW/gostyle)

## Usage

``` yaml
# .github/workflows/ci.yml
name: Lint and Test

on:
  push:
    branches:
      - main
  pull_request:

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      -
        uses: actions/checkout@v4
      -
        uses: actions/setup-go@v4
        with:
          go-version-file: go.mod
      -
        uses: k1LoW/gostyle-action@v1
        with:
          config-file: .gostyle.yml
      -
        name: Run tests
        run: go test ./...
```

See [action.yml](action.yml).

## Reference

- [golang/govulncheck-action](https://github.com/golang/govulncheck-action): GitHub Action for govulncheck
  - Referring to action.yml including variable names.
