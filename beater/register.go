package beater

import (
	"io/ioutil"
	"os"
	"regexp"

	"github.com/elastic/beats/libbeat/logp"
)

const registerFile = "/data/local/tmp/register"

var logTime = regexp.MustCompile(`\s*([0-9]*)-([0-9]*)\s*([0-9]*):([0-9]*):([0-9]*)\.([0-9]{3})`)

func getRegisterTime() string {
	f, err := ioutil.ReadFile(registerFile)
	if err != nil {
		logp.Err(err.Error())
		return ""
	}
	return string(f)
}

func setRegisterTime(msg string) error {
	date := logTime.FindStringSubmatch(msg)
	f, err := os.OpenFile(registerFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0640)
	if err != nil {
		return err
	}
	defer f.Close()
	if len(date) != 0 {
		_, err = f.Write([]byte(date[0]))
		if err != nil {
			return err
		}
	}
	return nil
}
