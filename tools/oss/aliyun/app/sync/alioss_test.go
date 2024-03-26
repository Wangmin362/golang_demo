package sync

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/golang/demo/tools"
	"log"
	"testing"
)

var bucket *oss.Bucket

func init() {
	var err error
	syncDir, err := tools.GetEnvVar(SyncDirKey)
	if err != nil {
		log.Fatalf("%s\n", syncDir)
	}

	endpoint, err := tools.GetEnvVar(EndpointKey)
	if err != nil {
		log.Fatalf("%s\n", syncDir)
	}
	bucketName, err := tools.GetEnvVar(BucketKey)
	if err != nil {
		log.Fatalf("%s\n", syncDir)
	}
	ossId, err := tools.GetEnvVar(OssIDKey)
	if err != nil {
		log.Fatalf("%s\n", syncDir)
	}
	ossSecret, err := tools.GetEnvVar(OssSecretKey)
	if err != nil {
		log.Fatalf("%s\n", syncDir)
	}

	// 创建阿里云OSS客户端
	client, err := oss.New(fmt.Sprintf("https://%s", endpoint), ossId, ossSecret)
	if err != nil {
		log.Fatalf("create aliyun oss client error:%s", err)
	}

	// 判断指定的桶是否存在
	exist, err := client.IsBucketExist(bucketName)
	if err != nil || !exist {
		log.Fatalf("query %s bucket exist error:%s", bucketName, err)
	}

	// 获取当前桶
	bucket, err = client.Bucket(bucketName)
	if err != nil {
		log.Fatalf("get %s bucket error:%s", bucketName, err)
	}
}

func TestAddObjTag(t *testing.T) {
	file := "test_data/img.png"
	objKey := file
	if err := SaveToAliOSS(file, objKey, bucket); err != nil {
		t.Fatal(err)
	}

	chat, err := NewWeChat()
	if err != nil {
		t.Fatal(err)
	}

	url, err := chat.ImageUpload(file)
	if err != nil {
		t.Fatal(err)
	}

	tagKey := WeChatURLTagName
	err = AddObjTag(objKey, tagKey, url, bucket)
	if err != nil {
		t.Fatal(err)
	}

	tag, exist := GetObjTag(objKey, tagKey, bucket)
	if !exist {
		t.Fatalf("%s obj not found %s tag", objKey, tagKey)
	}

	if url != tag {
		t.Fatalf("wechat[%s] not equal aliOss[%s]", url, tag)
	}
}
