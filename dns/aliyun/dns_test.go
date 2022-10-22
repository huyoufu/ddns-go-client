package aliyun

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"
)

//go:embed ali_key.txt
var ali_key_secret string

func TestAliDNSClient_CreateDNSRecord(t *testing.T) {
	keySecretArr := strings.Split(ali_key_secret, "\n")
	client := AliDNSClientOf(keySecretArr[0], keySecretArr[1])
	recordId, err := client.CreateDNSRecord("jk1123.com", "xxx", "AAAA", "2409:8a44:2fd0:e830:353f:10e6:6103:992e")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("记录id:", recordId)
	}
}
func TestAliDNSClient_UpdateDNSRecord(t *testing.T) {
	keySecretArr := strings.Split(ali_key_secret, "\n")
	client := AliDNSClientOf(keySecretArr[0], keySecretArr[1])
	client.UpdateDNSRecord("v8", "AAAA", "758035757678285824", "2409:8a44:2fd0:e830:353f:10e6:6103:992e")
}
func TestAliDNSClient_GetDNSRecords(t *testing.T) {
	keySecretArr := strings.Split(ali_key_secret, "\n")
	client := AliDNSClientOf(keySecretArr[0], keySecretArr[1])
	records, err := client.GetDNSRecords("jk1123.com", "xxx", 1, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(records)
	}
}
