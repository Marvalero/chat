package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("New should not return nil")
		return
	}

	tracer.Trace("HEllo trace package")
	if buf.String() != "HEllo trace package\n" {
		t.Errorf("Trace should not write '%s'.", buf.String())
	}
}
