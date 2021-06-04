package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func PostJson(incompleteURL string, request interface{}, response interface{}) error {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(request); err != nil {
		return err
	}
	httpResp, err := http.Post(incompleteURL, "application/json; charset=utf-8", &buf)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return errors.New("http.Status:" + httpResp.Status)
	}
	return json.NewDecoder(httpResp.Body).Decode(response)
}

func PostFile(incompleteURL string, fileName string, response interface{}) error {
	if fileName == "" {
		return errors.New("上传文件不存在")
	}

	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(fileName)
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("filename", filepath.Base(fileName))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		return errFile1
	}
	err := writer.Close()
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, incompleteURL, payload)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	httpResp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return errors.New("http.Status:" + httpResp.Status)
	}
	return json.NewDecoder(httpResp.Body).Decode(response)
}

func GetRequest(u string, request url.Values, response interface{}) error {
	if !strings.HasSuffix(u, "?") {
		u += "?"
	}
	httpResp, err := http.Get(u + request.Encode())
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return errors.New("http.Status:" + httpResp.Status)
	}
	return json.NewDecoder(httpResp.Body).Decode(response)
}

func AuthTokenUrlValues(authToken string) url.Values {
	v := make(url.Values)
	v.Set("access_token", authToken)
	return v
}
