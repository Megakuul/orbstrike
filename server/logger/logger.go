package logger

import (
	"io"
	"log"
	"os"
	"strings"
	"runtime"
	"path/filepath"
)

var (
	ERROR bool
	WARNING bool
	INFORMATION bool
)

func InitLogger(logname string, loglevel string, maxlogsize int, logconsole bool) error {

	file, err := os.OpenFile(logname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	err = clearLog(maxlogsize, file)
	if err != nil {
		return err
	}

	var mOutWriter io.Writer
	if logconsole {
		mOutWriter = io.MultiWriter(os.Stdout, file)
	} else {
		mOutWriter = io.MultiWriter(file)
	}

	log.SetOutput(mOutWriter)

	logleverUpper := strings.ToUpper(loglevel)

	ERROR = strings.Contains(logleverUpper, "ERROR")
	WARNING = strings.Contains(logleverUpper, "WARNING")
	INFORMATION = strings.Contains(logleverUpper, "INFORMATION")

	return nil
}

func clearLog(maxLogSize int, file *os.File) error {
	fileInfo, err := file.Stat()
	if err != nil {
		file.Close()
		return err
	}

	if fileInfo.Size() > int64(maxLogSize)*1024 {
		err = file.Truncate(0)
		if err != nil {
			file.Close()
			return err
		}

		_, err = file.Seek(0, 0)
		if err != nil {
			file.Close()
			return err
		}
	}

	return nil
}

func WriteInformationLogger(format string, v ...interface{}) {
	if INFORMATION {
		log.Printf("\n[ ORBSTRIKE Information ]:\n"+format+"\n\n", v...)
	}
}

func WriteWarningLogger(err error) {
	if WARNING {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			log.Printf("\n[ ORBSTRIKE Warning ]:\n%s\n(file: %s, line: %d)\n\n",
				err, filepath.Base(file), line)
		} else {
			log.Printf("\n[ ORBSTRIKE Warning ]:\n%s\n\n", err)
		}
	}
}

func WriteErrLogger(err error) {
	if ERROR {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			log.Printf("\n[ ORBSTRIKE Panic ]:\n%s\n(file: %s, line: %d)\n\n",
				err, filepath.Base(file), line)
		} else {
			log.Printf("\n[ ORBSTRIKE Panic ]:\n%s\n\n", err)
		}
	}
}
