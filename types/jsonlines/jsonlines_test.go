package jsonlines

import (
	"testing"
)

func TestIsJSON(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "empty",
			args:    args{content: ""},
			wantErr: true,
		},
		{
			name:    "string not json",
			args:    args{content: "foo"},
			wantErr: true,
		},
		{
			name:    "one line valid",
			args:    args{content: `{"foo": "bar"}`},
			wantErr: false,
		},
		{
			name: "multiple lines valid with no new line at the end",
			args: args{content: `{"foo": "bar"}
{"baz": "qux"}
{"quux": "quuz"}`},
			wantErr: false,
		},
		{
			name: "multiple lines valid with new line at the end",
			args: args{content: `{"foo": "bar"}
{"baz": "qux"}
{"quux": "quuz"}
`},
			wantErr: false,
		},
		{
			name: "multiple lines valid with escaped new line character",
			args: args{content: "{\"foo\\n\": \"bar\"}\n" +
				"{\"baz\": \"qux\\n\"}\n" +
				"{\"quux\": \"quuz\"}\n"},
			wantErr: false,
		},
		{
			name: "multiple lines invalid with one line missing a closing bracket",
			args: args{content: `{"foo": "bar"}
{"baz": "qux"
{"quux": "quuz"}`},
			wantErr: true,
		},
		{
			name: "multiple lines invalid with one line containing a newline character",
			args: args{content: "{\"foo\": \"bar\"}\n" +
				"{\"baz\": \"qux\n\"}\n" +
				"{\"quux\": \"quuz\"}\n"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsJSONLines(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("IsJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
