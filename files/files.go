package files

type FileSystem struct {
	files map[string]string
}

func NewFileSystem() *FileSystem {
	return &FileSystem{}
}

func (fs *FileSystem) ListFiles() string {
	// See the things with permissions
	return "to-do"
}

func (fs *FileSystem) ReadFile(name string) (string, bool) {
	content, ok := fs.files[name]
	return content, ok
}
