package lib

import (
	"github.com/apex/log"
	"github.com/hhrutter/pdfcpu/pkg/api"
	"github.com/hhrutter/pdfcpu/pkg/pdfcpu"
	"github.com/google/uuid"
)

func GetUUID() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.WithError(err)
		return ""
	}

	return uuid.String()
}

func SplitPage(srcPath string, outDir string) bool {
	config := pdfcpu.NewDefaultConfiguration()
	_, err := api.Process(api.SplitCommand(srcPath, outDir, 1, config))
	if err != nil {
		log.WithError(err)
		return false
	}

	return true
}
