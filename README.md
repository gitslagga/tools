# Tools

A useful command line tools.

## Introduction

Programmer command line develop tools for efficiency.

Influenced by [it-tools](https://github.com/CorentinTh/it-tools), develop command line version of tools.

## Install

Prerequisites to build Tools from source: 

- Standard edition: Go 1.22.5 or later

Build the standard edition:

```bash
$ go install github.com/gitslagga/tools@latest
```

## Usage

Crypto tools command line help:

```bash
$ tools crypto token
$ tools crypto token -a "ulns" -l 16

$ tools crypto hash
$ tools crypto hash -s "Hello World" -e "b"
```
