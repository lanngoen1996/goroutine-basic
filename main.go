package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Response struct {
	Status    int
	Url       string
	Time      float64
	StartTime string
}

func HttpRequest(url string, ch chan Response) {
	t := time.Now()
	resp, err := http.Get(url)
	secs := time.Since(t).Seconds()
	r := Response{
		Url:       url,
		Time:      secs,
		StartTime: t.Format("2006-01-02 15:04:05"),
	}
	if err != nil {
		r.Status = 500
		ch <- r
		return
	}
	r.Status = resp.StatusCode
	ch <- r
}

func HandleRequest(ch chan Response, wg *sync.WaitGroup) {
	defer wg.Done()
	r := <-ch
	m := fmt.Sprintf("%s | URL %s => status: %d, time: %f", r.StartTime, r.Url, r.Status, r.Time)
	fmt.Println(m)
}

func main() {
	urls := []string{
		"http://www.facebook.com/",
		"https://line.me/th/",
		"https://line.me/en/",
		"http://www.golang.org/",
		"https://line.me/jj/",
		"https://www.cloudflare.com/",
		"http://www.google.com/",
		"http://localhost:8000",
		"https://medium.com/",
	}
	chanel := make(chan Response)

	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go HttpRequest(url, chanel)
		go HandleRequest(chanel, &wg)
	}
	wg.Wait()

	fmt.Println("Done all process.")
}
