# allabolag-cli
A simple CLI tool written in Go to get information from the site allabolag.se, which serves information about Swedish companies.

## Features
- Search for a company
- Show general information and fiscal results from the last five calendar years.
- Get link to all data at website.

## Installation
The simplest way, if you are a Go developer, is to install it directly to $GOBIN:
```bash
$ go install github.com/vertan/allabolag-cli/cmd/allabolag
```

Or, you can download one of the binary distributions from the releases page on Github (make sure to pick the right binary for your system):
```bash
$ wget https://github.com/vertan/allabolag-cli/releases/download/v0.1.0/allabolag-v0.1.0-darwin-amd64.tar.gz
$ tar -xvfz allabolag-darwin-amd64.tar.gz /usr/bin/local
```

## Usage
The simplest command searches for a company name and returns the top result:
```bash
$ allabolag google
```
