package logging

import (
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("This is a test log message.")
	fmt.Println("Logger test completed.")
}

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is a Trace")
	logger.Debug("This is a Debug")
	logger.Info("This is an Info")
	logger.Warn("This is a Warn")
	logger.Error("This is an Error")
}

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("logrus_app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("This log message will be written to the file.")
	logger.Warn("This is a warning message in the file.")
	logger.Error("This is an error message in the file.")
}

func TestFormatter(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("logrus_app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("This log message will be in JSON format.")
	logger.Warn("This log message will also be in JSON format.")
	logger.Error("This log message will be in JSON format as well.")
}

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("logrus_app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.WithField("username", "iqbal").Info("This log message will be in JSON format.")
	logger.WithField("username", "iqball").WithField("name", "iqbal").Warn("This log message will also be in JSON format.")
}

func TestWithFields(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("logrus_app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.WithFields(logrus.Fields{
		"username": "ciaa",
		"name":     "cia",
	}).Info("This log message will be in JSON format.")
}

type MyHook struct{}

func (h *MyHook) Levels() []logrus.Level {
	//return logrus.AllLevels
	return []logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel}
}

func (h *MyHook) Fire(entry *logrus.Entry) error {
	fmt.Println("MyHook fired for entry:", entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()

	logger.AddHook(&MyHook{})

	logger.Info("This is an info message.")
	logger.Warn("This is a warning message.")
	logger.Error("This is an error message.")
}
