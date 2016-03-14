package inutil

import (
	"path/filepath"
	"flag"
)

var workingDirPath string = ""

type WorkingDir struct {
	Path string
}

func WorkingDirPath(subDir string) string {
	return filepath.Join(workingDirPath, subDir)
}

func GetWorkingDir() *WorkingDir {
	return &WorkingDir{workingDirPath}
}

func init() {
	workingDirPath = *flag.String("working_dir", "", "set working_dir")
}