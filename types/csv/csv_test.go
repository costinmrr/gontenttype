package csv

import "testing"

func TestIsCSV(t *testing.T) {
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
			name:    "simple string is csv",
			args:    args{content: "foo"},
			wantErr: false,
		},
		{
			name:    "csv comma separated",
			args:    args{content: "foo,bar\nbaz,qux"},
			wantErr: false,
		},
		{
			name:    "csv semicolon separated",
			args:    args{content: "foo;bar\nbaz;qux"},
			wantErr: false,
		},
		{
			name:    "csv tab separated",
			args:    args{content: "foo\tbar\nbaz\tqux"},
			wantErr: false,
		},
		{
			name:    "csv pipe separated",
			args:    args{content: "foo|bar\nbaz|qux"},
			wantErr: false,
		},
		{
			name:    "csv with quotes",
			args:    args{content: "\"foo\",\"bar\"\n\"baz\",\"qux\""},
			wantErr: false,
		},
		{
			name:    "csv with quotes and commas",
			args:    args{content: "\"foo,bar\",\"baz,qux\"\n\"foo,bar\",\"baz,qux\""},
			wantErr: false,
		},
		{
			name:    "different number of columns per line",
			args:    args{content: "foo,bar\nbaz"},
			wantErr: true,
		},
		{
			name:    "empty lines",
			args:    args{content: "foo,bar\n\nbaz,qux"},
			wantErr: false,
		},
		{
			name:    "different separators per line",
			args:    args{content: "foo,bar\nbaz;qux"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsCSV(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("IsCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
