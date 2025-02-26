package main

import (
	"fmt"
	"os"
	"io"
	"path/filepath"
	"sync"
	"time"
	"bufio"
	"strings"	
)


func copyFile(src, dst string, wg *sync.WaitGroup, errChan chan error)  {
	
	defer wg.Done()
	
	// open source file
	srcFile, err := os.Open(src)

	if err != nil {
		errChan <- fmt.Errorf("error opening source file %s: %v", src, err)
		return 
	}

	defer srcFile.Close()

	// create destination file
	dstFile, err := os.Create(dst)

	if err != nil {
		errChan <- fmt.Errorf("error creating destination file %s: %v", dst, err)
		return 
	}

	defer dstFile.Close()

	// copy
	_, err = io.Copy(dstFile, srcFile)

	if err != nil {
		errChan <- fmt.Errorf("error copying file %s: %v", src, err)
		return 
	}

	// save dst file
	err = dstFile.Sync()

	if err != nil {
		errChan <- fmt.Errorf("error syncing file %s: %v", dst, err)
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
	errChan := make(chan error, len(files))

	// loop
	for _, file := range files {
		srcFilePath := filepath.Join(srcDir, file.Name())
		dstFilePath := filepath.Join(dstDir, file.Name())

		wg.Add(1)
		
		go copyFile(srcFilePath, dstFilePath, &wg, errChan)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(errChan) 

	// Print errors if any
	for err := range errChan {
		fmt.Println("Error:", err)
	}

	return nil
}



func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter source directory: ")
	srcPath, _ := reader.ReadString('\n')
	srcPath = strings.TrimSpace(srcPath)
	
	fmt.Print("Enter destination directory: ")	
	dstPath, _ := reader.ReadString('\n')
	dstPath = strings.TrimSpace(dstPath)

	fmt.Println("Source:", srcPath)
	fmt.Println("Destination:", dstPath)
	
	startTime := time.Now()
	err := copyDir(srcPath, dstPath)
		
	if err != nil {
		fmt.Println("Error copying directory:", err)
	} else {
		fmt.Println("Directory copied successfully!")
		
		endTime := time.Now()
		fmt.Println("All files copied in: ", endTime.Sub(startTime))
	}		
}