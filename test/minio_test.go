/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-13 01:19:03
 */

package test

import (
	"context"
	"fmt"
	"github.com/YShiJia/IM/lib/sign"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

const (
	endpoint        string = "10.120.0.60:9000"
	accessKeyID     string = "heathyang"
	secretAccessKey string = "heathyang"
	useSSL          bool   = false // 不使用 https
)

var (
	client *minio.Client
	err    error
)

func Test_minio(t *testing.T) {
	client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL})
	if err != nil {
		log.Fatalln("minio连接错误: ", err)
	}
	log.Printf("连接minio成功：%#v\n", client)
	buckets := []string{"im.picture", "im.file", "im.video"}
	for _, bucket := range buckets {
		createBucket(bucket)
	}
	//listBucket()
	FileUploader()
	//FileGet()

}

func createBucket(bucketName string) {
	err = client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "cn-south-1", ObjectLocking: false})
	if err != nil {
		log.Println("创建bucket错误: ", err)
		exists, _ := client.BucketExists(context.Background(), bucketName)
		if exists {
			log.Printf("bucket: %s已经存在", bucketName)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
}

func listBucket() {
	buckets, _ := client.ListBuckets(context.Background())
	for _, bucket := range buckets {
		fmt.Println("已存在bucket：", bucket)
	}
}

func FileUploader() {
	bucketName := "im.file"
	filePath := "/Design/IM/logs/tlog.log"
	contextType := "application/text"

	openFile, err := os.Open(filePath)
	if err != nil {
		log.Println("获取文件错误: ", err)
	}
	defer openFile.Close()

	objectName, err := sign.FileSignatureBySHA256(openFile)
	if err != nil {
		log.Println("获取文件hash值错误: ", err)
	}

	object, err := client.FPutObject(context.TODO(), bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contextType})
	if err != nil {
		log.Println("上传失败：", err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, object.Size)
}

func FileGet() {
	bucketName := "im.file"
	objectName := "tlog.log"
	filePath := "/Design/IM/logs/tlog3.log"

	err = client.FGetObject(context.TODO(), bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		log.Println("下载错误: ", err)
	}
}
