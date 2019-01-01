package main

import (
	"os"
	"bytes"
	"log"
	"fmt"
	"sync"
	"net/http"
	"io/ioutil"
	"database/sql"
)

import "github.com/robfig/cron"
import _ "github.com/go-sql-driver/mysql"
import "github.com/joho/godotenv"

var internalWebhookUrl string = "https://hooks.slack.com/services/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"

func blockForever() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func loadEnvVariables() {
	err := godotenv.Load()

  if err != nil {
    log.Fatal("Error loading .env file", err)
  }
}

func getDataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}

func startDailyCron() {
	dsn := getDataSourceName()
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	c := cron.New()
	c.AddFunc("@every 2s", func() { postDailyMessage(db) })
	c.Start()
}

func postDailyMessage(db *sql.DB) {
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
	loadEnvVariables()
	startDailyCron()
	blockForever()
}
