package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	//	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	data map[string](string)
}

func NewConfig() *Config {
	var config Config
	config.data = make(map[string](string))

	return &config
}

func (this *Config) LoadFile(file_path string) error {
	f, err := os.Open(file_path)
	if err != nil {
		return err
	}

	defer f.Close()

	buff := bufio.NewReader(f)

	line_number := 0
	for {
		b, _, err := buff.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		line := string(b)
		line = strings.TrimSpace(line)
		if len(line) <= 0 {
			line_number++
			continue
		}

		i := strings.IndexRune(line, '=')
		if i <= 0 {
			return errors.New("Line " + strconv.Itoa(line_number) + " error, [" + line + "]")
		}
		if len(line) < 3 {
			line_number++
			continue
		}
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "\\") {
			line_number++
			continue
		}
		key := strings.TrimSpace(line[0:i])
		value := strings.TrimSpace(line[i+1:])
		if len(key) <= 0 || len(value) <= 0 {
			return errors.New("Line " + strconv.Itoa(line_number) + " error, [" + line + "]")
		}

		this.data[key] = value
		line_number++
	}

	return nil
}

func (this *Config) GetString(key string) string {
	value, ok := this.data[key]
	if ok {
		return value
	}
	return ""
}

func (this *Config) GetInt16(key string) int16 {
	return int16(this.GetInt(key))
}

func (this *Config) GetInt(key string) int {
	value, ok := this.data[key]
	if ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return 0
		}
		return i
	}
	return 0
}

func (this *Config) GetInt32(key string) int32 {
	return int32(this.GetInt(key))
}

func (this *Config) GetInt64(key string) int64 {
	value, ok := this.data[key]
	if ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return 0
		}
		return i
	}
	return 0
}

func (this *Config) GetFloat32(key string) float32 {
	value, ok := this.data[key]
	if ok {
		i, err := strconv.ParseFloat(value, 32)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		return float32(i)
	}
	return 0
}

func (this *Config) GetFloat64(key string) float64 {
	value, ok := this.data[key]
	if ok {
		i, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		return i
	}
	return 0
}

func (this *Config) String() string {
	b, err := json.Marshal(this.data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func (this *Config) GetBool(key string) bool {
	value, ok := this.data[key]
	if ok && strings.ToLower(value) == "true" {
		return true
	}
	return false
}

func (this *Config) Unmarshal(object interface{}) error {
	myref := reflect.ValueOf(object).Elem()
	typeOfType := myref.Type()
	for i := 0; i < myref.NumField(); i++ {
		field := myref.Field(i)
		value := this.GetString(typeOfType.Field(i).Name)
		if value == "" {
			return errors.New("Can not find config " + typeOfType.Field(i).Name)
		}
		switch field.Type().Kind() {
		case reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
			field.SetInt(this.GetInt64(typeOfType.Field(i).Name))
		case reflect.Float32, reflect.Float64:
			if field.CanSet() {
				field.SetFloat(this.GetFloat64(typeOfType.Field(i).Name))
			}
		case reflect.Bool:
			if field.CanSet() {
				field.SetBool(this.GetBool(typeOfType.Field(i).Name))
			}
		case reflect.String:
			if field.CanSet() {
				field.SetString(value)
			}
		default:
			return errors.New("Unknow type " + field.Type().String() + " for " + typeOfType.Field(i).Name)
		}
	}
	return nil
}
