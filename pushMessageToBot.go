package main

import (
	"bytes"
	"errors"
	"go11x5/play"
	"log"
	"net/http"
	"net/url"
	"time"
)

func pushMsgToBot(msg string) error {
	c := http.Client{
		Timeout: 60 * time.Second,
	}
	content := url.Values{
		"chat_id": []string{"615491801"},
		"text": []string{msg},
	}
	// 开始加密
	encBytes := play.Encrypt(content.Encode())
	log.Println("Encrypted data is:", encBytes)
	// 将加密后的请求发送给自由节点
	req, err := http.NewRequest(http.MethodPost, "http://log.fenr.men:8000/tg?method=sendMessage", bytes.NewBuffer(encBytes))
	req.Header.Add("Authorization", "c371b934-b18c-4a44-a4a9-4d830fd0a527")
	if err!=nil {
		log.Println(err)
		return err
	}
	resp, err := c.Do(req)
	if err!=nil {
		log.Println(err)
		return err
	}
	if resp.Header.Get("Complete") == "true" {
		return nil
	}else {
		return errors.New("unknown error")
	}
}