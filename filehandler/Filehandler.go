package filehandler

import (
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
)

type FileHandler struct {
	FilePath  string
	FileEvent chan map[string]string // map[filename]operation
}

func New(filePath string) *FileHandler {
	// if filepath is not a directory create it
	if err := os.Mkdir(filePath, 0777); err != nil {
		if os.IsExist(err) {
			// directory already exists
		} else {
			panic(err)
		}
	}

	fh := FileHandler{
		FilePath:  filePath,
		FileEvent: make(chan map[string]string),
	}
	return &fh
}

func (fh *FileHandler) Monitor() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()
	err = watcher.Add(fh.FilePath)
	if err != nil {
		panic(err)
	}
	func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fh.FileEvent <- map[string]string{
					filepath.Base(event.Name): event.Op.String(),
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				panic(err)
			}
		}
	}()

}
