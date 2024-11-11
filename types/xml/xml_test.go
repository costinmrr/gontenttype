package xml

import (
	"errors"
	"testing"
)

func TestIsXML(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		inheritedErr bool
		err          error
	}{
		{
			name:    "empty",
			args:    args{content: ""},
			wantErr: true,
			err:     ErrEmptyContent,
		},
		{
			name:    "string not xml",
			args:    args{content: "foo"},
			wantErr: true,
			err:     ErrRootNotFound,
		},
		{
			name:         "string not xml starting with <",
			args:         args{content: "<foo"},
			wantErr:      true,
			inheritedErr: true,
		},
		{
			name:    "xml",
			args:    args{content: "<foo>bar</foo>"},
			wantErr: false,
		},
		{
			name:    "xml with spaces at the end",
			args:    args{content: "<foo>bar</foo> "},
			wantErr: false,
		},
		{
			name:    "xml with spaces at the beginning",
			args:    args{content: " <foo>bar</foo>"},
			wantErr: false,
		},
		{
			name:    "xml with other chars at the end",
			args:    args{content: "<foo>bar</foo>>>>"},
			wantErr: true,
			err:     ErrContentAfterRoot,
		},
		{
			name:    "xml with other chars at the end 2",
			args:    args{content: "<foo>bar</foo><>"},
			wantErr: true,
			err:     ErrContentAfterRoot,
		},
		{
			name:    "xml with other chars at the end 3",
			args:    args{content: `<foo>bar</foo>{"a": "b"}`},
			wantErr: true,
			err:     ErrContentAfterRoot,
		},
		{
			name:    "complex xml",
			args:    args{content: "<foo><bar>baz</bar></foo>"},
			wantErr: false,
		},
		{
			name:    "very complex xml",
			args:    args{content: `<?xml version="1.0" encoding="UTF-8"?><foo bar="baz"><bar>baz</bar></foo>`},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsXML(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.inheritedErr {
				// If the error is inherited, we don't need to check the error
				return
			}
			if !errors.Is(err, tt.err) {
				t.Errorf("IsXML() error = %v, wantErr %v", err, tt.err)
				return
			}
		})
	}
}
