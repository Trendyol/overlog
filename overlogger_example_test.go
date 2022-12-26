package over

import (
	"errors"
	"github.com/rs/zerolog"
	"os"
)

func ExampleOverlog_New() {
	New(zerolog.New(os.Stdout))
	MDC().Set("x-correlation-id", "1234")
	MDC().Set("x-agent-name", "trendyol")
	SetGlobalFields([]string{"x-correlation-id", "x-agent-name"})

	Log().Info("hello world")

	// Output: {"level":"info","x-agent-name":"trendyol","x-correlation-id":"1234","message":"hello world"}
}

func ExampleOverlog_Trace() {
	New(zerolog.New(os.Stdout))
	MDC().Set("x-correlation-id", "1234")
	AddGlobalFields("x-correlation-id")

	Log().Trace("hello world")

	// Output: {"level":"trace","x-correlation-id":"1234","message":"hello world"}
}

func ExampleOverlog_Tracef() {
	New(zerolog.New(os.Stdout))
	MDC().Set("x-correlation-id", "1234")
	AddGlobalFields("x-correlation-id")

	Log().Tracef("hello %s", "world")

	// Output: {"level":"trace","x-correlation-id":"1234","message":"hello world"}
}

func ExampleOverlog_Debug() {
	New(zerolog.New(os.Stdout))
	MDC().Set("x-correlation-id", "1234")
	AddGlobalFields("x-correlation-id")

	Log().Debug("hello world")

	// Output: {"level":"debug","x-correlation-id":"1234","message":"hello world"}
}

func ExampleOverlog_Debugf() {
	New(zerolog.New(os.Stdout))
	MDC().Set("x-correlation-id", "1234")
	AddGlobalFields("x-correlation-id")

	Log().Debugf("hello %s", "world")

	// Output: {"level":"debug","x-correlation-id":"1234","message":"hello world"}
}

func ExampleOverlog_Info() {
	New(zerolog.New(os.Stdout))
	MDC().Set("x-correlation-id", "1234")
	AddGlobalFields("x-correlation-id")

	Log().Info("hello world")

	// Output: {"level":"info","x-correlation-id":"1234","message":"hello world"}
}

func ExampleOverlog_Infof() {
	New(zerolog.New(os.Stdout))
	MDC().Set("x-correlation-id", "1234")
	AddGlobalFields("x-correlation-id")

	Log().Infof("hello %s", "world")

	// Output: {"level":"info","x-correlation-id":"1234","message":"hello world"}
}

func ExampleOverlog_Warn() {
	New(zerolog.New(os.Stdout))
	MDC().Set("x-correlation-id", "1234")
	AddGlobalFields("x-correlation-id")

	Log().Warn("hello world")

	// Output: {"level":"warn","x-correlation-id":"1234","message":"hello world"}
}

func ExampleOverlog_Warnf() {
	New(zerolog.New(os.Stdout))
	MDC().Set("x-correlation-id", "1234")
	AddGlobalFields("x-correlation-id")

	Log().Warnf("hello %s", "world")

	// Output: {"level":"warn","x-correlation-id":"1234","message":"hello world"}
}

func ExampleOverlog_Error() {
	New(zerolog.New(os.Stdout))
	MDC().Set("x-correlation-id", "1234")
	AddGlobalFields("x-correlation-id")

	Log().Error(errors.New("hello world"))

	// Output: {"level":"error","x-correlation-id":"1234","message":"hello world"}
}

func ExampleOverlog_Errorf() {
	New(zerolog.New(os.Stdout))
	MDC().Set("x-correlation-id", "1234")
	AddGlobalFields("x-correlation-id")

	Log().Errorf("hello %s", errors.New("world"))

	// Output: {"level":"error","x-correlation-id":"1234","message":"hello world"}
}

func ExampleOverlog_Log() {
	New(zerolog.New(os.Stdout))
	MDC().Set("x-correlation-id", "1234")
	AddGlobalFields("x-correlation-id")

	Log().Log("hello world")

	// Output: {"x-correlation-id":"1234","message":"hello world"}
}

func ExampleOverlog_Logf() {
	New(zerolog.New(os.Stdout))
	MDC().Set("x-correlation-id", "1234")
	AddGlobalFields("x-correlation-id")

	Log().Logf("hello %s", "world")

	// Output: {"x-correlation-id":"1234","message":"hello world"}
}

func ExampleOverlog_Print() {
	New(zerolog.New(os.Stdout))
	MDC().Set("x-correlation-id", "1234")
	AddGlobalFields("x-correlation-id")

	Log().Print("hello world")

	// Output: {"level":"debug","x-correlation-id":"1234","message":"hello world"}
}

func ExampleOverlog_Printf() {
	New(zerolog.New(os.Stdout))
	MDC().Set("x-correlation-id", "1234")
	AddGlobalFields("x-correlation-id")

	Log().Printf("hello %s", "world")

	// Output: {"level":"debug","x-correlation-id":"1234","message":"hello world"}
}
