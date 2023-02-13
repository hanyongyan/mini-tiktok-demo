package cos

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

// Author: Hanyongyan

var Cos *cos.Client

// Init cos 初始化
func Init() {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse(viper.GetString("cos.url"))
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  viper.GetString("cos.secretid"),
			SecretKey: viper.GetString("cos.secretkey"),
		},
	})
	Cos = c
}

// LogStatus 用于检测是否报错
func LogStatus(err error) {
	if err == nil {
		return
	}
	if cos.IsNotFoundError(err) {
		// WARN
		fmt.Println("WARN: Resource is not existed")
	} else if e, ok := cos.IsCOSError(err); ok {
		fmt.Printf("ERROR: Code: %v\n", e.Code)
		fmt.Printf("ERROR: Message: %v\n", e.Message)
		fmt.Printf("ERROR: Resource: %v\n", e.Resource)
		fmt.Printf("ERROR: RequestId: %v\n", e.RequestID)
		// ERROR
	} else {
		fmt.Printf("ERROR: %v\n", err)
		// ERROR
	}
}
