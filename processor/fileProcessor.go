package processor

import (
	"os"
	"strings"
)

type FileProcessor struct {
	TargetDirectory string
	TargetFile      string
	files           []os.DirEntry
	error
}

func NewFileProcessor(targetDirectory string, targetFile string) *FileProcessor {
	processor := FileProcessor{
		TargetDirectory: targetDirectory,
		TargetFile:      targetFile,
	}

	processor.listFiles()

	return &processor
}

func (fp *FileProcessor) listFiles() {
	fp.files, fp.error = os.ReadDir(fp.TargetDirectory)
}

func (fp *FileProcessor) ListFiles() Result {
	result := Result{
		TargetDirectory: fp.TargetDirectory,
		Error:           fp.error,
	}

	if fp.error != nil {
		return result
	}

	result.FileSearch = make([]FileSearch, len(fp.files))

	for _, file := range fp.files {

		fileSearch := FileSearch{
			FileName: file.Name(),
		}
		result.FileSearch = append(result.FileSearch, fileSearch)
	}

	return result
}

func (fp *FileProcessor) ProcessFile(searhingWord string) Result {
	result := Result{
		TargetDirectory: fp.TargetDirectory,
		Error:           fp.error,
	}

	if result.Error != nil {
		return result
	}

	token := []byte(strings.ToLower(searhingWord))

	result.FileSearch = append(result.FileSearch, fileSearchWorker(fp, token)...)

	return result
}
