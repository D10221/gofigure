package inutil

import (
	"testing"
	"os"
	"path/filepath"
	"io/ioutil"
)

const testdata = "testdata"

func Test_read(t *testing.T) {

	dir := WorkingDirPath(testdata)

	if !exists(dir) {
		t.Error("dir doesn't exists %v", dir)
	}
}

func Test_Current_Dir(t *testing.T) {
	path := filepath.Dir(os.Args[0])
	path, e := filepath.Abs(path)
	list, e := ioutil.ReadDir(path)
	if e != nil {
		t.Error(e)
	}
	for _, item := range list {
		t.Log(item)
	}
}

func Test_TestData(t *testing.T) {

	dir := WorkingDirPath(testdata)

	json, err := ioutil.ReadFile(filepath.Join(dir, "data.json"))
	if err != nil {
		t.Error(err)
	}
	if (string(json) != "{\"key\" : \"value\"}") {
		t.Error("wrong")
	}

	t.Log("run ... ")
}

func exists(path string) bool {
	_, err := os.Stat(path);
	return err == nil
}
