package file_manager

type FileManager interface {
	CrateFile(file string, content string)
	DeleteFile(file string)
	UpdateFile(file string, content string)
	ReadFile(file string) string
}
