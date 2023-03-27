package utils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMoveToml(t *testing.T) {
	type args struct {
		tomlPath string
	}
	tests := []struct {
		name    string
		args    args
		want    *MoveToml
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				tomlPath: "./test_data/test.toml",
			},
			want: &MoveToml{
				Package: Package{
					Name:    "mokshya-staking",
					Version: "1.0.0",
				},
				Addresses: map[string]string{
					"Std":            "0x1",
					"aptos_std":      "0x1",
					"mokshyastaking": "_",
				},
				Dependencies: map[string]any{
					"AptosFramework": map[string]any{
						"local": "../../aptos-core/aptos-move/framework/aptos-framework",
					},
					"AptosToken": map[string]any{
						"local": "../../aptos-core/aptos-move/framework/aptos-token",
					},
					"AptosStdlib": map[string]any{
						"local": "../../aptos-core/aptos-move/framework/aptos-stdlib",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMoveToml(tt.args.tomlPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMoveToml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMoveToml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFillKeyValueToMoveToml(t *testing.T) {
	type args struct {
		tomlPath       string
		keyValueString string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				tomlPath:       "./test_data/test.toml",
				keyValueString: "Std=0x2",
			},
			wantErr: false,
		},
		{
			name: "test-2",
			args: args{
				tomlPath:       "./test_data/test.toml",
				keyValueString: "Std=0x1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FillKeyValueToMoveToml(tt.args.tomlPath, tt.args.keyValueString); (err != nil) != tt.wantErr {
				t.Errorf("FillKeyValueToMoveToml() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileToHexString(t *testing.T) {
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
				filePath: "./test_data/test.txt",
			},
			want:    "68656c6c6f20776f726c64",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileToHexString(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileToHexString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileToHexString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyValuesToString(t *testing.T) {
	keyValues := []KeyValue{
		{
			Key:   "Std",
			Value: "0x1",
		},
		{
			Key:   "AptosStdlib",
			Value: "0x2",
		},
	}
	want := "Std=0x1,AptosStdlib=0x2"
	got := KeyValuesToString(keyValues)
	if got != want {
		t.Errorf("KeyValuesToString() = %v, want %v", got, want)
	}
}

func TestKeyValuesFromString(t *testing.T) {
	keyValues := "Std=0x1,AptosStdlib=0x2"
	want := []KeyValue{
		{
			Key:   "Std",
			Value: "0x1",
		},
		{
			Key:   "AptosStdlib",
			Value: "0x2",
		},
	}
	got, err := KeyValuesFromString(keyValues)
	assert.NoError(t, err)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("KeyValuesFromString() = %v, want %v", got, want)
	}
}
