package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	args := os.Args
	webURL := args[1]
	cssURL := args[2]
	appURL := args[3]
	log.SetOutput(os.Stderr)
	ret, err := purge(webURL, cssURL, appURL)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(ret)
}

func purge(webURL, cssURL, appURL string) (string, error) {
	log.Printf("web: %s\n", webURL)
	log.Printf("css: %s\n", cssURL)
	log.Printf("app: %s\n", appURL)

	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)

	{
		fw, err := mw.CreateFormFile("html", "web.html")
		if err != nil {
			return "", err
		}

		resp, err := http.Get(webURL)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		_, err = io.Copy(fw, resp.Body)
		if err != nil {
			return "", err
		}
	}

	{
		fw, err := mw.CreateFormFile("css", "web.css")
		if err != nil {
			return "", err
		}

		resp, err := http.Get(cssURL)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		_, err = io.Copy(fw, resp.Body)
		if err != nil {
			return "", err
		}
	}

	if err := mw.Close(); err != nil {
		return "", err
	}

	resp, err := http.Post(appURL, mw.FormDataContentType(), body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
