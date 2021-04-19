package collector

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/solo-io/skv2/codegen/metrics"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/log"
	"github.com/solo-io/go-utils/stringutils"
)

type Collector interface {
	CollectImportsForFile(root, protoFile string) ([]string, error)
}

func NewCollector(customImports, commonImports []string) *collector {
	return &collector{
		customImports:    customImports,
		commonImports:    commonImports,
		importsExtractor: NewSynchronizedImportsExtractor(),
	}
}

type collector struct {
	customImports []string
	commonImports []string

	// The collector traverses a tree of files, opening and parsing each one.
	// These are costly operations that are often duplicated (ie fileA and fileB both import fileC)
	// The importsExtractor allows us to separate *how* to fetch imports from a file
	// from *when* to fetch imports from a file. This allows us to define a caching layer
	// in the importsExtractor and the collector doesn't have to be aware of it.
	importsExtractor ImportsExtractor
}

func (c *collector) CollectImportsForFile(root, protoFile string) ([]string, error) {
	return c.synchronizedImportsForProtoFile(root, protoFile, c.customImports)
}

var protoImportStatementRegex = regexp.MustCompile(`.*import "(.*)";.*`)

func (c *collector) detectImportsForFile(file string) ([]string, error) {
	metrics.IncrementFrequency("collector-opened-files")

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	var protoImports []string
	for _, line := range lines {
		importStatement := protoImportStatementRegex.FindStringSubmatch(line)
		if len(importStatement) == 0 {
			continue
		}
		if len(importStatement) != 2 {
			return nil, eris.Errorf("parsing import line error: from %v found %v", line, importStatement)
		}
		protoImports = append(protoImports, importStatement[1])
	}
	return protoImports, nil
}

func (c *collector) synchronizedImportsForProtoFile(absoluteRoot, protoFile string, customImports []string) ([]string, error) {
	// Define how we will extract the imports for the proto file
	importsFetcher := func(protoFileName string) ([]string, error) {
		return c.importsForProtoFile(absoluteRoot, protoFileName, customImports)
	}

	return c.importsExtractor.FetchImportsForFile(protoFile, importsFetcher)
}

func (c *collector) importsForProtoFile(absoluteRoot, protoFile string, customImports []string) ([]string, error) {
	importStatements, err := c.detectImportsForFile(protoFile)
	if err != nil {
		return nil, err
	}
	importsForProto := append([]string{}, c.commonImports...)
	for _, importedProto := range importStatements {
		importPath, err := c.findImportRelativeToRoot(absoluteRoot, importedProto, customImports, importsForProto)
		if err != nil {
			return nil, err
		}
		dependency := filepath.Join(importPath, importedProto)
		dependencyImports, err := c.synchronizedImportsForProtoFile(absoluteRoot, dependency, customImports)
		if err != nil {
			return nil, eris.Wrapf(err, "getting imports for dependency")
		}
		importsForProto = append(importsForProto, strings.TrimSuffix(importPath, "/"))
		importsForProto = append(importsForProto, dependencyImports...)
	}
	return stringutils.Unique(importsForProto), nil
}

func (c *collector) findImportRelativeToRoot(absoluteRoot, importedProtoFile string, customImports, existingImports []string) (string, error) {
	// if the file is already imported, point to that import
	for _, importPath := range existingImports {
		if _, err := os.Stat(filepath.Join(importPath, importedProtoFile)); err == nil {
			return importPath, nil
		}
	}
	rootsToTry := []string{absoluteRoot}

	for _, customImport := range customImports {
		absoluteCustomImport, err := filepath.Abs(customImport)
		if err != nil {
			return "", err
		}
		// Try the more specific custom imports first, rather than trying all of vendor
		rootsToTry = append([]string{absoluteCustomImport}, rootsToTry...)
	}

	// Sort by length, so longer (more specific paths are attempted first)
	sort.Slice(rootsToTry, func(i, j int) bool {
		elementsJ := strings.Split(rootsToTry[j], string(os.PathSeparator))
		elementsI := strings.Split(rootsToTry[i], string(os.PathSeparator))
		return len(elementsI) > len(elementsJ)
	})

	var possibleImportPaths []string
	for _, root := range rootsToTry {
		if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if strings.HasSuffix(path, importedProtoFile) {
				importPath := strings.TrimSuffix(path, importedProtoFile)
				possibleImportPaths = append(possibleImportPaths, importPath)

			}
			return nil
		}); err != nil {
			return "", err
		}
		// if found break
		if len(possibleImportPaths) > 0 {
			break
		}
	}
	if len(possibleImportPaths) == 0 {
		return "", eris.Errorf("found no possible import paths in root directory %v for import %v",
			absoluteRoot, importedProtoFile)
	}
	if len(possibleImportPaths) != 1 {
		log.Debugf("found more than one possible import path in root directory for "+
			"import %v: %v",
			importedProtoFile, possibleImportPaths)
	}
	return possibleImportPaths[0], nil

}

