package main

import (
	"path/filepath"
)

func main() {
	minSize, maxSize, suffixes, files := handleCommandLine()

	channel1 := source(files)                          // Возвращаем канал с файлами
	channel2 := filterSuffixes(suffixes, channel1)     // Фильтруем по суффиксам
	channel3 := filterSize(minSize, maxSize, channel2) // Фильтруем по размеру
	sink(channel3)
}

func handleCommandLine() {
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
			ext := strings.toLower(filepath.Ext(filename))
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

func sinc(chan string) {

}
