package clog

import (
	"github.com/apache/dubbo-kubernetes/operator/pkg/util/clog/log"
	"io"
	"os"
)

type ConsoleLogger struct {
	stdOut io.Writer
	stdErr io.Writer
	scope  *log.Scope
}

func (l *ConsoleLogger) LogAndError(v ...any) {
	if len(v) == 0 {
		return
	}
	s := fmt.Sprint(v...)
	l.PrintErr(s + "\n")
	l.scope.Infof(s)
}
func (l *ConsoleLogger) LogAndFatal(a ...any) {
	l.LogAndError(a...)
	os.Exit(-1)
}
func (l *ConsoleLogger) LogAndPrintf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	l.Print(s + "\n")
	l.scope.Infof(s)
}
func (l *ConsoleLogger) LogAndErrorf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	l.PrintErr(s + "\n")
	l.scope.Infof(s)
}

func (l *ConsoleLogger) LogAndFatalf(format string, a ...any) {
	l.LogAndErrorf(format, a...)
	os.Exit(-1)
}

type Logger interface {
	LogAndPrint(v ...any)
	LogAndError(v ...any)
	LogAndFatal(v ...any)
	LogAndPrintf(format string, a ...any)
	LogAndErrorf(format string, a ...any)
	LogAndFatalf(format string, a ...any)
}

func NewDefaultLogger() *ConsoleLogger {
	return nil
}

func (l *ConsoleLogger) LogAndPrint(v ...any) {
	return
}

func NewConsoleLogger(stdOut, stdErr io.Writer, scope *log.Scope) *ConsoleLogger {
	s := scope
	if s == nil {
		s = log.RegisterScope(log.DefaultScopeName, log.DefaultScopeName)
	}
	return &ConsoleLogger{
		stdOut: stdOut,
		stdErr: stdErr,
		scope:  s,
	}
}
