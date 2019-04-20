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

type Traffic struct {
	Allowed   []string `json:"allowed,omitempty"`
	Forbidden []string `json:"forbidden,omitempty"`
}

// NewProduct 新产品结构
type NewProduct struct {
	OfferID          int        `json:"offer_id,omitempty"`
	OfferName        string     `json:"offer_name,omitempty"`
	OfferCategory    []string   `json:"offer_category,omitempty"`
	OfferDescription string     `json:"offer_description,omitempty"`
	TrackingLink     string     `json:"tracking_link,omitempty"`
	IsS2s            bool       `json:"is_s2s,omitempty"`
	ConversionFlow   []string   `json:"conversion_flow,omitempty"`
	Targeting        Targeting  `json:"targeting,omitempty"`
	Traffic          Traffic    `json:"traffic,omitempty"`
	Financials       Financials `json:"financials,omitempty"`
	Icons            Icon       `json:"icons,omitempty"`
	Creatives        Creative   `json:"creatives,omitempty"`
	Videos           Video      `json:"videos,omitempty"`
	TrackingProvider string     `json:"tracking_provider,omitempty"`
	PkgName          string     `json:"pkg_name,omitempty"`
	PreviewURL       string     `json:"preview_url,omitempty"`
	OfferType        string     `json:"offer_type,omitempty"`
	Status           string     `json:"status,omitempty"`
	FileSizeBytes    int        `json:"file_size_bytes,omitempty"`
	RatingAverage    float64    `json:"rating_average,omitempty"`
	ProductSource    string     `json:"product_source,omitempty"`
}

// Targeting 定向
type Targeting struct {
	Countries []string `json:"countries,omitempty"`
	Platforms []string `json:"platforms,omitempty"`
}

// Financials 金融
type Financials struct {
	CapDaily          int     `json:"cap_daily,omitempty"`
	CapMonthly        int     `json:"cap_monthly,omitempty"`
	CapShared         bool    `json:"cap_shared,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	Payout            float64 `json:"payout,omitempty"`
	RemainingCapDaily int     `json:"remaining_cap_daily,omitempty"`
}

func parseYeahmobiData(r io.Reader, totalPage *int) ([]*NewProduct, error) {
	dec := json.NewDecoder(r)
	products := make([]*NewProduct, 0, 64)
	errCnt := 0
	for {
		t, err := dec.Token() // 令牌返回输入流中的下一个JSON令牌。在输入流结束时，令牌返回nil，io.EOF
		fmt.Printf("%T: %v", t, t)
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

				if dec.More() {
					t, err = dec.Token()
					fmt.Printf("%T: %v", t, t)
					fmt.Println(err)
					if key, ok := t.(string); ok {
						switch key {
						case "data":
							t, err = dec.Token()
							fmt.Printf("%T: %v", t, t)
							if err != nil {
								return nil, fmt.Errorf("unexpected error when reading data open bracket: %v", err)
							}

							for dec.More() {
								var item NewProduct
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
						}
					}
				}
			case "totalOffer":
				fmt.Println("totalOffer")
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
	url := "http://api.yeahmobi.com/v1/Offers/get?api_token=f3bed309b22&devapp_id=283"
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
	fmt.Println("--------------------------")
	fmt.Println(len(products))
}
