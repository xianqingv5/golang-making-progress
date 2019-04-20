package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type CommonNode struct {
	Url       string   `json:"url,omitempty"`
	Lang      string   `json:"language,omitempty"`
	Countries []string `json:"countries,omitempty"`
}

type Icon map[string][]CommonNode
type ScreenShot map[string][]CommonNode
type Creative map[string][]CommonNode
type Video map[string][]CommonNode

type GooglePlay struct {
	Lang        string     `json:"language,omitempty"`
	AppCat      string     `json:"app_category_primary,omitempty"`
	AppCatSec   []string   `json:"app_category_secondary,omitempty"`
	AppDesc     string     `json:"app_description,omitempty"`
	CurVer      string     `json:"current_version,omitempty"`
	Free        bool       `json:"free,omitempty"`
	Iap         bool       `json:"iap,omitempty"`
	Price       float32    `json:"price,omitempty"`
	RateAvg     float32    `json:"rating_average,omitempty"`
	RateCurVer  float32    `json:"rating_current_version,omitempty"`
	RateCntAvg  int        `json:"rating_counts_average,omitempty"`
	ContRate    string     `json:"content_rating,omitempty"`
	FileSz      int        `json:"file_size_bytes,omitempty"`
	Created     string     `json:"created,omitempty"`
	Updated     string     `json:"updated,omitempty"`
	Downloads   string     `json:"downloads,omitempty"`
	WhatIsNew   string     `json:"what_is_new,omitempty"`
	Permissions string     `json:"permissions,omitempty"`
	MinOsVer    string     `json:"min_os_version,omitempty"`
	SuppDevs    []string   `json:"supported_devices,omitempty"`
	SuppLangs   []string   `json:"supported_languages,omitempty"`
	ScreenShots ScreenShot `json:"screenshots,omitempty"`
	Icons       Icon       `json:"icons,omitempty"`
	Videos      Video      `json:"videos,omitempty"`
}

type Itune struct {
	Lang          string     `json:"language,omitempty"`
	AppCat        string     `json:"app_category_primary,omitempty"`
	AppCatSec     []string   `json:"app_category_secondary,omitempty"`
	AppDesc       string     `json:"app_description,omitempty"`
	CurVer        string     `json:"current_version,omitempty"`
	Free          bool       `json:"free,omitempty"`
	Iap           bool       `json:"iap,omitempty"`
	Price         float32    `json:"price,omitempty"`
	Currency      string     `json:"currency,omitempty"`
	Advisories    string     `json:"advisories,omitempty"`
	RateAvg       float32    `json:"rating_average,omitempty"`
	RateCurVer    float32    `json:"rating_current_version,omitempty"`
	RateCntAvg    int        `json:"rating_counts_average,omitempty"`
	RateCntCurVer int        `json:"rating_counts_current_version,omitempty"`
	ContRate      string     `json:"content_rating,omitempty"`
	FileSz        int        `json:"file_size_bytes,omitempty"`
	Created       string     `json:"created,omitempty"`
	Updated       string     `json:"updated,omitempty"`
	Downloads     string     `json:"downloads,omitempty"`
	WhatIsNew     string     `json:"what_is_new,omitempty"`
	MinOsVer      string     `json:"min_os_version,omitempty"`
	SuppDevs      []string   `json:"supported_devices,omitempty"`
	SuppLangs     []string   `json:"supported_languages,omitempty"`
	ScreenShots   ScreenShot `json:"screenshots,omitempty"`
	Icons         Icon       `json:"icons,omitempty"`
	Videos        Video      `json:"videos,omitempty"`
}

type Target struct {
	Platforms []string `json:"platforms,omitempty"`
	Devices   []string `json:"devices,omitempty"`
	DeviceOss []string `json:"device_os,omitempty"`
	Countries []string `json:"countries,omitempty"`
	Carriers  []string `json:"carriers,omitempty"`
}

type Traffic struct {
	Allowed   []string `json:"allowed,omitempty"`
	Forbidden []string `json:"forbidden,omitempty"`
}

type Financial struct {
	Payout         float32 `json:"payout,omitempty"`
	Currency       string  `json:"currency,omitempty"`
	CapShared      bool    `json:"cap_shared,omitempty"`
	CapDaily       int     `json:"cap_daily,omitempty"`
	CapMonthly     int     `json:"cap_monthly,omitempty"`
	RemainCapDaily int     `json:"remaining_cap_daily,omitempty"`
}

