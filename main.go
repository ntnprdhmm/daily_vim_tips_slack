package main

import "github.com/robfig/cron"
import "bytes"
import "fmt"
import "sync"
import "net/http"
import "io/ioutil"

var internalWebhookUrl string = "https://hooks.slack.com/services/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"

func blockForever() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func startDailyCron() {
	c := cron.New()
	c.AddFunc("@every 2s", postDailyMessage)
	c.Start()
}

func postDailyMessage() {
	var jsonBody = []byte(`{"text": "Hello from Go!"}`)
	req, err := http.NewRequest("POST", internalWebhookUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("An error happened while creating the request", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("An error happened while posting the daily message", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("An error happened while reading the daily message response body", err)
		return
	}

	fmt.Println(string(body[:]))
}

func main() {
	startDailyCron()
	blockForever()
}
