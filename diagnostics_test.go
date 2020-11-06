package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagnostics_Append(t *testing.T) {
	diags := Diagnostics{}
	diags = diags.Append(&Diagnostic{"test", "append test", 1})
	assert.Len(t, diags, 1)
	assert.Equal(t, "test", diags[0].Summary)
	assert.Equal(t, "append test", diags[0].Detail)
	assert.Equal(t, 1, diags[0].Line)
}

func TestDiagnostics_Extend(t *testing.T) {
	diags := Diagnostics{}
	someDiags := Diagnostics{}
	someDiags = someDiags.Append(&Diagnostic{"test 1", "a test", 1})
	someDiags = someDiags.Append(&Diagnostic{"test 2", "an other test", 2})
	diags = diags.Extend(someDiags)
	assert.Len(t, diags, 2)
	assert.Equal(t, "test 1", diags[0].Summary)
	assert.Equal(t, "test 2", diags[1].Summary)
}

func TestDiagnostics_HasErrors(t *testing.T) {
	t.Run("has no errors", func(t *testing.T) {
		diags := Diagnostics{}
		assert.Equal(t, false, diags.HasErrors())
	})

	t.Run("has errors", func(t *testing.T) {
		diags := Diagnostics{}
		diags = diags.Append(&Diagnostic{"test", "append test", 1})
		assert.Equal(t, true, diags.HasErrors())
	})
}

func TestDiagnostics_Print(t *testing.T) {
	diags := Diagnostics{}
	diags = diags.Append(&Diagnostic{"test", "append test", 1})
	buf := bytes.Buffer{}
	diags.Print(&buf)
	assert.Equal(t, "Summary: test\nLine:    1\nDetail:  append test\n\n", buf.String())
}
