package gontenttype

import (
	"os"
	"testing"
)

func TestDetect(t *testing.T) {
	parquetContentBytes, _ := os.ReadFile("./testdata/weather.parquet")

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
			name: "json lines",
			args: args{content: "{\"foo\":\"bar\"}\n{\"baz\":\"qux\"}\n{\"quux\":\"quuz\"}"},
			want: JSONLines,
		},
		{
			name: "xml",
			args: args{content: "<foo>bar</foo>"},
			want: XML,
		},
		{
			name: "parquet",
			args: args{content: string(parquetContentBytes)},
			want: Parquet,
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

// Benchmarks
func BenchmarkDetectJSON_SimpleString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Detect("{\"foo\":\"bar\"}")
	}
}

func BenchmarkDetectJSON_1KB(b *testing.B) {
	content, _ := os.ReadFile("testdata/json_1KB.json")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectJSON_100KB(b *testing.B) {
	content, _ := os.ReadFile("testdata/json_100KB.json")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectJSON_1MB(b *testing.B) {
	content, _ := os.ReadFile("testdata/json_1MB.json")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectJSON_10MB(b *testing.B) {
	content, _ := os.ReadFile("testdata/json_10MB.json")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectJSONLines_SimpleString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Detect("{\"foo\":\"bar\"}\n{\"baz\":\"qux\"}")
	}
}

func BenchmarkDetectJSONLines_1KB(b *testing.B) {
	content, _ := os.ReadFile("testdata/json_lines_1KB.json")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectJSONLines_100KB(b *testing.B) {
	content, _ := os.ReadFile("testdata/json_lines_100KB.json")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectJSONLines_1MB(b *testing.B) {
	content, _ := os.ReadFile("testdata/json_lines_1MB.json")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectJSONLines_10MB(b *testing.B) {
	content, _ := os.ReadFile("testdata/json_lines_10MB.json")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectXML_SimpleString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Detect("<foo>bar</foo>")
	}
}

func BenchmarkDetectXML_1KB(b *testing.B) {
	content, _ := os.ReadFile("testdata/xml_1KB.xml")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectXML_100KB(b *testing.B) {
	content, _ := os.ReadFile("testdata/xml_100KB.xml")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectXML_1MB(b *testing.B) {
	content, _ := os.ReadFile("testdata/xml_1MB.xml")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectXML_10MB(b *testing.B) {
	content, _ := os.ReadFile("testdata/xml_10MB.xml")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectCSV_SimpleString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Detect("foo,bar\nbaz,qux")
	}
}

func BenchmarkDetectCSV_1KB(b *testing.B) {
	content, _ := os.ReadFile("testdata/csv_1KB.csv")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectCSV_100KB(b *testing.B) {
	content, _ := os.ReadFile("testdata/csv_100KB.csv")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectCSV_1MB(b *testing.B) {
	content, _ := os.ReadFile("testdata/csv_1MB.csv")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectCSV_10MB(b *testing.B) {
	content, _ := os.ReadFile("testdata/csv_10MB.csv")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectUnsupported_SimpleString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Detect("lorem ipsum dolor sit amet")
	}
}

func BenchmarkDetectParquet_16KB(b *testing.B) {
	content, _ := os.ReadFile("testdata/weather.parquet")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}

func BenchmarkDetectParquet_6point5MB(b *testing.B) {
	content, _ := os.ReadFile("testdata/flights-1m.parquet")
	contentString := string(content)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(contentString)
	}
}
