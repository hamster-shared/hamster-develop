package utils

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

// MoveToml Move.toml
type MoveToml struct {
	Package      `toml:"package"`
	Addresses    map[string]string `toml:"addresses"`
	Dependencies map[string]any    `toml:"dependencies"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (m *MoveToml) GetAddressField() []KeyValue {
	var keyValues []KeyValue
	for key, value := range m.Addresses {
		keyValues = append(keyValues, KeyValue{
			Key:   key,
			Value: value,
		})
	}
	return keyValues
}

type Package struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
}

// FillKeyValueToMoveToml 填充键值对到 Move.toml 的 Addresses 字段
func FillKeyValueToMoveToml(tomlPath string, keyValueString string) error {
	moveToml, err := ParseMoveToml(tomlPath)
	if err != nil {
		return err
	}

	keyValues, err := GetKeyValuesFromString(keyValueString)
	if err != nil {
		return err
	}

	for key, value := range keyValues {
		moveToml.Addresses[key] = value
	}

	return saveMoveToml(tomlPath, moveToml)
}

func ParseMoveTomlWithString(tomlString string) (*MoveToml, error) {
	var moveToml MoveToml
	_, err := toml.Decode(tomlString, &moveToml)
	if err != nil {
		return nil, err
	}
	return &moveToml, nil
}

func SaveMoveTomlString(moveToml *MoveToml) (string, error) {
	var buffer bytes.Buffer
	err := toml.NewEncoder(&buffer).Encode(moveToml)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func ParseMoveToml(tomlPath string) (*MoveToml, error) {
	var moveToml MoveToml
	_, err := toml.DecodeFile(tomlPath, &moveToml)
	if err != nil {
		return nil, err
	}
	return &moveToml, nil
}

func saveMoveToml(tomlPath string, moveToml *MoveToml) error {
	var buffer bytes.Buffer
	err := toml.NewEncoder(&buffer).Encode(moveToml)
	if err != nil {
		return err
	}
	return saveStringToFile(tomlPath, buffer.String())
}

func GetKeyValuesFromString(keyValueString string) (map[string]string, error) {
	var result = make(map[string]string)
	// 以逗号分隔 keyValueString
	// keyValueString := "0x1:0x1,0x2:0x2"
	keyValues := strings.Split(keyValueString, ",")
	// 以等号分隔键值
	for _, keyValue := range keyValues {
		keyValueArray := strings.Split(keyValue, "=")
		if len(keyValueArray) != 2 {
			return nil, fmt.Errorf("keyValueString format error: %s", keyValueString)
		}
		// 去掉可能存在的空格
		key := strings.TrimSpace(keyValueArray[0])
		value := strings.TrimSpace(keyValueArray[1])
		result[key] = value
	}
	return result, nil
}

func KeyValuesToString(keyValues []KeyValue) string {
	// 以逗号分隔 keyValueString, 注意最后一个逗号不要
	var result string
	for _, keyValue := range keyValues {
		result += fmt.Sprintf("%s=%s,", keyValue.Key, keyValue.Value)
	}
	return strings.TrimRight(result, ",")
}

func KeyValuesFromString(keyValues string) ([]KeyValue, error) {
	var result []KeyValue
	// 以逗号分隔 keyValueString
	// keyValueString := "0x1:0x1,0x2:0x2"
	keyValuesArray := strings.Split(keyValues, ",")
	// 以等号分隔键值
	for _, keyValue := range keyValuesArray {
		keyValueArray := strings.Split(keyValue, "=")
		if len(keyValueArray) != 2 {
			return nil, fmt.Errorf("keyValueString format error: %s", keyValues)
		}
		// 去掉可能存在的空格
		key := strings.TrimSpace(keyValueArray[0])
		value := strings.TrimSpace(keyValueArray[1])
		result = append(result, KeyValue{
			Key:   key,
			Value: value,
		})
	}
	return result, nil
}

// FileToHexString 读取文件内容，转换为 hex 字符串
func FileToHexString(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(data), nil
}
