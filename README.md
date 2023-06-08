<p align="center" >
    <img src="assets/logo.png" alt="logo" width="250"/>
<h3 align="center">devstdout</h3>
<p align="center">Simple slog wrapper pkg</p>
<p align="center">Build with ❤ in Golang</p>
</p>

<p align="center" >
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/containerscrew/devstdout">
    <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/containerscrew/devstdout">
    <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/containerscrew/devstdout">
</p>


# Badges

![Build](https://github.com/containerscrew/devstdout/actions/workflows/build.yml/badge.svg)
![GitLeaks](https://github.com/containerscrew/devstdout/actions/workflows/gitleaks.yml/badge.svg)
[![License](https://img.shields.io/github/license/containerscrew/devstdout)](/LICENSE)

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [About](#about)
- [Example](#example)
- [Credits](#credits)
- [TO DO](#to-do)
- [Contribution](#contribution)
- [LICENSE](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# About

Simple slog wrapper pkg

# Example

```go
package main

import (
    logger "github.com/containerscrew/devstdout/pkg"
)

func main() {
    log := logger.NewLogger(
        logger.OptionsLogger{Level: "warning", AddSource: false, LoggerType: "pretty"},
    )

    log.Debug(
        "testing message",
        logger.PrintMessage("test", "Debug test"),
    )

    log.Info(
        "testing message",
        logger.PrintMessage("test", "test"),
    )

    log.Warning("warning message!")

    log.Success(
        "Success Message",
        logger.PrintMessage("test", "test"),
    )
}
```

![example](./assets/example.png)

# Credits
- [Slog](https://pkg.go.dev/golang.org/x/exp/slog)
- [Git leaks](https://github.com/gitleaks/gitleaks-action)
- [Color library](github.com/fatih/color)
- [To my teacher of Golang @gilmiriam](https://github.com/gilmiriam)

# TO DO

* Add tests
* Code refactor is certainly needed!

# Contribution

Pull requests are welcome! Any code refactoring, improvement, implementation.

# LICENSE

[LICENSE](./LICENSE)
