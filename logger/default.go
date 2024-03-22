package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	dlog "github.com/geniussheep/ymir/debug/log"
)

var (
	LoggerKey  = "_ymir-logger-request"
	TrafficKey = "X-Request-Id"
)

func init() {
	lvl, err := GetLevel(os.Getenv("YMIR_LOG_LEVEL"))
	if err != nil {
		lvl = InfoLevel
	}

	DefaultLogger = NewHelper(NewLogger(WithLevel(lvl)))
}

type defaultLogger struct {
	sync.RWMutex
	opts Options
}

func copyFields(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// logCallerfilePath returns a package/file:line description of the caller,
// preserving only the leaf directory name and file name.
func logCallerfilePath(loggingFilePath string) string {
	// To make sure we trim the path correctly on Windows too, we
	// counter-intuitively need to use '/' and *not* os.PathSeparator here,
	// because the path given originates from Go stdlib, specifically
	// runtime.Caller() which (as of Mar/17) returns forward slashes even on
	// Windows.
	//
	// See https://github.com/golang/go/issues/3335
	// and https://github.com/golang/go/issues/18151
	//
	// for discussion on the issue on Go side.
	idx := strings.LastIndexByte(loggingFilePath, '/')
	if idx == -1 {
		return loggingFilePath
	}
	idx = strings.LastIndexByte(loggingFilePath[:idx], '/')
	if idx == -1 {
		return loggingFilePath
	}
	return loggingFilePath[idx+1:]
}

func (l *defaultLogger) log(level Level, format string, fmtArg []interface{}, cfields []interface{}) {
	// TODO decide does we need to write message if log level not used?
	if !l.opts.Level.Enabled(level) {
		return
	}

	l.RLock()
	fields := copyFields(l.opts.Fields)
	l.RUnlock()

	fields["level"] = level.String()

	if _, file, line, ok := runtime.Caller(l.opts.CallerSkipCount); ok {
		fields["file"] = fmt.Sprintf("%s:%d", logCallerfilePath(file), line)
	}

	rec := dlog.Record{
		Timestamp: time.Now(),
		Metadata:  make(map[string]string, len(fields)),
	}
	message := ""
	if format == "" {
		for _, fa := range fmtArg {
			message += fmt.Sprintf("%v\t", fa)
		}
	} else {
		message = fmt.Sprintf(format, fmtArg...)
	}
	rec.Message = message

	metadata := ""

	keys := make([]string, 0)
	for k, v := range fields {
		keys = append(keys, k)
		rec.Metadata[k] = fmt.Sprintf("%v", v)
	}

	sort.Strings(keys)

	for i, k := range keys {
		if i == 0 {
			metadata += fmt.Sprintf("%s:%v\t", k, fields[k])
		} else {
			metadata += fmt.Sprintf(" %s:%v\t", k, fields[k])
		}
	}

	ckeys := make([]string, 0)
	for i := 0; i < len(cfields); {
		if i >= len(cfields)-1 {
			break
		}
		vi := i + 1
		if vi > len(cfields)+1 {
			vi = i
		}
		k, v := fmt.Sprintf("%s", cfields[i]), cfields[vi]
		ckeys = append(ckeys, k)
		rec.Metadata[k] = fmt.Sprintf("%v", v)
		i += 2
	}

	sort.Strings(ckeys)

	for i, k := range ckeys {
		if i == 0 {
			metadata += fmt.Sprintf("%s:%v\t", k, rec.Metadata[k])
		} else {
			metadata += fmt.Sprintf(" %s:%v\t", k, rec.Metadata[k])
		}
	}

	var name string
	if l.opts.Name != "" {
		name = "[" + l.opts.Name + "]"
	}
	t := rec.Timestamp.Format("2006-01-02 15:04:05.000Z0700")
	//fmt.Printf("%s\n", t)
	//fmt.Printf("%s\n", name)
	//fmt.Printf("%s\n", metadata)
	//fmt.Printf("%v\n", rec.Message)
	logStr := ""
	if name == "" {
		logStr = fmt.Sprintf("%s %s %v\n", t, metadata, rec.Message)
	} else {
		logStr = fmt.Sprintf("%s %s %s %v\n", name, t, metadata, rec.Message)
	}
	_, err := l.opts.Out.Write([]byte(logStr))
	if err != nil {
		log.Printf("log [Logf] write error: %s \n", err.Error())
	}
}

// Init (opts...) should only overwrite provided options
func (l *defaultLogger) Init(opts ...Option) error {
	for _, o := range opts {
		o(&l.opts)
	}
	return nil
}

func (l *defaultLogger) String() string {
	return "default"
}

func (l *defaultLogger) Fields(fields map[string]interface{}) Logger {
	l.Lock()
	l.opts.Fields = copyFields(fields)
	l.Unlock()
	return l
}

func (l *defaultLogger) Log(level Level, v ...interface{}) {
	l.log(level, "", v, nil)
}

func (l *defaultLogger) Logf(level Level, format string, v ...interface{}) {
	l.log(level, format, v, nil)
}

func (l *defaultLogger) Logw(level Level, msg string, f ...interface{}) {
	l.log(level, msg, nil, f)
}

func (l *defaultLogger) Options() Options {
	// not guard against options Context values
	l.RLock()
	opts := l.opts
	opts.Fields = copyFields(l.opts.Fields)
	l.RUnlock()
	return opts
}

// NewLogger builds a new logger based on options
func NewLogger(opts ...Option) Logger {
	// Default options
	options := Options{
		Level:           InfoLevel,
		Fields:          make(map[string]interface{}),
		Out:             os.Stderr,
		CallerSkipCount: 3,
		Context:         context.Background(),
		Name:            "",
	}

	l := &defaultLogger{opts: options}
	if err := l.Init(opts...); err != nil {
		l.Log(FatalLevel, err)
	}

	return l
}
