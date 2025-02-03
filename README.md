
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/Trendyol/overlog/badge)](https://scorecard.dev/viewer/?uri=github.com/Trendyol/overlog)
# Overlog
_Golang Logging with **Mapped Diagnostic Context**_

## What is it? 
This package is [Zerolog](https://github.com/rs/zerolog) wrapper with Mapped Diagnostic Context structure like Sl4j.

## Installation
```bash
go get -u github.com/Trendyol/overlog
```

## Getting Started
> over.New() and over.NewDefault() initialize global log field.
> You can use over.Log() from anywhere after initialize.

Quick start: 
```go
package main

import (
	over "github.com/Trendyol/overlog"
)

func main() {
	over.NewDefault()
	over.MDC().Set("x-correlation-id", "1234")
	over.MDC().Set("x-agent-name", "trendyol.com")
	over.SetGlobalFields([]string{"x-correlation-id", "x-agent-name"})
    
	over.Log().Info("hello world")
}

// Output: {"level":"INFO","time":"2020-05-20 13:02:10,000","source":"overlog/overlogger.go:70","x-agent-name":"trendyol","x-correlation-id":"1234","message":"hello world"}
```

You can set your own [zerolog](https://github.com/rs/zerolog) settings.
```go
package main

import (
	over "github.com/Trendyol/overlog"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	zlogger := zerolog.New(os.Stderr).With().Str("foo", "bar").Logger()
	over.New(zlogger)
	over.MDC().Set("x-correlation-id", "1234")
	over.AddGlobalFields("x-correlation-id")

	over.Log().Info("hello world")
}

// Output: {"level":"info","foo":"bar","x-correlation-id":"1234","message":"hello world"}
```

#### Audition middleware example for echo
```go
func main() {
	...
	over.NewDefault()
	over.SetGlobalFields([]string{"x-agent-name", "x-correlation-id"})
	...
}
```
```go
func AuditionPreHandler(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        var agentName = c.Request().Header.Get("x-agent-name")
        if agentName == "" {
        	agentName = a.instance.Config().Host
        }
        over.MDC().Set("x-agent-name", agentName)

        var correlationId = c.Request().Header.Get("x-correlation-id")
        if correlationId == "" {
            correlationId = uuid.New().String()
        }
        over.MDC().Set("x-correlation-id", correlationId)

        return next(c)
    }
}
```
```go
func (a *Application) AuditionAfterCompletion(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        c.Response().After(func() {
            over.MDC().Remove("x-agent-name")
            over.MDC().Remove("x-correlation-id")
        })
        return next(c)
    }
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.
## License
[MIT](https://github.com/Trendyol/overlog/blob/master/LICENSE)
