package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BASEURL = "https://android.clients.google.com/fdfe/"
)

var country2Lang = map[string]string{
	"CN": "zh",
	"US": "en",
}

func GetAcceptLang(country string) string {
	if lang, ok := country2Lang[country]; ok {
		return lang + "_" + country
	}
	return "en_US"
}

func main() {
	pkg := "com.oyo.consumer"
	url := fmt.Sprintf(BASEURL+"details?doc=%s", pkg)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Accept-Language", GetAcceptLang("JP"))
	req.Header.Set("User-Agent", "Android-Finsky/7.7.31.O-all%20%5B0%5D%20%5BPR%5D%20153394623 (api=3,versionCode=80773100,sdk=24,device=hero2qltechn,hardware=samsungexynos8890,product=hero2qltezc)")
	req.Header.Set("X-DFE-Device-Id", "39d495e073c1540c")
	req.Header.Set("Authorization", "GoogleLogin auth=WwXTRX03-lnD-aeSdMZxK1sTbeQwmdMTbLH6NzFDzqMSejR_q0gW_s16BQvRxTy_l8y-Fg.")

	if err != nil {
		fmt.Println("[GooglePlay] build request error: ", err, ", pkg: ", pkg)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println("[GooglePlay] request error: ", err, ", pkg: ", pkg)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[Details] read resp body error: ", err, ", pkg: ", pkg)
	}
	fmt.Println(data)
}
