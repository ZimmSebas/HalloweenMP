package files

import "fmt"

type FileSystem struct {
	files map[string]string
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		files: map[string]string{
			"test.md":             "Test file",
			"ReporteVatela-v2.md": "Reporte de Ines Vatela",
		},
	}
}

func (fs *FileSystem) ListFiles() (string, bool) {
	// See the things with permissions and list of files to see
	fmt.Println("Lista de archivos:")
	fmt.Println("")

	for file := range fs.files {
		fmt.Println("\t", file)
		fmt.Println("\t", fs.files[file])
		fmt.Println()
	}

	fmt.Println("")
	fmt.Println("Nota: La lista de archivos varia según el permiso de seguridad del usuario")
	fmt.Println("")

	return "", true
}

func (fs *FileSystem) ReadFile(name string) (string, bool) {
	result, ok := fs.files[name]

	if !ok {
		result = "Archivo no encontrado."
	} else {
		result = name
	}

	return result, ok
}
