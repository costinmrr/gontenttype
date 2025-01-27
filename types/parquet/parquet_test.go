package parquet

import (
	"os"
	"testing"
)

func TestIsParquet_ValidFile(t *testing.T) {
	content1, err := os.ReadFile("../../testdata/weather.parquet")
	if err != nil {
		t.Fatal(err)
	}

	err = IsParquet(string(content1))
	if err != nil {
		t.Errorf("IsParquet() error = %v, wantErr %v", err, false)
	}

	content2, err := os.ReadFile("../../testdata/flights-1m.parquet")
	if err != nil {
		t.Fatal(err)
	}

	err = IsParquet(string(content2))
	if err != nil {
		t.Errorf("IsParquet() error = %v, wantErr %v", err, false)
	}
}

func TestIsParquet_Invalid(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr bool
	}{
		{
			name:    "empty",
			content: "",
			wantErr: true,
		},
		{
			name:    "invalid - no magic sequence",
			content: "invalid",
			wantErr: true,
		},
		{
			name:    "invalid - only beginning magic sequence",
			content: "PAR1",
			wantErr: true,
		},
		{
			name:    "invalid - invalid footer",
			content: "PAR1\ncontent\nPAR1",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsParquet(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsParquet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
