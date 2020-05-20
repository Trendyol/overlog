package over

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var overLogger *Overlog

type Overlog struct {
	log *zerolog.Logger
}

func New(logger zerolog.Logger) *Overlog {
	log.Logger = logger.Hook(MDCHook{})
	ResetGlobalMdcAdapter()
	ClearGlobalFields()

	overLogger = &Overlog{
		log: &log.Logger,
	}

	return overLogger
}

func NewDefault() *Overlog {
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger().Hook(MDCHook{})
	ResetGlobalMdcAdapter()
	ClearGlobalFields()

	zerolog.TimeFieldFormat = "2006-01-02 15:04:05,000"
	zerolog.LevelFieldMarshalFunc = UppercaseLevelEncoder()
	zerolog.CallerMarshalFunc = ShortCallerEncoder()
	zerolog.CallerFieldName = "source"

	overLogger = &Overlog{
		log: &log.Logger,
	}

	return overLogger
}

func Log() *Overlog {
	return overLogger
}

func Z() *zerolog.Logger {
	return overLogger.log
}

func (o *Overlog) Trace(args ...interface{}) {
	o.log.Trace().Msg(message("", args...))
}

func (o *Overlog) Tracef(format string, args ...interface{}) {
	o.log.Trace().Msg(message(format, args...))
}

func (o *Overlog) Debug(args ...interface{}) {
	o.log.Debug().Msg(message("", args...))
}

func (o *Overlog) Debugf(format string, args ...interface{}) {
	o.log.Debug().Msg(message(format, args...))
}

func (o *Overlog) Info(args ...interface{}) {
	o.log.Info().Msg(message("", args...))
}

func (o *Overlog) Infof(format string, args ...interface{}) {
	o.log.Info().Msg(message(format, args...))
}

func (o *Overlog) Warn(args ...interface{}) {
	o.log.Warn().Msg(message("", args...))
}

func (o *Overlog) Warnf(format string, args ...interface{}) {
	o.log.Warn().Msg(message(format, args...))
}

func (o *Overlog) Error(args ...interface{}) {
	o.log.Error().Msg(message("", args...))
}

func (o *Overlog) Errorf(format string, args ...interface{}) {
	o.log.Error().Msg(message(format, args...))
}

func (o *Overlog) Fatal(args ...interface{}) {
	o.log.Fatal().Msg(message("", args...))
}

func (o *Overlog) Fatalf(format string, args ...interface{}) {
	o.log.Fatal().Msg(message(format, args...))
}

func (o *Overlog) Panic(args ...interface{}) {
	o.log.Panic().Msg(message("", args...))
}

func (o *Overlog) Panicf(format string, args ...interface{}) {
	o.log.Panic().Msg(message(format, args...))
}

func (o *Overlog) Log(args ...interface{}) {
	o.log.Log().Msg(message("", args...))
}

func (o *Overlog) Logf(format string, args ...interface{}) {
	o.log.Log().Msg(message(format, args...))
}

func (o *Overlog) Print(args ...interface{}) {
	o.log.Print(args...)
}

func (o *Overlog) Printf(format string, args ...interface{}) {
	o.log.Printf(format, args...)
}

func message(format string, args ...interface{}) string {
	msg := format
	if msg == "" && len(args) > 0 {
		msg = fmt.Sprint(args...)
	} else if msg != "" && len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}
	return msg
}
