package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"net/http"
)

type Gcm struct {
	Client    *http.Client
	Api_key   string
	Regist_id string
}

func NewGcm() *Gcm {

	var config Config
	toml.DecodeFile("config.toml", &config)

	gcm := &Gcm{
		Client:    new(http.Client),
		Api_key:   config.Setting.Api_key,
		Regist_id: config.Setting.Regist_id,
	}

	return gcm
}

type Config struct {
	Setting Setting
}

type Setting struct {
	Api_key   string
	Regist_id string
}

type Message struct {
	Registration_ids []string          `json:"registration_ids"`
	Data             map[string]string `json:"data"`
}

type Response struct {
	Body string
}

func (gcm *Gcm) SendMessage(message string, key string, value string) (*Response, error) {

	fmt.Println(gcm.Api_key)
	fmt.Println(gcm.Regist_id)

	param := map[string]string{
		"message": message,
		"params":  fmt.Sprintf("%s:%s", key, value),
	}

	msg := &Message{
		Registration_ids: []string{gcm.Regist_id},
		Data:             param,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("errr")
	}

	req, err := http.NewRequest("POST", GCM_SERVER, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Authorization", fmt.Sprintf("key=%s", gcm.Api_key))

	res, err := gcm.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Response{
		Body: string(body),
	}

	return response, nil
}

func main() {

	msg := flag.String("msg", "default message", "GCMで送信するメッセージ")
	key := flag.String("key", "", "送信するパラメータのkey")
	value := flag.String("value", " ", "送信するパラメータのvalue")
	flag.Parse()

	gcm := NewGcm()

	res, err := gcm.SendMessage(*msg, *key, *value)
	if err != nil {
		fmt.Println("send err.")
	} else {
		fmt.Println(res.Body)
	}
}
