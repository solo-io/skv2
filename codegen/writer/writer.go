package writer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/solo-io/skv2/codegen/render"
	"golang.org/x/tools/imports"
)

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

		// set default comment char to "#" as this is the most common
		commentChar := "#"
		switch filepath.Ext(name) {
		case ".go":
			commentChar = "//"
		}

		if w.Header != "" {
			content = fmt.Sprintf("%s %s\n\n", commentChar, w.Header) + content
		}

		if err := ioutil.WriteFile(name, []byte(content), perm); err != nil {
			return err
		}

		if !strings.HasSuffix(name, ".go") {
			continue
		}

		formatted, err := imports.Process(name, []byte(content), nil)
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(name, formatted, 0644); err != nil {
			return err
		}
	}
	return nil
}
