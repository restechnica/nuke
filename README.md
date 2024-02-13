# Nuke

[![github.com release badge](https://img.shields.io/github/release/restechnica/nuke.svg)](https://github.com/restechnica/nuke/)
[![github.com workflow badge](https://github.com/restechnica/nuke/workflows/main/badge.svg)](https://github.com/restechnica/nuke/actions?query=workflow%3Amain)
[![go.pkg.dev badge](https://pkg.go.dev/badge/github.com/restechnica/nuke)](https://pkg.go.dev/github.com/restechnica/nuke)
[![goreportcard.com badge](https://goreportcard.com/badge/github.com/restechnica/nuke)](https://goreportcard.com/report/github.com/restechnica/nuke)
[![img.shields.io MPL2 license badge](https://img.shields.io/github/license/restechnica/nuke)](./LICENSE)

A modern task runner experience for Nushell.

## Table of Contents

* [Features](#features)
* [Usage](#usage)
* [Requirements](#requirements)
* [How to install](#how-to-install)
  * [Github](#github)
  * [Homebrew](#homebrew)
* [How to configure](#how-to-configure)

## Features

- run commands from any subdirectory
- a set of default script names
- `.env` support

## Usage

Each command has a `-h, --help` flag available.

### `nuke`

Walks up the current directory structure in search of a `nu` [script](#default-script-names).

Once found, it will change directory to the script's directory and execute the script with any arguments passed on to `nuke`.

Initially environment variables are loaded from a `.env` file.

### `nuke version`

prints `nuke` version information.

## Requirements

`nuke` requires a [`nu`](https://www.nushell.sh/) installation.

## How to install

`nuke` can be retrieved from GitHub or a Homebrew tap. Run `nuke version` to validate the installation.
The tool is available for Windows, Linux and macOS.

### github

`nuke` is available through GitHub. The following example works for a GitHub Workflow, other CI/CD tooling will require a different path setup.

```shell
NUKE_VERSION=0.1.0
mkdir bin
echo "$(pwd)/bin" >> $GITHUB_PATH
curl -o bin/nuke -L https://github.com/restechnica/nuke/releases/download/v$NUKE_VERSION/nuke-linux-amd64
chmod +x bin/nuke
```

### homebrew

`nuke` is available through the public tap [github.com/restechnica/homebrew-tap](https://github.com/restechnica/homebrew-tap)

```shell
brew tap restechnica/tap git@github.com:restechnica/homebrew-tap.git
brew install restechnica/tap/nuke
```

### golang

`nuke` is written in golang, which means you can use `go install`. Make sure the installation folder, which depends on your golang setup, is in your system PATH.

```shell
go install github.com/restechnica/nuke/cmd/nuke@v0.1.0
```

## How to configure

`nuke` supports a set of default `nu` script names and environment variables. It currently does not support a config file.

### Environment variables

- `NUKE_LOG_LEVEL`: sets the `nuke` log level, `DEBUG` and `INFO` are supported.

### Script names

Supported default script names:

- `main.nu`
- `nuke.nu`
- `make.nu`
- `tasks.nu`
- `nukefile`
- `nukefile.nu`
- `Nukefile`
- `Nukefile.nu`