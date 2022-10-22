package config

type DomainNameServiceProvider int

const (
	DnsAli DomainNameServiceProvider = iota
	DnsTencent
)

type DDNSConfig struct {
	Active   string           `json:"active"`
	Interval int              `json:"interval"` //间隔时间
	Ipv4     []string         `json:"ipv4"`
	Ipv6     []string         `json:"ipv6"`
	Ali      AliDNSConfig     `json:"ali"`
	Tencent  TencentDNSConfig `json:"tencent"`
}

func (d *DDNSConfig) DetectDomainNameServiceProvider() DomainNameServiceProvider {
	if d.Active == "ali" {
		return DnsAli
	} else {
		return DnsTencent
	}

}

type AliDNSConfig struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
}
type TencentDNSConfig struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
}
