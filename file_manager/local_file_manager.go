package file_manager

import (
	"io/ioutil"
	"os"
)

type LocalFileManager struct {
}

func (this *LocalFileManager) CrateFile(file string, content string) {
	fout, _ := os.Create(file)
	defer fout.Close()
	fout.WriteString(content)
}

func (this *LocalFileManager) DeleteFile(file string) {
	os.Remove(file)
}

func (this *LocalFileManager) UpdateFile(file string, content string) {
	if f, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm); err == nil {
		defer f.Close()
		f.WriteString(content)
	}
}

func (this *LocalFileManager) ReadFile(file string) string {
	f, _ := os.Open(file)
	fg, _ := ioutil.ReadAll(f)
	return string(fg)
}

func NewLocalFileManager() *LocalFileManager {
	return &LocalFileManager{}
}
