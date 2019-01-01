package main

import "github.com/robfig/cron"
import "fmt"
import "sync"

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
	fmt.Println("Hello from Cron")
}

func main() {
	startDailyCron()
	blockForever()
}
