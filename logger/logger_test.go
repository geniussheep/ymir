package logger

import (
	"context"
	"fmt"
	"testing"
)

func TestLogger(t *testing.T) {
	l := NewLogger(WithLevel(TraceLevel), WithName("zap"))
	h1 := NewHelper(l).WithFields(map[string]interface{}{"key1": "val1"})
	h1.Trace("trace_msg1")
	h1.Warn("warn_msg1")

	h2 := NewHelper(l).WithFields(map[string]interface{}{"key2": "val2"})
	h2.Trace("trace_msg2")
	h2.Warn("warn_msg2")

	h3 := NewHelper(l).WithFields(map[string]interface{}{"key3": "val4"})
	h3.Info("test_msg")
	ctx := context.TODO()
	ctx = context.WithValue(ctx, &loggerKey{}, h3)
	v := ctx.Value(&loggerKey{})
	ll := v.(*Helper)
	ll.Info("test_msg")
	ll.Infof("test_msg:%s", "args1")
	ll.Logw(InfoLevel, fmt.Sprintf("test_msg:%s", "arg"),
		"test-key", "test-value")
	ll.Logw(InfoLevel, "test_msg_value",
		"test-key", "test-value",
		"test-kv")
}
