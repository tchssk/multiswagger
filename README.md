# Multiswagger

A [goa](https://github.com/goadesign/goa) plugin package to generate multiple swagger specification files.

[![CircleCI](https://circleci.com/gh/tchssk/multiswagger.svg?style=shield&circle-token=7b399438f203142830a3a0bfd79e1d90d5cd9476)](https://circleci.com/gh/tchssk/multiswagger)
[![license](https://img.shields.io/github/license/tchssk/multiswagger.svg)]()

## Overview

Multiswagger is compatible with internal swagger generator of goagen (`goagen swagger`) but it expects [apidsl.Description()](https://godoc.org/github.com/goadesign/goa/design/apidsl#Description) in design to be JSON object like below.

```go
apidsl.Description(`{
	"key1": "value1",
	"key2": "value2"
}`)
```

Multiswagger generates `description: value1` as swagger.key1.json and swagger.key1.yaml, `description: value2` as swagger.key2.json and swagger.key2.yaml.

### Use case

This plugin can be used for i18n.

```go
apidsl.Description(`{
	"en": "This is an english description.",
	"ja": "これは日本語の説明です。"
}`)
```

## Installation

```sh
$ go get github.com/tchssk/multiswagger
```

## Usage


```sh
$ goagen gen --pkg-path github.com/tchssk/multiswagger --design path/to/your/design
```

## License

MIT License

## Author

Taichi Sasaki
