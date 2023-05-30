package interfaceRouter

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
)

type Router struct {
	Method     string
	Path       string
	HandleFunc http.HandlerFunc
}

func Use() {
	// 模拟路由器配置信息
	routerConfigs := []Router{
		{
			Method: "GET",
			Path:   "/",
			HandleFunc: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Welcome to the homepage!")
			},
		},
		{
			Method: "GET",
			Path:   "/about",
			HandleFunc: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Welcome to the about page!")
			},
		},
		{
			Method: "GET",
			Path:   "/contact",
			HandleFunc: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Welcome to the contact page!")
			},
		},
	}

	// 创建路由器并配置路由信息
	router := http.NewServeMux()
	for _, config := range routerConfigs {
		handler := reflect.ValueOf(config.HandleFunc)
		router.HandleFunc(config.Path, func(w http.ResponseWriter, r *http.Request) {
			// 判断请求方法是否匹配
			if r.Method == config.Method {
				// 调用处理函数
				args := []reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)}
				handler.Call(args)
			} else {
				http.NotFound(w, r)
			}
		})
	}

	// 在 8080 端口启动 HTTP 服务器
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
