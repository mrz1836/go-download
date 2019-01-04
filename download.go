/*
Package main (go-download) takes a JSON configuration of urls, paths and file names and downloads all the files locally
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

//Default package constants
const (
	defaultDownloadPath = "downloads"
)

//file is the default attributes for a parsed file object
type file struct {
	FileName   string `json:"file_name"`
	FileURL    string `json:"file_url"`
	FolderPath string `json:"folder_path"`
}

//main to run the application
func main() {

	//Create the default folder to store the downloaded files
	err := createDefaultDirectory()
	if err != nil {
		log.Panicf("failed creating default directory: %s", err.Error())
	}

	//Load the JSON file (or change this to point to a URL)
	files, err := parseFromFile("example.json")
	if err != nil {
		log.Panicf("error parsing JSON file: %s", err.Error())
	}

	//Download all the files
	err = downloadAllFiles(files)
	if err != nil {
		log.Panicf("failed in downloading all files: %s", err.Error())
	}

}

//createDefaultDirectory create the default directory
func createDefaultDirectory() (err error) {
	if _, err := os.Stat(defaultDownloadPath); os.IsNotExist(err) {
		if err = os.MkdirAll(defaultDownloadPath, os.ModePerm); err != nil {
			err = fmt.Errorf("error creating default directory %s error: %s", defaultDownloadPath, err.Error())
		}
	}

	return
}

//parseFromFile will parse the JSON file into file struct
func parseFromFile(filePath string) (files []file, err error) {

	//Load the JSON file (or change this to point to a URL)
	var configFile *os.File
	if configFile, err = os.Open(filePath); err != nil {
		err = fmt.Errorf("error opening config file: %s", err.Error())
		return
	}

	//Decode the JSON from the file
	parser := json.NewDecoder(configFile)
	if err = parser.Decode(&files); err != nil {
		err = fmt.Errorf("error parsing config file: %s", err.Error())
	}

	return
}

//downloadAllFiles downloads all given files
func downloadAllFiles(files []file) (err error) {

	//No files found?
	if len(files) == 0 {
		err = fmt.Errorf("no url/files found in JSON file to download")
		return
	}

	//Loop and download //todo: make this multi-threaded using Channels and Go Routine
	for _, f := range files {
		err = downloadFile(f.FileURL, f.FolderPath, f.FileName)
		if err != nil {
			err = fmt.Errorf("error downloading file: %s error: %s", f.FileURL, err.Error())
			return
		}
	}

	//Get the total number / list of files downloaded in the default folder
	var downloadCount int
	err = filepath.Walk(defaultDownloadPath, func(path string, f os.FileInfo, err error) error {

		//Are we a file vs directory
		if !f.IsDir() {
			downloadCount++
		}
		return nil
	})

	//Get the file count in the downloads folder
	log.Printf("Total URLs in JSON file: %d \n", len(files))
	log.Printf("Files downloaded and found locally: %d \n", downloadCount)

	return
}

//downloadFile will take a url and download the file locally if not found
func downloadFile(fromURL, folderPath, fileName string) (err error) {

	//Get the extension from the file being downloaded
	var extension = filepath.Ext(fromURL)

	//Locate the folder path into the global downloads folder (global-path/local-path)
	folderPath = fmt.Sprintf("%s%c%s", defaultDownloadPath, os.PathSeparator, folderPath)

	//Create the final path + file name + extension
	finalFilePath := fmt.Sprintf("%s%c%s%s", folderPath, os.PathSeparator, fileName, extension)

	//Logs for tracking the progress
	log.Printf("Downloading file from url: %s to local path: %s%c%s%s", fromURL, folderPath, os.PathSeparator, fileName, extension)

	//Directory exists? (if not, create the directory first)
	if _, err = os.Stat(folderPath); os.IsNotExist(err) {
		if err = os.MkdirAll(folderPath, os.ModePerm); err != nil {
			log.Printf("error creating directory %s error: %s", folderPath, err.Error())
			return
		}
	}

	//File already exists in this directory?
	if _, err = os.Stat(finalFilePath); !os.IsNotExist(err) {
		log.Printf("file %s already exists, skipped downloading", fileName)
		return
	}

	//Create the file
	var output *os.File
	if output, err = os.Create(finalFilePath); err != nil {
		log.Printf("error while creating file %s error: %s", fileName, err.Error())
		return
	}

	//Defer the closing of the output
	defer func() {
		if err = output.Close(); err != nil {
			log.Printf("error occurred in closing output: %s", err.Error())
		}
	}()

	//Get the file
	var response *http.Response
	if response, err = http.Get(fromURL); err != nil {
		log.Printf("error while downloading from url: %s error: %s", fromURL, err.Error())
		return
	}

	//Defer the closing of the response body
	defer func() {
		if err = response.Body.Close(); err != nil {
			log.Printf("error occurred in closing the response body: %s", err.Error())
		}
	}()

	//Copy the contents
	var numberOfBytes int64
	if numberOfBytes, err = io.Copy(output, response.Body); err != nil {
		log.Printf("error while downloading from url: %s error: %s", fromURL, err.Error())
		return
	}

	//Success
	log.Printf("File downloaded successfully and %d bytes downloaded", numberOfBytes)
	return
}
