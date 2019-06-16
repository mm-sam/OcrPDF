package lib_test

import (
	"github.com/mm-sam/OcrPDF/lib"
	"path/filepath"
	"runtime"
	"testing"
)

func getLocalConfigPath(file string) string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), file)
}

func Test_NewConfig(t *testing.T) {
	conf_file := getLocalConfigPath("../config.json")
	err, conf := lib.NewConfig(conf_file)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	err = conf.Save()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	t.Log("PASS")
}
