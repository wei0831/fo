package utli

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"time"

	homedir "github.com/mitchellh/go-homedir"
)

var log = getNewLogger()

type logger struct {
	logFolder, logFileName string
	fileMode               os.FileMode
}

func getNewLogger() logger {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logFolder := path.Join(home, "folog")
	return logger{logFolder, "", 0755}
}

func (l *logger) init() {
	l.logFileName = fmt.Sprintf("%d.log", time.Now().UnixNano())
	if err := os.MkdirAll(l.logFolder, l.fileMode); err != nil {
		fmt.Printf("[ERR] %s\n", err)
	}
}

func (l *logger) getFileName() string {
	if l.logFileName == "" {
		l.init()
	}
	return l.logFileName
}

func (l *logger) info(action string) {
	if l.logFileName == "" {
		l.init()
	}
	f, err := os.OpenFile(path.Join(l.logFolder, l.logFileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, l.fileMode)
	if err != nil {
		fmt.Printf("[ERR] %s\n", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	w := bufio.NewWriter(f)
	fmt.Fprintln(w, action)
	w.Flush()
}
