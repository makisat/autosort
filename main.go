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

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("err loading .env:", err)
	}

	// add download path to .env file
	// add your own directories here
	download_directory := os.Getenv("DOWNLOAD_DIRECTORY")
	video_directory := download_directory + "/videos"
	image_directory := download_directory + "/images"
	document_directory := download_directory + "/documents"

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
					file := filepath.Base(event.Name)
					ext := path.Ext(event.Name)

					// add your own extensions here
					imageExtensions := []string{".jpeg", ".jpg", ".png"}
					videoExtensions := []string{".mp4", ".avi", ".mov"}
					documentExtensions := []string{".pdf", ".docx", ".html", ".txt"}

					// todo: create function to check if directory exists
					// todo: DRY!!!
					if _, err := os.Stat(video_directory); os.IsNotExist(err) {
						err := os.Mkdir(video_directory, os.ModePerm)
						if err != nil {
							fmt.Println("err at creating directory:", err)
						}
					}

					if _, err := os.Stat(image_directory); os.IsNotExist(err) {
						err := os.Mkdir(image_directory, os.ModePerm)
						if err != nil {
							fmt.Println("err at creating directory:", err)
						}
					}

					if _, err := os.Stat(document_directory); os.IsNotExist(err) {
						err := os.Mkdir(document_directory, os.ModePerm)
						if err != nil {
							fmt.Println("err at creating directory:", err)
						}
					}

					if contains(strings.ToLower(ext), imageExtensions){
						err := os.Rename(event.Name, fmt.Sprintf(image_directory+"\\%s", file))
						if err != nil {
							fmt.Println("err at moving file:", err)
						}
						fmt.Println("move success")
					}

					if contains(strings.ToLower(ext), videoExtensions){
						err := os.Rename(event.Name, fmt.Sprintf(video_directory+"\\%s", file))
						if err != nil {
							fmt.Println("err at moving file:", err)
						}
						fmt.Println("move success")
					}

					if contains(strings.ToLower(ext), documentExtensions){
						err := os.Rename(event.Name, fmt.Sprintf(document_directory+"\\%s", file))
						if err != nil {
							fmt.Println("err at moving file:", err)
						}
						fmt.Println("move success")
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

func contains(value string, list []string) bool {
	for _, e := range list {
		if e == value {
			return true
		}
	}
	return false
}
