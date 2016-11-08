package main

import (
	"github.com/goroom/config"
	"github.com/goroom/logger"
)

type TestStruct struct {
	A int
	B int16
	C int32
	D int64
	E float32
	F float64
	G string
	H bool
}

func main() {
	logger_config := logger.NewDefaultConfig()
	logger_config.Level = logger.OFF
	logger_config.ConsoleLevel = logger.ALL
	log, _ := logger.NewLogger(logger_config)
	logger.SetDefaultLogger(log)

	//初始化
	cf := config.NewConfig()
	//从文件读取配置
	err := cf.LoadFile("./config.ini")
	if err != nil {
		logger.Error(err)
		return
	}

	//打印配置
	logger.Debug(cf.GetInt("A"))
	logger.Debug(cf.GetInt16("B"))
	logger.Debug(cf.GetInt32("C"))
	logger.Debug(cf.GetInt64("D"))
	logger.Debug(cf.GetFloat32("E"))
	logger.Debug(cf.GetFloat64("F"))
	logger.Debug(cf.GetString("G"))
	logger.Debug(cf.GetBool("H"))
	logger.Debug(cf)

	//解析到结构体
	var ts TestStruct
	err = cf.Unmarshal(&ts)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Debug("ts:", ts)

	//不初始化直接将文件解析到结构体
	var ts2 TestStruct
	err = config.FileUnmarshal("./config.ini", &ts2)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Debug("ts2:", ts2)
}
