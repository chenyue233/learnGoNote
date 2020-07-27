package fetcher

import (
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding/unicode"
	"log"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)
// 发起HTTP请求
func Fetch(url string) ([]byte,error) {
	<- rateLimiter
	clien := &http.Client{}
	rep,err := http.NewRequest("GET",url,nil)
	if err != nil {
		log.Printf("Get Url error %s",url)
	}
	rep.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.75 Safari/537.36")
	rep.Header.Add("Host","album.zhenai.com")
	resp,err := clien.Do(rep)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("wrong statuscode:%d",resp.StatusCode)
	}
	bodyRead := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyRead)
	utf8Reader := transform.NewReader(bodyRead,e.NewDecoder()) // 转换格式
	return ioutil.ReadAll(utf8Reader)
}

// 判断读取的内容的格式是什么
func determineEncoding(r *bufio.Reader) encoding.Encoding{
	bytes,err := r.Peek(1024)
	if err != nil{
		log.Printf("Fetcher error:%v",err)
		return unicode.UTF8
	}
	e,_,_ :=charset.DetermineEncoding(bytes,"")
	return e
}
