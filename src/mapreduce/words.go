package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
//
// Split load optimally across processor cores.
func WordCount(text string) map[string]int {
	freqs := make(map[string]int)
	ch := make(chan map[string]int)

	words := strings.Fields(text)

	ChunkCount := 14
	ChunkSize := int(len(words) / ChunkCount)

	for i := 0; i < ChunkCount; i++ {
		// Last chunk reaches end words array to account for rounding
		if i == ChunkCount-1 {
			go SliceWordCount(words[i*ChunkSize:], ch)
		} else {
			go SliceWordCount(words[i*ChunkSize:(i+1)*ChunkSize], ch)
		}
	}

	// Merge chunks by adding values of every key together
	for i := 0; i < ChunkCount; i++ {
		for key, val := range <-ch {
			freqs[key] += val
		}
	}

	return freqs
}

func SliceWordCount(words []string, ch chan<- map[string]int) {
	chunk := make(map[string]int)

	// Increment counter of word for every occurence of word
	for _, word := range words {
		chunk[strings.ToLower(strings.Trim(word, ".,;:!?'\""))]++
	}

	ch <- chunk
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read in DataFile as a string called data
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", WordCount(string(data)))

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
