package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

import "github.com/robfig/cron"
import _ "github.com/go-sql-driver/mysql"
import "github.com/joho/godotenv"

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

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	fmt.Println("Connected to the database")

	c := cron.New()
	c.AddFunc(os.Getenv("CRON_EXPRESSION"), func() { postDailyMessage(db) })
	c.Start()
}

func getNextTip(db *sql.DB) (string, string) {
	var command string
	var description string

	err := db.QueryRow(`
		SELECT command, description
		FROM tips
		WHERE posted = 0
		ORDER BY `+"`"+`index`+"`"+`
		LIMIT 1;
	`).Scan(&command, &description)

	if err != nil {
		log.Fatal("Failed to retrieve the next tip from database", err)
	}

	return command, description
}

func markTipAsPosted(db *sql.DB, command string) {
	_, err := db.Exec(`
		UPDATE tips
		SET posted = 1
		WHERE command = "` + command + `";
	`)

	if err != nil {
		log.Fatal("Failed to mark the tip as posted", err)
	}
}

func formatDailyMessage(command string, description string) string {
	return fmt.Sprintf("*%s*\n\n%s", command, description)
}

func postDailyMessage(db *sql.DB) {
	command, description := getNextTip(db)
	dailyMessage := formatDailyMessage(command, description)
	var jsonString = fmt.Sprintf(`{"text": "%s"}`, dailyMessage)
	var jsonBytes = []byte(jsonString)

	req, err := http.NewRequest("POST", os.Getenv("SLACK_WEBHOOK_URL"), bytes.NewBuffer(jsonBytes))
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

	markTipAsPosted(db, command)
	fmt.Println(string(body[:]))
}

func main() {
	loadEnvVariables()
	startDailyCron()
	blockForever()
}
