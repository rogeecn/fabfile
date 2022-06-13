package fabfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFind(t *testing.T) {
	wd, _ := os.Getwd()

	type args struct {
		dst string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"test1", args{dst: "fabfile.go"}, filepath.Join(wd, "fabfile.go"), false},
		{"test1", args{dst: "http\\go.mod"}, "f:\\Projects\\http\\go.mod", false},
		{"test2", args{dst: "a_no_exist.go"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Find(tt.args.dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
