package render

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

func PrependMockgenDirective(files []OutFile) {
	// prepend each .go file with a go:generate directive to run mockgen on the file
	for i, file := range files {
		if !strings.HasSuffix(file.Path, ".go") {
			continue
		}
		// only add the directive if the file contains interfaces
		if !containsInterface(file.Content) {
			continue
		}

		baseFile := filepath.Base(file.Path)
		mockgenComment := fmt.Sprintf("//go:generate mockgen -source ./%v -destination mocks/%v", baseFile, baseFile)

		file.Content = fmt.Sprintf("%v\n\n%v", mockgenComment, file.Content)
		files[i] = file
	}
}

var interfaceRegex = regexp.MustCompile("type .* interface")

func containsInterface(content string) bool {
	return interfaceRegex.MatchString(content)
}
