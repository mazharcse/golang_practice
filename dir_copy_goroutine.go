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