package logger

import (
	"os"
	"github.com/op/go-logging"
)

var (

)

type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func init() {
	logLevel := os.Getenv("DA_LOG")
	if logLevel == "" {
		os.Setenv("DA_LOG", logging.INFO.String())
	}
}

func GetLogger() (*logging.Logger) {
	logger := logging.MustGetLogger("dalog")
	//backend:=logging.NewLogBackend(os.Stderr, "", 0)
	//formatter := logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} -▶ %{level:.4s} %{color:reset} %{message}`, )
	//backendFormatter := logging.NewBackendFormatter(backend, formatter)
	//backendLeveled := logging.AddModuleLevel(backend)
	//backendLeveled.SetLevel(logging.GetLevel(os.Getenv("DA_LOG")), "da")
	//logging.SetBackend(backend, backendFormatter)
	//logger.Debugf("Log Level: %s", os.Getenv("DA_LOG"))

	return logger;
}

