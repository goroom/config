package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	//	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	config_file_path string
	data             map[string](string)
}

func NewConfig(config_file_path string) (*Config, error) {
	var config Config
	config.config_file_path = config_file_path
	config.data = make(map[string](string))

	err := config.Load()
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (this *Config) Load() error {
	f, err := os.Open(this.config_file_path)
	if err != nil {
		return err
	}
	defer f.Close()

	buff := bufio.NewReader(f)

	line_number := 0
	for {
		b, _, err := buff.ReadLine()
		if err != nil || io.EOF == err {
			//			fmt.Println(err)
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

		//		fmt.Println("[" + key + "][" + value + "]")

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
	value, ok := this.data[key]
	if ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return 0
		}
		return int16(i)
	}
	return 0
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
		return err.Error()
	}
	return string(b)
}
