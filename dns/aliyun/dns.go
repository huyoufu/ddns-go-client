package aliyun

import (
	alidns "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/huyoufu/ddns-go-client/config"
	"github.com/huyoufu/ddns-go-client/dns"
	"github.com/huyoufu/ddns-go-client/logger"
	"github.com/huyoufu/ddns-go-client/util/iputli"
	"time"
)

type AliDNSConfig struct {
	// accessKey id
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// accessKey secret
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
}
type AliDNSClient struct {
	client alidns.Client
}

func AliDNSClientOf(accessKeyId string, accessKeySecret string) (aliDNSClient *AliDNSClient) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: tea.String(accessKeyId),
		// 您的 AccessKey Secret
		AccessKeySecret: tea.String(accessKeySecret),
		// 访问的域名
		Endpoint: tea.String("alidns.cn-hangzhou.aliyuncs.com"),
	}
	client, _ := alidns.NewClient(config)

	return &AliDNSClient{*client}
}

func (c *AliDNSClient) CreateDNSRecord(domainName, rr, Type, value string) (string, error) {
	request := &alidns.AddDomainRecordRequest{
		//语言
		Lang:       tea.String("en"),
		DomainName: tea.String(domainName),
		RR:         tea.String(rr),
		Type:       tea.String(Type),
		Value:      tea.String(value),
	}
	resp, e := c.client.AddDomainRecord(request)
	if e != nil {
		return "", e
	} else {
		return *resp.Body.RecordId, e
	}
}
func (c *AliDNSClient) GetDNSRecords(domainName, rrKeyWord string, pageNumber, pageSize int64) (records *alidns.DescribeDomainRecordsResponseBody, e error) {
	request := &alidns.DescribeDomainRecordsRequest{
		//语言
		Lang:       tea.String("en"),
		DomainName: tea.String(domainName),
		RRKeyWord:  tea.String(rrKeyWord),
		PageNumber: tea.Int64(pageNumber),
		PageSize:   tea.Int64(pageSize),
	}
	resp, e := c.client.DescribeDomainRecords(request)
	return resp.Body, e
}

func (c *AliDNSClient) UpdateDNSRecord(RR string, Type string, RecordId string, Value string) (_err error) {
	request := &alidns.UpdateDomainRecordRequest{
		//语言
		Lang:     tea.String("en"),
		RR:       tea.String(RR),
		Type:     tea.String(Type),
		RecordId: tea.String(RecordId),
		Value:    tea.String(Value),
		//专用路线 企业级可以设置专用的解析路线
		Line: tea.String("default"),
	}
	_, _err = c.client.UpdateDomainRecord(request)
	return _err
}

var ddnsConfig *config.DDNSConfig

func DNSTask(c *config.DDNSConfig) {
	ddnsConfig = c
	task()
}
func task() {
	logger.Log.Info("定时任务开始了")
	aliDNSClient := AliDNSClientOf(ddnsConfig.Ali.AccessKeyId, ddnsConfig.Ali.AccessKeySecret)
	//先查询
	ipv6s := ddnsConfig.Ipv6
	for _, ipv6 := range ipv6s {
		logger.Log.Infof("正在为域名:%s进行同步ddns记录", ipv6)
		domainStruct := dns.Analysis(ipv6)
		records, err := aliDNSClient.GetDNSRecords(domainStruct.Main, domainStruct.RR, 1, 10)
		if err != nil {
			logger.Log.Error(err)
		} else {
			if *records.TotalCount == 0 {
				//如果为0 说明 没有创建 所以创建新记录就行了
				ip, err := iputli.GetIp()
				if err != nil {
					logger.Log.Error("获取ip地址失败", err)
				} else {
					record, err := aliDNSClient.CreateDNSRecord(domainStruct.Main, domainStruct.RR, "AAAA", ip.Ip)
					if err != nil {
						logger.Log.Errorf("创建dns记录:%s 失败", ipv6)
					} else {
						logger.Log.Infof("创建dns记录:%s 成功,recordId是:%s", ipv6, record)
					}
				}
			} else {
				//如果不为0  我们这里认为最多只会只有一条记录的
				record := records.DomainRecords.Record[0]
				ip, err := iputli.GetIp()
				if err != nil {
					logger.Log.Error("获取ip地址失败", err)
					return
				}
				logger.Log.Infof("获取的ip地址为:%s", ip.Ip)

				if *record.Value == ip.Ip {
					//如果记录相同 不做事情
					logger.Log.Infof("当前ip和dns记录的值:%s相同,无需更新!", ip.Ip)
				} else {
					//不相同就更新
					err := aliDNSClient.UpdateDNSRecord(domainStruct.RR, "AAAA", *record.RecordId, ip.Ip)
					if err != nil {
						logger.Log.Error("更新失败", err)
					} else {
						logger.Log.Infof("域名:%s,值为:%s,更新成功", ipv6, ip.Ip)
					}

				}
			}
		}
	}

	time.AfterFunc(time.Minute*time.Duration(ddnsConfig.Interval), task)
}
