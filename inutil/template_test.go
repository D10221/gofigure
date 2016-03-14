package inutil

import (
	"testing"
	"io/ioutil"
	"path/filepath"
	"html/template"
	"bytes"
)

func Test_Templating(t *testing.T) {
	wd := GetWorkingDir()
	templateData, e := ioutil.ReadFile(filepath.Join(wd.Path, "testdata", "template.html"))
	if e != nil {
		t.Error(e)
	}
	tt := template.Must(template.New("?").Parse(string(templateData)))

	templateContext := struct{ Name string }{"me"}

	writer := new(bytes.Buffer)
	e = tt.Execute(writer, templateContext)
	if e != nil {
		t.Error(e)
	}
	if string(writer.Bytes()) != "<div>me</div>" {
		t.Error("?")
	}

}

func Test_Teamplate2(t *testing.T) {
	writer := new(bytes.Buffer)
	data := struct{ Name string }{"x"}
	template.Must(template.New("?").Parse("{{.Name}}")).Execute(writer, data)
	if string(writer.Bytes()) != "x" {
		t.Error("not x")
	}
}

func Test_template3(t *testing.T) {
	wd := GetWorkingDir()
	path := filepath.Join(wd.Path, "testdata", "template.html")
	t.Logf("path: %v", path)
	tpt, e := template.New("A").ParseFiles(path)
	if (e == nil) {
		t.Error(e)
	}
	writer := new(bytes.Buffer)
	data := map[string]interface{}{"Name": "me"}

	e = tpt.Execute(writer, data)
	if e != nil {
		t.Error(e)
	}
	if string(writer.Bytes()) != "<div>me</div" {
		t.Error("?")
	}

}