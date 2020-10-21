package writer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/solo-io/skv2/codegen/render"
)

var commentPrefixes = map[string]string{
	".go":    "//",
	".proto": "//",
	".js":    "//",
	".ts":    "//",
}

type FileWriter interface {
	WriteFiles(files []render.OutFile) error
}

// writes to the filesystem
type DefaultFileWriter struct {
	Root   string
	Header string // prepended to files before write
}

func (w *DefaultFileWriter) WriteFiles(files []render.OutFile) error {
	for _, file := range files {
		name := filepath.Join(w.Root, file.Path)
		content := file.Content

		if err := os.MkdirAll(filepath.Dir(name), 0777); err != nil {
			return err
		}

		perm := file.Permission
		if perm == 0 {
			perm = 0644
		}

		log.Printf("Writing %v", name)

		commentPrefix := commentPrefixes[filepath.Ext(name)]
		if commentPrefix == "" {
			// set default comment char to "#" as this is the most common
			commentPrefix = "#"
		}

		if w.Header != "" {
			content = fmt.Sprintf("%s %s\n\n", commentPrefix, w.Header) + content
		}

		if err := ioutil.WriteFile(name, []byte(content), perm); err != nil {
			return err
		}

	}
	return nil
}
