# weather-cli

A simple weather forecasting command line tool by go.

[![LICENSE](https://img.shields.io/badge/license-MIT-FF0080.svg)](https://github.com/leeifme/weather-cli/blob/master/LICENSE)
[![leeifme on github](https://img.shields.io/badge/github-@leeifme-red.svg)](https://github.com/leeifme)
[![package version](https://img.shields.io/badge/package-v0.1.0-blue.svg)](https://github.com/leeifme/weather-cli)
[![language by golang](https://img.shields.io/badge/language-@golang-green.svg)](https://github.com/leeifme/weather-clii)

## Usage

**Native**

```bash
go build -o weather-cli utils.go types.go main.go && ./weather-cli
```

**Use Lib**

```bash
go build -o weather-cli utils.go types.go cli_main.go && ./weather-cli
```

## Package

```bash
go build -ldflags "-s -w" -o weather-cli utils.go types.go cli_main.go && upx ./weather-cli
```

**Linux Arch**

```bash
GOOS=linux GOARCH=amd64 go build ...
```

**Windows Arch**

```bash
GOOS=windows GOARCH=amd64 go build ...
```

**MacOSX Arch**

```bash
GOOS=darwin GOARCH=amd64 go build ...
```

## Compression

```bash
go build -ldflags "-s -w" -o weather-cli utils.go types.go cli_main.go && upx ./weather-cli
```
