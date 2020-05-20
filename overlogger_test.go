package over

import (
	"bytes"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	zerolog := zerolog.New(os.Stdout).With().Timestamp().Logger()
	zerolog_with_hook := zerolog.Hook(MDCHook{})
	actual := &zerolog_with_hook

	overlog := New(zerolog)
	expected := overlog.log

	assert.Equal(t, actual, expected)
}

func TestOverlog_Trace(t *testing.T) {
	t.Run("trace log", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		Log().Trace("trace log")

		expected := "{\"level\":\"trace\",\"message\":\"trace log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
	})

	t.Run("trace log with mdc", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		SetGlobalFields([]string{"x-correlation-id"})
		MDC().Set("x-correlation-id", "1234")

		Log().Trace("trace log")

		expected := "{\"level\":\"trace\",\"x-correlation-id\":\"1234\",\"message\":\"trace log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
		ClearGlobalFields()
	})
}

func TestOverlog_Tracef(t *testing.T) {
	t.Run("tracef log", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		Log().Tracef("tracef %s",  "log")

		expected := "{\"level\":\"trace\",\"message\":\"tracef log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
	})

	t.Run("tracef log with mdc", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		SetGlobalFields([]string{"x-correlation-id"})
		MDC().Set("x-correlation-id", "1234")

		Log().Tracef("tracef %s",  "log")

		expected := "{\"level\":\"trace\",\"x-correlation-id\":\"1234\",\"message\":\"tracef log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
		ClearGlobalFields()
	})
}

func TestOverlog_Debug(t *testing.T) {
	t.Run("debug log", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		Log().Debug("debug log")

		expected := "{\"level\":\"debug\",\"message\":\"debug log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
	})

	t.Run("debug log with mdc", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		SetGlobalFields([]string{"x-correlation-id"})
		MDC().Set("x-correlation-id", "1234")

		Log().Debug("debug log")

		expected := "{\"level\":\"debug\",\"x-correlation-id\":\"1234\",\"message\":\"debug log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
		ClearGlobalFields()
	})
}

func TestOverlog_Debugf(t *testing.T) {
	t.Run("debugf log", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		Log().Debugf("debugf %s",  "log")

		expected := "{\"level\":\"debug\",\"message\":\"debugf log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
	})

	t.Run("debugf log with mdc", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		SetGlobalFields([]string{"x-correlation-id"})
		MDC().Set("x-correlation-id", "1234")

		Log().Debugf("debugf %s",  "log")

		expected := "{\"level\":\"debug\",\"x-correlation-id\":\"1234\",\"message\":\"debugf log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
		ClearGlobalFields()
	})
}

func TestOverlog_Info(t *testing.T) {
	t.Run("info log", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		Log().Info("info log")

		expected := "{\"level\":\"info\",\"message\":\"info log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
	})

	t.Run("info log with mdc", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		SetGlobalFields([]string{"x-correlation-id"})
		MDC().Set("x-correlation-id", "1234")

		Log().Info("info log")

		expected := "{\"level\":\"info\",\"x-correlation-id\":\"1234\",\"message\":\"info log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
		ClearGlobalFields()
	})
}

func TestOverlog_Infof(t *testing.T) {
	t.Run("infof log", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		Log().Infof("infof %s",  "log")

		expected := "{\"level\":\"info\",\"message\":\"infof log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
	})

	t.Run("infof log with mdc", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		SetGlobalFields([]string{"x-correlation-id"})
		MDC().Set("x-correlation-id", "1234")

		Log().Infof("infof %s",  "log")

		expected := "{\"level\":\"info\",\"x-correlation-id\":\"1234\",\"message\":\"infof log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
		ClearGlobalFields()
	})
}

func TestOverlog_Warn(t *testing.T) {
	t.Run("warn log", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		Log().Warn("warn log")

		expected := "{\"level\":\"warn\",\"message\":\"warn log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
	})

	t.Run("warn log with mdc", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		SetGlobalFields([]string{"x-correlation-id"})
		MDC().Set("x-correlation-id", "1234")

		Log().Warn("warn log")

		expected := "{\"level\":\"warn\",\"x-correlation-id\":\"1234\",\"message\":\"warn log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
		ClearGlobalFields()
	})
}

func TestOverlog_Warnf(t *testing.T) {
	t.Run("warnf log", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		Log().Warnf("warnf %s",  "log")

		expected := "{\"level\":\"warn\",\"message\":\"warnf log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
	})

	t.Run("warnf log with mdc", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		SetGlobalFields([]string{"x-correlation-id"})
		MDC().Set("x-correlation-id", "1234")

		Log().Warnf("warnf %s",  "log")

		expected := "{\"level\":\"warn\",\"x-correlation-id\":\"1234\",\"message\":\"warnf log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
		ClearGlobalFields()
	})
}

func TestOverlog_Error(t *testing.T) {
	t.Run("error log", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		Log().Error("error log")

		expected := "{\"level\":\"error\",\"message\":\"error log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
	})

	t.Run("error log with mdc", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		SetGlobalFields([]string{"x-correlation-id"})
		MDC().Set("x-correlation-id", "1234")

		Log().Error("error log")

		expected := "{\"level\":\"error\",\"x-correlation-id\":\"1234\",\"message\":\"error log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
		ClearGlobalFields()
	})
}

func TestOverlog_Errorf(t *testing.T) {
	t.Run("errorf log", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		Log().Errorf("errorf %s",  "log")

		expected := "{\"level\":\"error\",\"message\":\"errorf log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
	})

	t.Run("errorf log with mdc", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		SetGlobalFields([]string{"x-correlation-id"})
		MDC().Set("x-correlation-id", "1234")

		Log().Errorf("errorf %s",  "log")

		expected := "{\"level\":\"error\",\"x-correlation-id\":\"1234\",\"message\":\"errorf log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
		ClearGlobalFields()
	})
}

func TestOverlog_Log(t *testing.T) {
	t.Run("no level log", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		Log().Log("no level log")

		expected := "{\"message\":\"no level log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
	})

	t.Run("no level log with mdc", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		SetGlobalFields([]string{"x-correlation-id"})
		MDC().Set("x-correlation-id", "1234")

		Log().Log("no level log")

		expected := "{\"x-correlation-id\":\"1234\",\"message\":\"no level log\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
		ClearGlobalFields()
	})
}

func TestOverlog_Logf(t *testing.T) {
	t.Run("logf log", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		Log().Logf("no level %s",  "logf")

		expected := "{\"message\":\"no level logf\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
	})

	t.Run("logf log with mdc", func(t *testing.T) {
		out := &bytes.Buffer{}
		New(zerolog.New(out))

		SetGlobalFields([]string{"x-correlation-id"})
		MDC().Set("x-correlation-id", "1234")

		Log().Logf("no level %s",  "logf")

		expected := "{\"x-correlation-id\":\"1234\",\"message\":\"no level logf\"}\n"
		actual := string(out.Bytes())
		assert.Equal(t, expected, actual)
		ClearGlobalFields()
	})
}