package logrus

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

// 参考：
//https://juejin.cn/post/6844904061393698823
//https://www.cnblogs.com/jiujuan/p/15542743.html

//func init() {
//	log.SetOutput(os.Stdout)
//	log.SetFormatter(&log.JSONFormatter{
//		TimestampFormat: "2006-01-02 15:04:05",
//		DataKey:         "test",
//	})
//	log.SetLevel(log.WarnLevel)
//}

func Use() {
	log.WithFields(log.Fields{"game": "mines", "func": "reward"}).Infof("start infof:%v", 111)
	log.WithFields(log.Fields{"game": "mines", "func": "reward"}).Warnf("start Warnf:%v", 222)
}

func Use2() {
	writer1 := &bytes.Buffer{}
	writer1.Write([]byte("11"))
	writer2 := os.Stdout
	writer3, err := os.OpenFile("./logrus_test_out.log", os.O_WRONLY|os.O_CREATE, 0744)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	log.Infof("test123 ")
	fmt.Println(writer1, "11")
}
