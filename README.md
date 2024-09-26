<p align="center" >
    <img src="img/logo.png" alt="logo" width="250"/>
<h3 align="center">devstdout</h3>
<p align="center">Simple slog wrapper pkg</p>
<p align="center">Built with ❤ in Golang</p>
</p>

<p align="center" >
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/containerscrew/devstdout">
    <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/containerscrew/devstdout">
    <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/containerscrew/devstdout">
    <img alt="LICENSE" src="https://img.shields.io/github/license/containerscrew/devstdout">
    <img alt="LICENSE" src="https://github.com/containerscrew/devstdout/actions/workflows/test.yml/badge.svg">
</p>

<!-- START OF TOC !DO NOT EDIT THIS CONTENT MANUALLY-->
**Table of Contents**  *generated with [mtoc](https://github.com/containerscrew/mtoc)*
- [devstdout](#devstdout)
- [Examples](#examples)
  - [Full code examples](#full-code-examples)
  - [Pretty output](#pretty-output)
  - [Json output](#json-output)
- [TO DO](#to-do)
- [License](#license)
<!-- END OF TOC -->

# devstdout

Simple slog wrapper pkg for my Golang projects.

# Examples

```go
package main

import (
	devstdout "github.com/containerscrew/devstdout/pkg"
)

func main() {
	log := devstdout.NewLogger(
		devstdout.OptionsLogger{Level: "debug", AddSource: false, LoggerType: "pretty"},
	)

	log.Debug(
		"testing message",
		devstdout.Argument("hello", "world"),
	)

	log.Info(
		"testing message",
		devstdout.Argument("bob", "marley"),
	)

	log.Warning("warning message!")

	log.Success(
		"Success Message",
		devstdout.Argument("alice", "bob"),
	)

	log.Error("error in your app!", devstdout.Argument("error", "your_error_is_here"))

	log.ErrorWithExit("fatal error, app must stop!", devstdout.Argument("error", "your_error_is_here"))
}
```

## Full code examples

* [**json**](./examples/json/json.go)
* [**pretty**](./examples/pretty/pretty.go)
* [**console**](./examples/console/console.go)

## Pretty output
![example](./img/example.png)

## Json output
![example2](./img/example2.png)

# TO DO

* Add tests
* Code refactor is certainly needed!
* AddSource option in logger is too much verbose

# License

[LICENSE](./LICENSE)
