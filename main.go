package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
)

func check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", e)
		os.Exit(0)
	}
}

type Job struct {
  person map[string]string
  morph Morph
}

func main() {
	log.Println("Going to open file")
	file, err := os.Open(options.input_filename)
	check(err)
	defer file.Close()

  log.Println("File opened")
	morphedName := make(chan map[string]string)
	go getPeople(file, morphedName)

	var people []map[string]string
	for v := range morphedName {
		people = append(people, v)
	}
	log.Printf("Got %d People\n", len(people))

	workerCount :=runtime.NumCPU()
  if workerCount > 20 {
    workerCount = 20
  }

  jobChan := make(chan Job, workerCount*100)
	resultChan := make(chan string, workerCount*100)
  var wg sync.WaitGroup

	for range workerCount {
		wg.Add(1)
    go worker(jobChan, resultChan, &wg)
  }
  
  go func(){
    defer close(jobChan)
    for _, person := range people{
      for _, morph := range morphs{
          jobChan <- Job{person: person, morph: morph}
      }
    }
  }()
  
  go func(){
    wg.Wait()
    close(resultChan)
  }()

  saveOutput(resultChan)
}

func worker(jobChan <-chan Job,resultChan chan<- string, wg *sync.WaitGroup ) {
  defer wg.Done()

  n := &Name{}
  if len(jobChan) == 0 {
    fmt.Println("Empty Jobchan")
  }
  for job := range jobChan{
    n.Init(job.person)
    username := job.morph.Generate(n)
    if len(username) > 0 {
      resultChan <- username
    }
  }
}

