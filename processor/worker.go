package processor

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

func fileSearchWorker(fp *FileProcessor, searchToken []byte) []FileSearch {

	filesearch := []FileSearch{}

	var wg sync.WaitGroup
	resultMutex := &sync.Mutex{}
	maxWorkers := 10
	semaphore := make(chan struct{}, maxWorkers)

	for _, fileInf := range fp.files {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(fileInf fs.DirEntry) {
			defer wg.Done()
			defer func() { <-semaphore }()
			info, err := fileInf.Info()

			if err != nil || info.IsDir() {
				return
			}

			fmt.Println("File open and search -> " + info.Name())

			file, err := os.Open(filepath.Join(fp.TargetDirectory, fileInf.Name()))

			if err != nil {
				return
			}

			defer file.Close()

			foundData := 0

			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				lowerLine := bytes.ToLower(scanner.Bytes())
				count := bytes.Count(lowerLine, searchToken)
				foundData += count
			}

			fmt.Println("File close and search completed -> " + info.Name())

			fileSearch := FileSearch{
				FileName:   file.Name(),
				MatchCount: foundData,
			}

			// Sonuç slice'ını güvenli bir şekilde güncelleme
			resultMutex.Lock()
			filesearch = append(filesearch, fileSearch)
			resultMutex.Unlock()
		}(fileInf)
	}

	wg.Wait()
	return filesearch
}
