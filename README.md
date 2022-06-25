# envconfig

[![Go](https://github.com/nanernunes/envconfig/actions/workflows/go.yml/badge.svg)](https://github.com/nanernunes/envconfig/actions/workflows/go.yml)
[![version](https://img.shields.io/github/tag/nanernunes/envconfig.svg)](https://github.com/nanernunes/envconfig/releases/latest)
[![GoDoc](https://godoc.org/github.com/envconfig?status.png)](https://godoc.org/github.com/nanernunes/envconfig)
[![license](https://img.shields.io/github/license/nanernunes/envconfig.svg)](../LICENSE.md)
[![LoC](https://tokei.rs/b1/github/nanernunes/envconfig?category=lines)](https://github.com/nanernunes/envconfig)
[![codecov](https://codecov.io/gh/nanernunes/envconfig/branch/master/graph/badge.svg)](https://codecov.io/gh/nanernunes/envconfig)

This library is conceptually similar to [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig), with the following major behavioral differences:

- Look for `AUTO_SPLIT_VARS` by default, use `underscore:"false"` to disable it
- Adds a new annotation `underscore:"false"` to process envs like `TheEnd`
- Adds a new annotation `env:"NAME"` to fetch envs outside the scoped `PREFIX_`

```go
import "github.com/nanernunes/envconfig"
```

## Documentation

See [godoc](http://godoc.org/github.com/nanernunes/envconfig)

## Usage

Set some environment variables:

```bash
export MYAPP_HOST="localhost"
export MYAPP_PORT=9999
export MYAPP_LIVE=true
export MYAPP_DEAD="0"
```

Write some code:

```go
package main

import (
    "log"

    "github.com/nanernunes/envconfig"
)

type Entity struct {
    Host string
    Port int
    Live bool
    Dead bool
}

func main() {
    var e Entity
    err := envconfig.Process("MYAPP", &e)
    if err != nil {
        log.Fatal(err.Error())
    }

}
```

## Struct Tag Support

Envconfig supports the use of struct tags to specify alternate, default, and required
environment variables.

For example, consider the following struct:

```bash
export MYAPP_TEXT=
export MYAPP_THEZ="Z"
export JUST=alone
```

```go
type Entity struct {
    Text string `default:"example"`
    TheZ string `underscore:"false"`
    Just string `env:"JUST"`
}
```

Envconfig has automatic support for UpperCased with underscores struct elements with no required additional tags. Note that numbers
will get globbed into the previous word. If the setting does not do the
right thing, you may use a manual override.

### underscore

Envconfig will process value for `TheZ` by populating it with the
value for `MYAPP_THE_Z`. With the `underscore:"false"` tag
it would have looked up `MYAPP_THEZ`.

### default

If envconfig can't find an environment variable value for `MYAPP_TEXT`,
it will populate it with "example" as a default value.

### env

If envconfig can't find an environment variable in the form `PREFIX_JUST`, and there is a struct tag defined, it will try to populate your variable with an environment variable that directly matches the envconfig tag in your struct definition by setting an `env:"JUST"` override:

## Supported Struct Field Types

envconfig supports these struct field types:

- string
- int8, int16, int32, int64
- bool
- float32, float64
