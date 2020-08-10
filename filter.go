package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Program started")
	minSize, maxSize, suffixes, files := handleCommandLine()
	channel1 := source(files)                          // Возвращаем канал с файлами
	channel2 := filterSuffixes(suffixes, channel1)     // Фильтруем по суффиксам
	channel3 := filterSize(minSize, maxSize, channel2) // Фильтруем по размеру
	sink(channel3)
	fmt.Println("Program finished")
}

func handleCommandLine() (int, int, [2]string, [5]string) {
	suffixes := [...]string{".txt", ".log"}
	files := [...]string{"test.txt", "smth.log", "lala.txt", "notValid.aa", "withoutExt"}
	return 1, 500, suffixes, files
}

func source(files [5]string) <-chan string {
	out := make(chan string, 1000)
	go func() {
		for _, filename := range files {
			out <- filename
		}
		close(out)
	}()
	return out
}

func filterSuffixes(suffixes [2]string, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		for filename := range in {
			if len(suffixes) == 0 {
				out <- filename
				continue
			}
			ext := strings.ToLower(filepath.Ext(filename))
			for _, suffix := range suffixes {
				if ext == suffix {
					out <- filename
					break
				}
			}
		}
		close(out)
	}()
	return out
}

func filterSize(maxSize int, minSize int, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		for filename := range in {
			if fileSize(filename) > maxSize && fileSize(filename) <= minSize {
				out <- filename
			}
		}
		close(out)
	}()
	return out
}

func fileSize(filename string) int {
	return 100
}

func sink(ch <-chan string) {
	for each := range ch {
		fmt.Printf("%s ", each)
	}
}
