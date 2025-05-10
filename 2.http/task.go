package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/trace"
	"strconv"
	"sync"
)

type comment struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func HttpWorkers() []comment {

	var wgWorker sync.WaitGroup
	var wgResult sync.WaitGroup

	const jobsCount int = 500

	resSl := make([]comment, 0, jobsCount)

	jobs := make(chan int)
	results := make(chan comment, 100)

	wgResult.Add(1)
	go func() {
		defer wgResult.Done()
		for { //вместо него select
			result, ok := <-results
			if !ok {
				return
			}
			resSl = append(resSl, result)
		}

	}()

	for j := 0; j < 40; j++ { // ErrorGroup //Залить на гитхаб задачи
		wgWorker.Add(1)
		go func() {
			defer wgWorker.Done()
			for {
				num, ok := <-jobs
				if !ok {
					return
				}
				resp, err := http.Get("https://jsonplaceholder.typicode.com/comments/" + strconv.Itoa(num))
				if err != nil {
					log.Printf("Ошибка при попытке отправить GET-запрос: %v", err)
					continue
				}

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Printf("Ошибка при чтении ответа: %v", err)
					continue
				}
				resp.Body.Close()

				var c comment
				err = json.Unmarshal(body, &c)
				if err != nil {
					log.Printf("Ошибка парсинга: %v", err)
					continue
				}
				results <- c
			}

		}()
	}

	for i := 1; i <= jobsCount; i++ {
		jobs <- i
	}
	close(jobs)

	wgWorker.Wait()
	close(results)

	wgResult.Wait()
	return resSl
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := trace.Start(f); err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	resSl := HttpWorkers()
	fmt.Println(resSl)
}
