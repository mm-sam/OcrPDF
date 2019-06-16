package lib_test

import (
	"github.com/mm-sam/OcrPDF/lib"
	"os"
	"testing"
)

func Test_SplitPage(t *testing.T) {
	pdfFile := getLocalConfigPath("..//temp/test.pdf")
	outDir := getLocalConfigPath("../temp/" + lib.GetUUID())
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		os.MkdirAll(outDir, 0777)
	}
	defer func() {
		os.RemoveAll(outDir)
	}()
	t.Log(outDir)

	r := lib.SplitPage(pdfFile, outDir)
	if r {
		t.Log("PASS")
		return
	}
	t.Fail()
}

func Test_GetUUID(t *testing.T)  {
	uuid := lib.GetUUID()

	t.Log(uuid)

	t.Log("PASS")
}