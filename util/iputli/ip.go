package iputli

import (
	"encoding/json"
	"errors"
	. "github.com/huyoufu/ddns-go-client/logger"
	"net/http"
	"time"
)

type Ip struct {
	Ip      string `json:"ip"`
	Type    string `json:"type"`
	Subtype string `json:"subtype"`
	Via     string `json:"via"`
	Padding string `json:"padding"`
}

var ipSites = [...]string{
	"https://ipv6.jp2.test-ipv6.com/ip/?testdomain=test-ipv6.com&testname=test_aaaa",
	"https://ipv6.test-ipv6.is/ip/?testdomain=test-ipv6.is&testname=test_aaaa",
	"http://ipv6.test-ipv6.cl/ip/?testdomain=test-ipv6.cl&testname=test_aaaa",
	"http://ipv6.test-ipv6.hu/ip/?testdomain=test-ipv6.hu&testname=test_aaaa",
	"http://ipv6.test-ipv6.alpinedc.ch/ip/?testdomain=test-ipv6.alpinedc.ch&testname=test_aaaa",
	"http://ipv6.test-ipv6.ttk.ru/ip/?testdomain=test-ipv6.ttk.ru&testname=test_aaaa",
	"http://ipv6.test-ipv6.com.au/ip/?testdomain=test-ipv6.com.au&testname=test_aaaa",
}

func GetIp() (result *Ip, err error) {
	for _, ipSite := range ipSites {
		Log.Infof("正在从:%s获取ip地址", ipSite)
		result, err = getIp(ipSite)
		if err == nil {
			return result, nil
		}
	}
	return result, err
}

func getIp(ipSite string) (result *Ip, err error) {
	client := &http.Client{}
	client.Timeout = time.Millisecond * 50000
	request, err := http.NewRequest("GET", ipSite, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		Log.Warn("重定向过多!")
		return http.ErrUseLastResponse
	}
	var res *http.Response
	res, err = client.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		Log.Warn("url %s,status code error: %d %s\n", ipSite, res.StatusCode, res.Status)
		return nil, errors.New("状态码错误")
	}
	//读取内容
	result = new(Ip)
	buff := make([]byte, 1024)
	var n int
	n, err = res.Body.Read(buff)
	//callback({...})
	b := buff[9 : n-2]
	err = json.Unmarshal(b, result)
	return result, err
}
