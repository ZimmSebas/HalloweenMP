package files

import "fmt"

type FileSystem struct {
	files map[string]string
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		files: map[string]string{
			"readme.txt":           "Test file",
			"ReporteVatela-v2.pdf": "Reporte de Ines Vatela",
		},
	}
}

func (fs *FileSystem) ListFiles() (string, bool) {
	// See the things with permissions and list of files to see
	fmt.Println("Lista de archivos:")
	for file := range fs.files {
		fmt.Println(file)
	}
	return "", true
}

func (fs *FileSystem) ReadFile(name string) (string, bool) {
	content, ok := fs.files[name]

	return content, ok
}
