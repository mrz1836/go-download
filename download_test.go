package main

import (
	"testing"
)

//TestDefaultDirectory test the default directory being present
func TestDefaultDirectory(t *testing.T) {
	err := createDefaultDirectory()
	if err != nil {
		t.Fatal("unable to create default directory", err.Error())
	}
}

//TestParseFromFile test parsing from a file
func TestParseFromFile(t *testing.T) {

	//Test bad file
	_, err := parseFromFile("bad-file.json")
	if err == nil {
		t.Fatal("error should have occurred, file not found")
	} else if err.Error() != "error opening config file: open bad-file.json: no such file or directory" {
		t.Fatal(err.Error())
	}

	//Test parsing a valid file
	files, err := parseFromFile("example.json")
	if err != nil {
		t.Fatal("error should have  not occurred", err.Error())
	} else if len(files) == 0 {
		t.Fatal("no files found, there should be at least 1")
	}

	//Test the data
	if files[0].FileName != "wiki-logo" {
		t.Fatal("file name expected was not found", files[0].FileName)
	}

	if files[0].FolderPath != "Test-Folder-Path" {
		t.Fatal("folder path expected was not found", files[0].FolderPath)
	}

	if files[0].FileURL != "https://en.wikipedia.org/wiki/File:Wikipedia-logo-v2.svg" {
		t.Fatal("file url expected was not found", files[0].FileURL)
	}

}

//TestDownloadAllFiles test parsing from a file and downloading
func TestDownloadAllFiles(t *testing.T) {

	//Test parsing a valid file
	files, err := parseFromFile("example.json")
	if err != nil {
		t.Fatal("error should have  not occurred", err.Error())
	} else if len(files) == 0 {
		t.Fatal("no files found, there should be at least 1")
	}

	//Download all the files
	err = downloadAllFiles(files)
	if err != nil {
		t.Fatal("failed in downloading all files: ", err.Error())
	}
}

//BenchmarkCreateDefaultDirectory benchmarks the createDefaultDirectory method
func BenchmarkCreateDefaultDirectory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = createDefaultDirectory()
	}
}

//BenchmarkParseFromFile benchmarks the parseFromFile method
func BenchmarkParseFromFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = parseFromFile("example.json")
	}
}
