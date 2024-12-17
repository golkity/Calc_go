package main

import (
	"encoding/json"
	"fmt"
	"github.com/golkity/Calc_go/rpn/Errors"
	"io"
	"os"
	"runtime"
	"time"
)

type Config struct {
	DataFormat string   `json:"data_format"`
	Prefix     []string `json:"prefix"`
}

type Logger struct {
	out        io.Writer
	datePrefix string
	prefixes   [2]string
}

func NewFromConfig(out io.Writer, configPath string) (*Logger, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	if len(cfg.Prefix) != 2 {
		return nil, Errors.ErrInvalidPrefix
	}

	return &Logger{
		out:        out,
		datePrefix: cfg.DataFormat,
		prefixes:   [2]string{cfg.Prefix[0], cfg.Prefix[1]},
	}, nil
}

func (l *Logger) writeln(lvl int, a []any) {
	fmt.Fprintln(l.out, append([]any{time.Now().Format(l.datePrefix), l.prefixes[lvl]}, a...)...)
}

func (l *Logger) Println(a ...any) {
	l.writeln(0, a)
}

func (l *Logger) Fatalln(a ...any) {
	l.writeln(1, a)
	l.printStackTrace()
	os.Exit(1)
}

func (l *Logger) printStackTrace() {
	stackTrace := make([]uintptr, 10)
	length := runtime.Callers(2, stackTrace)
	for i := 0; i < length; i++ {
		function := runtime.FuncForPC(stackTrace[i] - 1)
		file, line := function.FileLine(stackTrace[i] - 1)
		fmt.Fprintf(l.out, "STACK TRACE: %s:%d %s\n", file, line, function.Name())
	}
}
