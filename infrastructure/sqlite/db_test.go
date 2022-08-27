package sqlite

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestConnect(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
	}{
		{"runs", args{fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Connect(tt.args.file)
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tt.args.file)

			if got == nil {
				t.Errorf("Connect() = %v, want %v", got, nil)
			}

			err = got.Close()
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
