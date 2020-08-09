package filter

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	minSize, maxSize, suffixes, files := handleCommandLine()

	channel1 := source(files)                          // Возвращаем канал с файлами
	channel2 := filterSuffixes(suffixes, channel1)     // Фильтруем по суффиксам
	channel3 := filterSize(minSize, maxSize, channel2) // Фильтруем по размеру
	sink(channel3)
}

func handleCommandLine() (int, int, []string, []string) {
	suffixes := make([]string, 5)
	files := make([]string, 5)
	return 1, 5, suffixes, files
}

func source(files []string) <-chan string {
	out := make(chan string, 1000)
	go func() {
		for _, filename := range files {
			out <- filename
		}
		close(out)
	}()
	return out
}

func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
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
			if fileSize(filename) > maxSize && fileSize(filename) < minSize {
				continue
			}
		}
	}()
	return out
}

func fileSize(filename string) int {
	return 1
}

func sink(ch <-chan string) {
	for each := range ch {
		fmt.Printf("%s", each)
	}
}
