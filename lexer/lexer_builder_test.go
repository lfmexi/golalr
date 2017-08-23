package lexer

import (
	"reflect"
	"testing"

	"github.com/lfmexi/golalr/types"
)

func TestNewLexerBuilder(t *testing.T) {
	tests := []struct {
		name   string
		wanted Builder
	}{
		{"New Builder", Builder{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLexerBuilder()
			if &got == nil {
				t.Errorf("NewLexerBuilder() = %v", got)
			}
			if reflect.DeepEqual(got, tt.wanted) {
				t.Errorf("NewLexerBuilder() %v != %v", got, tt.wanted)
			}
		})
	}
}

func TestBuilder_SetFirstLine(t *testing.T) {
	type args struct {
		line int
	}
	tests := []struct {
		name string
		args args
	}{
		{"line 0", args{0}},
		{"line 1", args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewLexerBuilder()
			b.SetFirstLine(tt.args.line)
			if b.firstLine != tt.args.line {
				t.Errorf("firstLine %v = %v", b.firstLine, tt.args.line)
			}
		})
	}
}

func TestBuilder_SetFirstColumn(t *testing.T) {
	type args struct {
		line int
	}
	tests := []struct {
		name string
		args args
	}{
		{"line 0", args{0}},
		{"line 1", args{100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewLexerBuilder()
			b.SetFirstColumn(tt.args.line)
			if b.firstColumn != tt.args.line {
				t.Errorf("firstLine %v = %v", b.firstColumn, tt.args.line)
			}
		})
	}
}

func TestBuilder_BuildSimpleLexer(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		notWant Lexer
	}{
		{"New Simple Lexer", args{"abc"}, Lexer{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewLexerBuilder()
			got := b.Build(tt.args.input)
			if &got == nil {
				t.Errorf("Builder.Build() %v = nil", got)
			}
			if reflect.DeepEqual(got, tt.notWant) {
				t.Errorf("Builder.Build() = %v, want %v", got, tt.notWant)
			}
			if got.input != tt.args.input {
				t.Errorf("Builder.Build().input = %v, want %v", got, tt.args.input)
			}
		})
	}
}

func TestBuilder_ParseErrorOnBuild(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Did not recover")
		}
	}()
	builder := NewLexerBuilder()
	const aWrongDef types.SymbolType = "aWrongDef"
	builder.AddTerminalDefinition(aWrongDef, "abc*)", nil)
}

func TestBuilder_ErrorOnBuildingRange(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Did not recover")
		}
	}()
	builder := NewLexerBuilder()
	const aWrongDef types.SymbolType = "aWrongDef"
	builder.AddTerminalDefinition(aWrongDef, "a[z-a]|csd", nil)
}
