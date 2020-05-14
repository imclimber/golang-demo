package main

import (
	"fmt"
	"path"
	"runtime"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// Logger  global logger
var Logger = logrus.New()

// InitLogger  set logger
func InitLogger() {
	// Only log the InfoLevel or above
	Logger.SetLevel(logrus.InfoLevel)
	if true {
		// log all the level
		Logger.SetLevel(logrus.TraceLevel)

		// set whether print caller
		Logger.SetReportCaller(true)
	}

	Logger.Formatter = &logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Dir(f.File) + "/" + path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}

	// logInfo output file
	// logInfoFileName := "entity-collection.log"
	// infoFile, err := os.OpenFile(logInfoFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	Logger.WithFields(logrus.Fields{
	// 		"fileName": logInfoFileName,
	// 	}).Info("failed to open logFile: ")
	// }
	// Output to stdout instead of the default stderr
	// MultiWriter
	// Logger.SetOutput(io.MultiWriter(infoFile, os.Stdout))

	// rotatelogs
	path := "./LogFiles/SysLog"
	errorPath := "./LogFiles/SysErrorLog"
	/*
		WithMaxAge 和 WithRotationCount二者只能设置一个
		WithMaxAge 设置文件清理前的最长保存时间
		WithRotationCount 设置文件清理前最多保存的个数
	*/
	// 每隔 1 天轮转一个新文件，保留最近 30 天的日志文件，多余自动清理
	writerSysLog, err := rotatelogs.New(
		path+"_%Y%m%d_%H%M.log",
		rotatelogs.WithLinkName(path),                            // 为最新的日志建立软连接，指向最新日志文件
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour), // 设置日志分割的时间，隔多久分割一次
		rotatelogs.WithMaxAge(time.Duration(30*24)*time.Hour),    // 设置文件清理前的最长保存时间
	)
	if err != nil {
		Logger.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	// 每隔 1 天轮转一个新文件，保留最近 30 天的日志文件，多余自动清理
	writerSysErrorLog, err := rotatelogs.New(
		errorPath+"_%Y%m%d_%H%M.log",
		rotatelogs.WithLinkName(path),                            // 为最新的日志建立软连接，指向最新日志文件
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour), // 设置日志分割的时间，隔多久分割一次
		rotatelogs.WithMaxAge(time.Duration(30*24)*time.Hour),    // 设置文件清理前的最长保存时间
	)
	if err != nil {
		Logger.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.TraceLevel: writerSysLog,
		logrus.DebugLevel: writerSysLog, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writerSysLog,
		logrus.WarnLevel:  writerSysLog,
		logrus.ErrorLevel: writerSysErrorLog,
		logrus.FatalLevel: writerSysLog,
		logrus.PanicLevel: writerSysLog,
	}, Logger.Formatter)
	Logger.AddHook(lfHook)
}

func main() {
	InitLogger()

	for {
		Logger.Traceln("hello, world! - trace")
		Logger.Info("hello, world!")
		time.Sleep(time.Duration(1) * time.Second)

		Logger.Errorln("hello, world! - error")
		time.Sleep(time.Duration(1) * time.Second)
	}
}

// 参考：https://blog.csdn.net/qianghaohao/article/details/104103717
// 参考：https://studygolang.com/articles/10534
