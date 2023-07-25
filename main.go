package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
)

type fileTypes struct {
	filetype string
	directory string
	extension []string
}

func (f fileTypes) contains(ext string) bool {
	for _, e := range f.extension {
		if e == strings.ToLower(ext) {
			return true
		}
	}
	return false
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("err loading .env:", err)
	}

	download_directory := os.Getenv("DOWNLOAD_DIRECTORY")

	directories := []fileTypes{
		{
			filetype: "images",
			directory: download_directory + "\\images",
			extension: []string{".jpeg", ".jpg", ".png", ".gif"},
		},
		{
			filetype: "videos",
			directory: download_directory + "\\videos",
			extension: []string{".mp4", ".avi", ".mov"},
		},
		{
			filetype: "sounds",
			directory: download_directory + "\\sounds",
			extension: []string{".mp3", ".wav", ".ogg"},
		},
		{
			filetype: "documents",
			directory: download_directory + "\\documents",
			extension: []string{".pdf", ".docx", ".html", ".txt"},
		},
		{
			filetype: "zip",
			directory: download_directory + "\\zip",
			extension: []string{".zip"},
		},
		{
			filetype: "exe",
			directory: download_directory + "\\exe",
			extension: []string{".exe"},
		},
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("err creating watcher:", err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Has(fsnotify.Create) {
					ext := path.Ext(event.Name)

					for _, directory := range directories {
						checkDirectory(directory.directory)
						if directory.contains(ext) {
							movefile(event.Name, directory.directory)
						}
					}
				}
			case err := <-watcher.Errors:
				fmt.Println("err at select:", err)
			}
		}
	}()

	err = watcher.Add(download_directory)
	if err != nil {
		fmt.Println("err at watcher.add:", err)
	}

	<-make(chan any)
}

func movefile(file string, directory string) {
	base := filepath.Base(file)
	err := os.Rename(file, directory + "\\" + base)
	if err != nil {
		fmt.Println("err at moving file:", err)
	}
}

func checkDirectory(directory string) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, os.ModePerm)
		if err != nil {
			fmt.Println("err at creating directory:", err)
		}
	}
}

