package logger

import (
	"os"
)

type Helper struct {
	Logger
	fields map[string]interface{}
}

func NewHelper(log Logger) *Helper {
	return &Helper{Logger: log}
}

func (h *Helper) Info(args ...interface{}) {
	h.Log(InfoLevel, args)
}

func (h *Helper) Infof(format string, args ...interface{}) {
	h.Logf(InfoLevel, format, args...)
}

func (h *Helper) Infow(fields map[string]interface{}, format string, args ...interface{}) {
	h.Logw(InfoLevel, fields, format, args...)
}

func (h *Helper) Trace(args ...interface{}) {
	h.Log(TraceLevel, args...)
}

func (h *Helper) Tracef(format string, args ...interface{}) {
	h.Logf(TraceLevel, format, args...)
}

func (h *Helper) Tracew(fields map[string]interface{}, format string, args ...interface{}) {
	h.Logw(TraceLevel, fields, format, args...)
}

func (h *Helper) Debug(args ...interface{}) {
	h.Log(DebugLevel, args...)
}

func (h *Helper) Debugf(format string, args ...interface{}) {
	h.Logf(DebugLevel, format, args...)
}

func (h *Helper) Debugw(fields map[string]interface{}, format string, args ...interface{}) {
	h.Logw(DebugLevel, fields, format, args...)
}

func (h *Helper) Warn(args ...interface{}) {
	h.Log(WarnLevel, args...)
}

func (h *Helper) Warnf(format string, args ...interface{}) {
	h.Logf(WarnLevel, format, args...)
}

func (h *Helper) Warnw(fields map[string]interface{}, format string, args ...interface{}) {
	h.Logw(WarnLevel, fields, format, args...)
}

func (h *Helper) Error(args ...interface{}) {
	h.Log(ErrorLevel, args...)
}

func (h *Helper) Errorf(format string, args ...interface{}) {
	h.Logf(ErrorLevel, format, args...)
}

func (h *Helper) Errorw(fields map[string]interface{}, format string, args ...interface{}) {
	h.Logw(ErrorLevel, fields, format, args...)
}

func (h *Helper) Fatal(args ...interface{}) {
	h.Log(FatalLevel, args...)
	os.Exit(1)
}

func (h *Helper) Fatalf(format string, args ...interface{}) {
	h.Logf(FatalLevel, format, args...)
	os.Exit(1)
}

func (h *Helper) Fatalw(fields map[string]interface{}, format string, args ...interface{}) {
	h.Logw(FatalLevel, fields, format, args...)
	os.Exit(1)
}

func (h *Helper) WithError(err error) *Helper {
	fields := copyFields(h.fields)
	fields["error"] = err
	return &Helper{Logger: h.Logger, fields: fields}
}

func (h *Helper) WithFields(fields map[string]interface{}) *Helper {
	nfields := copyFields(fields)
	for k, v := range h.fields {
		nfields[k] = v
	}
	return &Helper{Logger: h.Logger, fields: nfields}
}

func (h *Helper) Log(level Level, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level) {
		return
	}
	h.Logger.Fields(h.fields).Log(level, args...)
}

func (h *Helper) Logf(level Level, format string, v ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level) {
		return
	}
	h.Logger.Fields(h.fields).Logf(level, format, v...)
}

// Logf writes a msg log entry with some custom field
func (h *Helper) Logw(level Level, fileds map[string]interface{}, format string, v ...interface{}) {
	if !h.Logger.Options().Level.Enabled(level) {
		return
	}
	h.Logger.Fields(h.fields).Logw(level, fileds, format, v...)
}
