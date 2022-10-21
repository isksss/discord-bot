package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

func main() {

	f, err := os.Open("config.yml")
	if err != nil {
		return
	}
	defer f.Close()

	d := yaml.NewDecoder(f)

	var m map[string]string

	if err := d.Decode(&m); err != nil {
		log.Fatal(err)
	}

	var whurl string
	var username string
	var message string

	whurl = m["url"]
	username = m["name"]
	message = m["message"]

	const format = "2006-01-02 15:04:05"
	day := time.Now().Add(9 * time.Hour)
	message += day.Format(format)

	dw := &discordWebhook{UserName: username, Content: message}

	sendWebhook(whurl, dw)
}

func sendWebhook(whurl string, dw *discordWebhook) {
	j, err := json.Marshal(dw)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}

	req, err := http.NewRequest("POST", whurl, bytes.NewBuffer(j))
	if err != nil {
		fmt.Println("new request err:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client err:", err)
		return
	}
	if resp.StatusCode == 204 {
		fmt.Println("sent", dw)
	} else {
		fmt.Printf("%#v\n", resp)
	}
}

type discordWebhook struct {
	UserName string `json:"username"`
	Content  string `json:"content"`
	TTS      bool   `json:"tts"`
}
