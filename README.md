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
$ wget https://github.com/vertan/allabolag-cli/releases/download/v0.2.0/allabolag-v0.2.0-darwin-amd64.tar.gz
$ tar -xzvf allabolag-darwin-amd64.tar.gz /usr/bin/local
```

## Usage
The simplest command searches for a company name and returns the top result:
```bash
$ allabolag google

Google Sweden AB
https://www.allabolag.se/5566566880/
--------------------
Year  Revenue   Result
2018  1115347k  89113k
2017  802348k   64858k
2016  630652k   50998k
2015  525204k   43359k
2014  389061k   32456k
```

## TODO
* Automate binary building for releases
