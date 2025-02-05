package configuration

import (
	"farhan/exception"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

// CustomLogger membungkus logrus.Logger dengan method helper untuk level log.
type CustomLogger struct {
	logger *logrus.Logger
}

// NewCustomLogger menginisialisasi logger dengan konfigurasi yang diinginkan.
func NewCustomLogger() *CustomLogger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	})

	// Pastikan direktori logs ada
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		if err := os.Mkdir("logs", 0770); err != nil {
			// Jika gagal membuat direktori, kita langsung log error di konsol dan keluar
			logger.Fatalf("Failed to create logs directory: %v", err)
		}
	}

	// Buka atau buat file log dengan format log_<tanggal>.log
	date := time.Now()
	logFile := "logs/log_" + date.Format("01-02-2006") + ".log"
	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		// Jika gagal membuka file, kita gunakan output default (stderr)
		logger.Infof("Failed to log to file, using default stderr: %v", err)
	} else {
		logger.SetOutput(file)
	}
	return &CustomLogger{logger: logger}
}

func (l *CustomLogger) Info(msg string, fields logrus.Fields) {
	l.logger.WithFields(fields).Info(msg)
}

func (l *CustomLogger) Warn(msg string, fields logrus.Fields) {
	l.logger.WithFields(fields).Warn(msg)
}

func (l *CustomLogger) Error(msg string, fields logrus.Fields) {
	l.logger.WithFields(fields).Error(msg)
}

func (l *CustomLogger) Critical(msg string, fields logrus.Fields) {
	// Gunakan Fatal untuk critical, yang akan mengakhiri program
	l.logger.WithFields(fields).Fatal(msg)
}

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyMsg:  "message",
		},
		PrettyPrint: true,
	})

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0770)
		exception.PanicLogging(err)
	}

	date := time.Now()
	file, err := os.OpenFile("logs/log_"+date.Format("01-02-2006")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	logger.SetOutput(file)
	return logger
}

func NewLoggerConfig() logger.Config {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0770)
		exception.PanicLogging(err)
	}

	date := time.Now()
	file, err := os.OpenFile("logs/log_"+date.Format("01-02-2006")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	exception.PanicLogging(err)

	return logger.Config{
		Output: file,
	}
}

/*
ref:
- https://dev.to/koddr/go-fiber-by-examples-working-with-middlewares-and-boilerplates-3p0m#explore-logging-middleware
*/
