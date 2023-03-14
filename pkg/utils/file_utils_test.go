package utils

import (
	"fmt"
	"strings"
	"testing"
)

func Test_GetRedundantPath(t *testing.T) {
	shortPath := "/a/b/"
	longPath := "/a/b/c/d.txt"
	index := strings.Index(longPath, shortPath)
	fmt.Println(index)
	if index == 0 {
		relativePath := longPath[len(shortPath):]
		if relativePath[0] == '/' {
			fmt.Println(relativePath[1:])
		} else {
			fmt.Println(relativePath)
		}
	}
}

func Test_readStringFromFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				filePath: "./test_data/test.toml",
			},
			want: `[package]
name = 'mokshya-staking'
version = '1.0.0'


[addresses]
mokshyastaking = "_"
Std = "0x1"
aptos_std = "0x1"

[dependencies]
AptosFramework = { local = "../../aptos-core/aptos-move/framework/aptos-framework"}
AptosToken = { local = "../../aptos-core/aptos-move/framework/aptos-token" }
AptosStdlib = { local = "../../aptos-core/aptos-move/framework/aptos-stdlib" }`,
			wantErr: false,
		}, {
			name: "test file not exist",
			args: args{
				filePath: "./test_data/testtest.toml",
			},
			want:    ``,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readStringFromFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("readStringFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readStringFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
