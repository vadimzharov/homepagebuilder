package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const sourcePathPrefix string = "source/"
const assetsPathPrefix string = "assets/"

func SourceFileExists(filename string) bool {

	filePath := sourcePathPrefix + filename

	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func WriteToSourceFile(data string, filename string) {

	log.Println("Writing content to file ", filename)

	filePath := sourcePathPrefix + filename

	f, err := os.Create(filePath)

	if err != nil {
		panic(fmt.Errorf("Cannot create file "+filename+" - %w", err))
	}

	defer f.Close()

	writtendata, err := f.WriteString(data)

	if err != nil {
		panic(fmt.Errorf("Cannot write to "+filename+" - %w", err))
	}

	log.Printf("Wrote %d bytes to "+filename+"\n", writtendata)

	f.Close()

}

func ReadFromSourceFile(filename string) []byte {

	log.Println("Reading content from file", filename)

	filePath := sourcePathPrefix + filename

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)

	}

	return bytes

}

func CopySourceToAssets(srcfilename string, destfilename string) {

	srcFilePath := sourcePathPrefix + srcfilename
	destFilePath := assetsPathPrefix + destfilename

	// Open the source file for reading
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	// Create the destination file for writing
	dstFile, err := os.Create(destFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}

}

func EraseAllSourceFiles() {

	files, err := filepath.Glob(filepath.Join(sourcePathPrefix, "*"))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		err = os.Remove(file)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func EraseSourceFile(filename string) {

	filePath := sourcePathPrefix + filename

	err := os.Remove(filePath)
	if err != nil {
		log.Fatal(err)
	}

}
