package main

import (
	"fmt"
	"io"
	"os"
)

func createFileError(err error) bool {
	if err != nil {
		fmt.Println("--Cannot make file--")
		return true
	}
	return err != nil
}

func createFile(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if createFileError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Printf("--Success create file in directory: %s\n", path)
}

func writeFileError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return err != nil
}

func writeFile(path string) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	writeFileError(err)
	defer file.Close()

	_, err = file.WriteString("Haloo\n")
	writeFileError(err)
	_, err = file.WriteString("Novalita Kusuma Wardhana\n")
	writeFileError(err)
	_, err = file.WriteString("Yogyakarta\n")
	writeFileError(err)
	_, err = file.WriteString("Indonesia\n")

	err = file.Sync()
	writeFileError(err)
	fmt.Printf("--Success edit file in: %s\n", path)
}

func readFileError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return err != nil
}

func readFile(path string) {
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if readFileError(err) {
		return
	}
	defer file.Close()

	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if readFileError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	fmt.Printf("Isi file %s:\n %s", path, string(text))
}

func removeFileError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return err != nil
}

func removeFile(path string) {
	var err = os.Remove(path)
	if removeFileError(err) {
		return
	}
	fmt.Printf("--Success delete file: %s\n", path)
}

func main() {
	fmt.Println("Part 50 File")
	fmt.Println("")

	fmt.Println("Part 50.1 Membuat File")
	pathFile := "/home/novalwardhana/novalAgung/text2.txt"
	createFile(pathFile)
	fmt.Println("")

	fmt.Println("Part 50.2 Mengedit isi file")
	writeFile(pathFile)
	fmt.Println("")

	fmt.Println("Part 50.3 Membaca file")
	readFile(pathFile)
	fmt.Println("")

	fmt.Println("Part 50.4 Menghapus file")
	removeFile(pathFile)
}
