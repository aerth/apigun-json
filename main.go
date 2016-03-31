// The MIT License (MIT)
//
// Copyright (c) 2016 aerth
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

/*

apigun is a laser gun for testing your API server. Easy to modify.
This version is for ndjinn specifically. An you may want to use it only on localhost.

*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Status struct {
	IP        string `json:"ip"`
	UserAgent string `json:"user-agent"`
	Time      string `json:"time"`
	Response  int
	Error     error
}

const useragent = "apigun v1"

// Set User Agent
var tr = &http.Transport{
	DisableCompression: true,
}

// Set User Agent...
var apigun = &http.Client{
	CheckRedirect: redirectPolicyFunc,
	Transport:     tr,
}

// Set User Agent once mor--
func redirectPolicyFunc(req *http.Request, reqs []*http.Request) error {
	req.Header.Add("Content-Type", "[application/json; charset=utf-8")
	req.Header.Set("User-Agent", useragent)
	return nil
}

func pinger(c chan Status, url string) {
	for {
		var status Status

		req, err := http.NewRequest("GET", url, nil)
		// Here is the real one!
		req.Header.Set("User-Agent", useragent)

		if err != nil {
			log.Fatalln(err)
		}

		resp, err := apigun.Do(req)

		if err != nil {
			status.Error = err
			c <- status
			return
		}

		status.Response = resp.StatusCode
		if resp.StatusCode != http.StatusOK {
			log.Println("ERROR" + string(resp.StatusCode))

		}

		result := json.NewDecoder(resp.Body).Decode(&status)

		if result != nil {
			status.Error = result
		}

		c <- status

		// Options....
		//time.Sleep(time.Microsecond * 1000)
		//time.Sleep(time.Microsecond * 100)
		//time.Sleep(time.Microsecond * 1)

	}
}

func responsePrinter(c chan Status) {
	for i := 0; ; i++ {
		msg := <-c
		//fmt.Println(msg)
		if msg.Error != nil {
			fmt.Println(msg.Error)
			return
		}
		fmt.Println(msg.Response)
		fmt.Println(msg.IP)
		fmt.Println(msg.UserAgent)
		//timestamp := strings.Split(msg.Time, " ")
		fmt.Println("Date: " + msg.Time)
		fmt.Println(i)
	}
}
func main() {
	var c chan Status = make(chan Status)
	var d chan Status = make(chan Status)

	url := "http://localhost:8080/status"
	badurl := "http://localhost:8080/ap"

	// Get the values from API
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(c, url)
	go pinger(d, badurl)

	// Print the values
	go responsePrinter(c)
	go responsePrinter(d)

	// New line quit
	var input string
	fmt.Scanln(&input)
}