type Offer struct {
	Id           int       `json:"offer_id,omitempty"`
	Name         string    `json:"offer_name,omitempty"`
	Status       string    `json:"status,omitempty"` // 这个字段文档上面是缺失的
	Cat          []string  `json:"offer_category,omitempty"`
	Desc         string    `json:"offer_description,omitempty"`
	Incent       bool      `json:"is_incent,omitempty"`
	TrackLink    string    `json:"tracking_link,omitempty"`
	IsS2s        bool      `json:"is_s2s"`
	S2sTrackLink string    `json:"s2s_tracking_link"`
	ConvFlow     []string  `json:"conversion_flow,omitempty"`
	Targeting    Target    `json:"targeting,omitempty"`
	Traffic      Traffic   `json:"traffic,omitempty"`
	Financials   Financial `json:"financials,omitempty"`
}

type Product struct {
	AppId        int    `json:"app_id,omitempty"`
	AppName      string `json:"app_name,omitempty"`
	PkgName      string `json:"pkg_name,omitempty"`
	AppShortDesc string `json:"app_short_description,omitempty"`
	AppUrl       string `json:"app_url,omitempty"`
	ProCat       string `json:"product_category,omitempty"`
	ProCatSec    string `json:"product_category_secondary,omitempty"`

	GooglePlays []GooglePlay `json:"google_play,omitempty"`
	Itunes      []Itune      `json:"itune,omitempty"`
	Icons       Icon         `json:"icons,omitempty"`
	Creatives   Creative     `json:"creatives,omitempty"`
	Videos      Video        `json:"videos,omitempty"`
	Offers      []Offer      `json:"offers,omitempty"`
}

func parseYeahmobiData(r io.Reader, totalPage *int) ([]*Product, error) {
	dec := json.NewDecoder(r)
	products := make([]*Product, 0, 64)
	errCnt := 0
	for {
		t, err := dec.Token() // 令牌返回输入流中的下一个JSON令牌。在输入流结束时，令牌返回nil，io.EOF
		if err != nil {
			return nil, fmt.Errorf("dec token err: %v", err)
		}
		if _, ok := t.(json.Delim); ok {
			if dec.More() {
				continue
			}
			break
		}

		if key, ok := t.(string); ok {
			switch key {
			case "flag":
				t, err = dec.Token()
				if err != nil {
					return nil, fmt.Errorf("unexpected error when parsing flag: %v", err)
				}
				if v, ok := t.(string); ok {
					if v == "fail" {
						// 访问API失败，设置totalPage为0
						*totalPage = 0
						return nil, fmt.Errorf("call api fail")
					}
				} else {
					return nil, fmt.Errorf("unexpectedd flag type: %T, %v", t, t)
				}
			case "msg":
				t, err = dec.Token()
				if err != nil {
					return nil, fmt.Errorf("unexpected error when parsing msg: %v", err)
				}
				if v, ok := t.(string); ok {
					if !strings.Contains(v, "success") {
						return nil, fmt.Errorf("call api msg: %v", v)
					}
				} else {
					return nil, fmt.Errorf("unexpecetd msg type: %T %v", t, t)
				}
			case "data":
				t, err = dec.Token()
				fmt.Printf("%T: %v", t, t)
				if err != nil {
					return nil, fmt.Errorf("unexpected error when reading data open bracket: %v", err)
				}

				for dec.More() {
					var item Product
					if err := dec.Decode(&item); err != nil {
						fmt.Println("[JSON] unexpected error when reading data item: ", err)
						if errCnt > 50 {
							return nil, fmt.Errorf("json decode err: %v", err)
						}
						errCnt++
						continue
					}

					products = append(products, &item)
				}
				t, err = dec.Token()
				fmt.Printf("%T: %v", t, t)
				if err != nil {
					fmt.Println("unexpected error when reading close bracket: ", err)
				}
			case "totalPage":
				t, err = dec.Token()
				if err != nil {
					return nil, fmt.Errorf("unexpected error when parsing totalPage: %v", err)
				}
				if v, ok := t.(float64); ok {
					*totalPage = int(v)
				} else {
					return nil, fmt.Errorf("unexpecetd totalPage type: %T %v", t, t)
				}
			}
		}
	}

	return products, nil
}

func main() {
	url := "http://api.yeahmobi.com/v1/Apps/get?api_token=f3bed309b22&devapp_id=283&limit=2&page=1"
	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
	}

	totalPage := 1
	products, err := parseYeahmobiData(resp.Body, &totalPage)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(products))
}
