package main

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	proxy    string
	user     string
	password string
	dest     string
)

func main() {

	// flag.StringVar(&proxy, "proxy", "", "provide proxy URL: IP:PORT")
	// flag.StringVar(&user, "user", "", "provide proxy user")
	// flag.StringVar(&password, "password", "", "provide proxy password")
	// flag.StringVar(&dest, "dest", "", "provide URL to access")
	// flag.Parse()

	proxy = "http://zproxy.lum-superproxy.io:22225"
	user = "lum-customer-yeahmobi-zone-sdk1-country-kr"
	password = "9hgzqlj9tq0a"
	dest = "http://tracking.yohoads.com/trace?offer_id=15946509&app_id=951&type=6f5a6d2500000028&aff_sub=yoh7526292_15946509_54589970_KR___b___1563523204___04FA6DC5-552C-44C7-8D81-1B7AF14241C9_Y29tLmludHNpZy5DYW1TY2FubmVyTGl0ZQ%3D%3D_iOS_1.408_ODg0MDQzNDYy&idfa=04FA6DC5-552C-44C7-8D81-1B7AF14241C9&sub_affiliate_id=54589970&aff_sub2=com.intsig.CamScannerLite"

	req, _ := http.NewRequest("GET", dest, nil)
	req.Header.Set("User-Agent", "")

	proxyURL, err := url.Parse(proxy)
	if err != nil {
		log.Println(err)
	}

	auth := fmt.Sprintf("%s:%s", user, password)
	basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Proxy-Authorization", basic)

	transport := &http.Transport{
		Proxy:           http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	transport.ProxyConnectHeader = req.Header
	client := &http.Client{Transport: transport}
	req.RequestURI = ""

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("erro: %s", err)
		return
	}
	fmt.Printf("code: %d", resp.StatusCode)
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(os.Stdout, string(htmlData))
}
