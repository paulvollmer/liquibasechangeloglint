package main

import (
	"fmt"
	"io"
)

type Diagnostic struct {
	Summary string
	Detail  string
	Line    int
}

type Diagnostics []*Diagnostic

func (d Diagnostics) HasErrors() bool {
	if len(d) == 0 {
		return false
	}
	return true
}

func (d Diagnostics) Append(diag *Diagnostic) []*Diagnostic {
	return append(d, diag)
}

func (d Diagnostics) Extend(diags Diagnostics) []*Diagnostic {
	return append(d, diags...)
}

func (d Diagnostics) Print(w io.Writer) {
	for _, diagnostic := range d {
		fmt.Fprintf(w, "Summary: %s\nLine:    %v\nDetail:  %s\n\n", diagnostic.Summary, diagnostic.Line, diagnostic.Detail)
	}
}
