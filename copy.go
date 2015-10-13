package main

import (
	"fmt"
	"io"
	"os"
)

// CopyFile copy file
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err == nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		} else {
			return err
		}
	} else {
		return err
	}
	return
}

// CopyDir copy directory recursively
func CopyDir(source string, dest string) (err error) {
	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)
	for _, obj := range objects {
		sourcefilePointer := source + "/" + obj.Name()
		destinationfilePointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories recursively
			err = CopyDir(sourcefilePointer, destinationfilePointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform Copy
			err = CopyFile(sourcefilePointer, destinationfilePointer)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return
}
