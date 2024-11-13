package gontenttype

import "testing"

func TestDetect(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want ContentType
	}{
		{
			name: "empty",
			args: args{content: ""},
			want: Unsupported,
		},
		{
			name: "random string defaults to csv",
			args: args{content: "lorem ipsum dolor sit amet"},
			want: CSV,
		},
		{
			name: "random string with commas, semicolons, tabs, and pipes defaults to csv",
			args: args{content: "foo,bar;baz\tqux|quux"},
			want: CSV,
		},
		{
			name: "csv comma separated",
			args: args{content: "foo,bar\nbaz,qux"},
			want: CSV,
		},
		{
			name: "json",
			args: args{content: "{\"foo\":\"bar\"}"},
			want: JSON,
		},
		{
			name: "xml",
			args: args{content: "<foo>bar</foo>"},
			want: XML,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Detect(tt.args.content); got != tt.want {
				t.Errorf("Detect() = %v, want %v", got, tt.want)
			}
		})
	}
}
