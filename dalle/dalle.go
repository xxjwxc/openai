package dalle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type DallE struct {
	apiKey  string
	userId  string
	timeOut time.Duration // 超时时间, 0表示不超时
}

type DallEReq struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
}

type DallEResp struct {
	Created int64       `json:"created"`
	Data    []DallEData `json:"data"`
	Error   DallError   `json:"error"`
}

type DallEData struct {
	Url     string `json:"url"`
	B64Json string `json:"b64_json"`
}

type DallError struct {
	Message string `json:"message"`
}

// NewDallE 新建一个智能绘图
func NewDallE(ApiKey, UserId string, timeOut time.Duration) *DallE {
	return &DallE{
		apiKey:  ApiKey,
		userId:  UserId,
		timeOut: timeOut,
	}
}

func (d *DallE) GenPhoto(prompt string, n int, size string) ([]string, error) {
	if len(size) == 0 {
		size = "512x512"
	}

	requestBody := DallEReq{
		Prompt:         prompt,
		N:              n,
		Size:           size,
		ResponseFormat: "url",
	}

	postData, _ := json.Marshal(requestBody)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", bytes.NewReader(postData))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", d.apiKey))
	resp, e := client.Do(req)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, e
	}
	// mylog.Debug(string(body))

	var dallEResp DallEResp
	err = json.Unmarshal(body, &dallEResp)
	if err != nil {
		return nil, err
	}

	if len(dallEResp.Error.Message) > 0 {
		return nil, fmt.Errorf("%v", dallEResp.Error.Message)
	}

	var out []string
	for _, v := range dallEResp.Data {
		out = append(out, v.Url)
	}
	return out, nil
}

func (d *DallE) GenPhotoBase64(prompt string, n int, size string) ([]string, error) {
	if len(size) == 0 {
		size = "512x512"
	}

	requestBody := DallEReq{
		Prompt:         prompt,
		N:              n,
		Size:           size,
		ResponseFormat: "b64_json",
	}

	postData, _ := json.Marshal(requestBody)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", bytes.NewReader(postData))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", d.apiKey))
	resp, e := client.Do(req)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, e
	}
	// mylog.Debug(string(body))

	var dallEResp DallEResp
	err = json.Unmarshal(body, &dallEResp)
	if err != nil {
		return nil, err
	}

	if len(dallEResp.Error.Message) > 0 {
		return nil, fmt.Errorf("%v", dallEResp.Error.Message)
	}

	var out []string
	for _, v := range dallEResp.Data {
		out = append(out, v.B64Json)
	}
	return out, nil
}
