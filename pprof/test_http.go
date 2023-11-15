package pprof

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"strings"
	"sync"
	"time"
)

func TestHttp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/fib/", fibHandler)       // 正常的业务http请求处理函数
	mux.HandleFunc("/repeat/", repeatHandler) // 正常的业务http请求处理函数

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	newProfileHttpServer(":9999") // pprof处理函数

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	// 模拟业务请求:
	do()
}

func doHTTPRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("ret:", len(data))
	resp.Body.Close()
}

func do() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			doHTTPRequest(fmt.Sprintf("http://localhost:8080/fib/%d", rand.Intn(30)))
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			doHTTPRequest(fmt.Sprintf("http://localhost:8080/repeat/%s/%d", generate(rand.Intn(200)), rand.Intn(200)))
			time.Sleep(500 * time.Millisecond)
		}
	}()
	wg.Wait()
}

func newProfileHttpServer(addr string) {
	go func() {
		log.Fatalln(http.ListenAndServe(addr, nil))
	}()
}

func fibHandler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Path[len("/fib/"):])
	if err != nil {
		responseError(w, err)
		return
	}

	var result int
	for i := 0; i < 1000; i++ {
		result = fib(n)
	}
	response(w, result)
}

func repeatHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.SplitN(r.URL.Path[len("/repeat/"):], "/", 2)
	if len(parts) != 2 {
		responseError(w, errors.New("invalid params"))
		return
	}

	s := parts[0]
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		responseError(w, err)
		return
	}

	var result string
	for i := 0; i < 1000; i++ {
		result = repeat(s, n)
	}
	response(w, result)
}

func response(w http.ResponseWriter, res interface{}) {

}

func responseError(w http.ResponseWriter, err error) {

}
