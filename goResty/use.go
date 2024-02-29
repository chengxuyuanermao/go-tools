package goResty

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"time"
)

func SimpleGet() {
	client := resty.New() // 创建一个restry客户端
	resp, err := client.R().EnableTrace().Get("https://httpbin.org/get")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	// fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	// fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())
}

func StrongGet() {
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"page_no":      "1",
			"limit":        "20",
			"internalSort": "name",
			"order":        "asc",
			"random":       strconv.FormatInt(time.Now().Unix(), 10),
		}).
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("http://127.0.0.1:10240/ping")
	fmt.Println(resp, err)
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  header      :", resp.Header())

	// Request.SetQueryString method
	resp, err = client.R().
		SetQueryString("productId=232&template=fresh-sample&cat=resty&source=google&kw=buy a lot more").
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("http://127.0.0.1:10240/ping")
	fmt.Println(resp, err)

	// 解析返回的内容，内容是json解析到struct
	//resp, err = client.R().
	//	SetResult(result).
	//	ForceContentType("application/json").
	//	Get("v2/alpine/mainfestes/latest")
	//fmt.Println(resp, err)
}

func Post() {
	// POST Map, default is JSON content type. No need to set oneresp, err := client.R().
	client := resty.New()
	resp, err := client.R().
		SetBody(map[string]interface{}{"username": "testuser", "password": "testpass"}).
		SetResult(&result{}).
		Post("http://127.0.0.1:10240/ping")
	fmt.Println(resp, err)
}

type result struct{}

func TestMiddleware() {
	client := resty.New().OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		// Now you have access to Client and current Request object
		// manipulate it as per your need
		fmt.Println("请求前。。。")
		return nil // if its success otherwise return error
	})
	client.OnAfterResponse(func(c *resty.Client, req *resty.Response) error {
		// Now you have access to Client and current Request object
		// manipulate it as per your need
		fmt.Println("请求后。。。")
		return nil // if its success otherwise return error
	})
	resp, err := client.R().
		SetBody(map[string]interface{}{"username": "testuser", "password": "testpass"}).
		SetResult(&result{}).
		Post("http://127.0.0.1:10240/ping")
	fmt.Println(resp, err)
}

func TestRetry() {
	client := resty.New().
		SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(3 * time.Second).
		SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
			return 0, errors.New("quota exceeded")
		})
	resp, err := client.R().
		SetBody(map[string]interface{}{"username": "testuser", "password": "testpass"}).
		SetResult(&result{}).
		Post("http://127.0.0.1:10240/ping")
	fmt.Println(resp, err)
}