type ImportsFetcher func(file string) ([]string, error)

type ImportsExtractor interface {
	FetchImportsForFile(protoFile string, importsFetcher ImportsFetcher) ([]string, error)
}

// thread-safe
// The synchronizedImportsExtractor provides synchronized access to imports for a given proto file.
// It provides 2 useful features for tree traversal:
//	1. Imports for each file are cached, ensuring that if we attempt to access that file
//		during traversal again, we do not need to duplicate work.
//	2. If imports for a file are unknown, and simultaneous go routines attempt to load
//		the imports, only 1 will execute and the other will block, waiting for the result.
//		This reduces the number of times we open and parse files.
type synchronizedImportsExtractor struct {
	// cachedImports contains a map of fileImports, each indexed by their file name
	cachedImports sync.Map

	// protect access to activeRequests
	activeRequestsMu sync.RWMutex
	activeRequests   map[string]*importsFetchRequest
}

type fileImports struct {
	imports []string
	err     error
}

func NewSynchronizedImportsExtractor() ImportsExtractor {
	return &synchronizedImportsExtractor{
		activeRequests: map[string]*importsFetchRequest{},
	}
}

func (i *synchronizedImportsExtractor) FetchImportsForFile(protoFile string, importsFetcher ImportsFetcher) ([]string, error) {
	// Attempt to load the imports from the cache
	cachedImports, ok := i.cachedImports.Load(protoFile)
	if ok {
		typedCachedImports := cachedImports.(*fileImports)
		return typedCachedImports.imports, typedCachedImports.err
	}

	i.activeRequestsMu.Lock()

	// Ensure that we do not wait forever
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// If there's not a current active request for this file, create one.
	activeRequest := i.activeRequests[protoFile]
	if activeRequest == nil {
		activeRequest = newImportsFetchRequest()
		i.activeRequests[protoFile] = activeRequest

		// This goroutine has exclusive ownership over the current request.
		// It releases the resource by nil'ing the importRequest field
		// once the goroutine is done.
		go func(requestCtx context.Context) {
			// fetch the imports
			imports, err := importsFetcher(protoFile)

			// update the cache
			i.cachedImports.Store(protoFile, &fileImports{imports: imports, err: err})

			// Signal to waiting goroutines
			activeRequest.done(imports, err)

			// Free active request so a different request can run.
			i.activeRequestsMu.Lock()
			defer i.activeRequestsMu.Unlock()
			delete(i.activeRequests, protoFile)
		}(ctxWithTimeout)
	}

	inflightRequest := i.activeRequests[protoFile]
	i.activeRequestsMu.Unlock()

	select {
	case <-ctxWithTimeout.Done():
		// We should never reach this. This can only occur if we deadlock on file imports
		// which only happens with cyclic dependencies
		// Perhaps a safer alternative to erroring is just to execute the importsFetcher:
		// 	return importsFetcher(protoFile)
		return nil, ctxWithTimeout.Err()
	case <-inflightRequest.wait():
		// Wait for the existing request to complete, then return
		return inflightRequest.result()
	}
}

// importsFetchRequest is used to wait on some in-flight request from multiple goroutines.
// Derived from: https://github.com/coreos/go-oidc/blob/08563f61dbb316f8ef85b784d01da503f2480690/oidc/jwks.go#L53
type importsFetchRequest struct {
	doneCh  chan struct{}
	imports []string
	err     error
}

func newImportsFetchRequest() *importsFetchRequest {
	return &importsFetchRequest{doneCh: make(chan struct{})}
}

// wait returns a channel that multiple goroutines can receive on. Once it returns
// a value, the inflight request is done and result() can be inspected.
func (i *importsFetchRequest) wait() <-chan struct{} {
	return i.doneCh
}

// done can only be called by a single goroutine. It records the result of the
// inflight request and signals other goroutines that the result is safe to
// inspect.
func (i *importsFetchRequest) done(imports []string, err error) {
	i.imports = imports
	i.err = err
	close(i.doneCh)
}

// result cannot be called until the wait() channel has returned a value.
func (i *importsFetchRequest) result() ([]string, error) {
	return i.imports, i.err
}

// This is the old implementation, that does not include caching or locking
type primitiveImportsExtractor struct {
}

func (p *primitiveImportsExtractor) FetchImportsForFile(protoFile string, importsFetcher ImportsFetcher) ([]string, error) {
	return importsFetcher(protoFile)
}
