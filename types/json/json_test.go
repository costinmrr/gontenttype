package json

import "testing"

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
			name:    "json",
			args:    args{content: `{"foo": "bar"}`},
			wantErr: false,
		},
		{
			name:    "json with spaces",
			args:    args{content: ` { "foo": "bar" } `},
			wantErr: false,
		},
		{
			name: "json with newlines",
			args: args{content: `{
"foo": "bar"
}`},
			wantErr: false,
		},
		{
			name: "json with tabs",
			args: args{content: `{
				"foo": "bar"
			}`},
			wantErr: false,
		},
		{
			name:    "json with other chars at the end",
			args:    args{content: `{"foo": "bar"}!`},
			wantErr: true,
		},
		{
			name:    "json with other chars at the end 2",
			args:    args{content: `{"foo": "bar"}>>`},
			wantErr: true,
		},
		{
			name:    "json with other chars at the end 3",
			args:    args{content: `{"foo": "bar"}}}`},
			wantErr: true,
		},
		{
			name:    "json with other chars at the end 4",
			args:    args{content: `{"foo": "bar"}{"baz": "qux"}`},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsJSON(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("IsJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
