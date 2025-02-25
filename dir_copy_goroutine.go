package main

import (
	"fmt"
	"os"
	"io"
	"path/filepath"
	"sync"
	"time"
)


func copyFile(src, dst string, wg *sync.WaitGroup)  {
	
	defer wg.Done()
	
	// open source file
	srcFile, err := os.Open(src)

	if err != nil {
		return 
	}

	defer srcFile.Close()

	// create destination file
	dstFile, err := os.Create(dst)

	if err != nil {
		return 
	}

	defer dstFile.Close()

	// copy
	_, err = io.Copy(dstFile, srcFile)

	if err != nil {
		return 
	}

	// save dst file
	err = dstFile.Sync()

	if err != nil {
		return 
	}
	
	fmt.Println("Copied: ", src, "->", dst)
}



func copyDir(srcDir, dstDir string) error {

	// read file dir
	files, err := os.ReadDir(srcDir)

	if err != nil {
		return err
	}

	err = os.MkdirAll(dstDir, os.ModePerm)

	if err != nil {
		return err
	}
	
	var wg sync.WaitGroup

	// loop
	for _, file := range files {
		srcFilePath := filepath.Join(srcDir, file.Name())
		dstFilePath := filepath.Join(dstDir, file.Name())

		wg.Add(1)
		
		go copyFile(srcFilePath, dstFilePath, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	return nil
}



func main() {
	srcPath := "/home/mazhar/books/python/django_course/The Ultimate Django Series Part 1/5.Django ORM (100m)"
	dstPath := "./backup"

	startTime := time.Now()
	err := copyDir(srcPath, dstPath)
		
	
	if err != nil {
		fmt.Println("Error copying directory:", err)
	} else {
		fmt.Println("Directory copied successfully!")
	}

	endTime := time.Now()
	fmt.Println("All files copied in: ", endTime.Sub(startTime))	
}