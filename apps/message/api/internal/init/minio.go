/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-04-29 12:17:48
 */

package init

import (
	"context"
	"fmt"
	conf "github.com/YShiJia/IM/apps/message/api/internal/config"
	"github.com/YShiJia/IM/apps/message/api/internal/dao/fs"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
)

func InitMinio() error {
	client, err := minio.New(
		fmt.Sprintf("%s:%d", conf.Conf.MinioConf.Host, conf.Conf.MinioConf.Port),
		&minio.Options{
			Creds:  credentials.NewStaticV4(conf.Conf.MinioConf.AccessKeyID, conf.Conf.MinioConf.SecretAccessKey, ""),
			Secure: conf.Conf.MinioConf.UseSSL,
		})
	if err != nil {
		log.Errorf("minio连接错误 err %v", err)
		return err
	}
	log.Infof("连接minio成功：%#v\n", client)
	fs.MinioClient = client
	buckets := []string{fs.FileCommonBucket, fs.FilePictureBucket, fs.FileVideoBucket}
	for _, bucket := range buckets {
		if err := initBucket(bucket); err != nil {
			log.Errorf("初始化bucket[%s]失败 err: %v", bucket, err)
			return fmt.Errorf("初始化bucket[%s]失败 err: %v", bucket, err)
		}
	}
	log.Infof("初始化 minio 成功")
	return nil
}

func initBucket(bucketName string) error {
	ctx := context.TODO()
	// 先判断是否存在
	exists, err := fs.MinioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Infof("获取bucket[%s]存在信息失败 err: %v", bucketName, err)
		return fmt.Errorf("获取bucket[%s]存在信息失败 err: %v", bucketName, err)
	}
	if exists {
		log.Infof("bucket[%s] 已存在无需创建", bucketName)
		return nil
	}
	if err = fs.MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "cn-south-1", ObjectLocking: false}); err != nil {
		log.Infof("创建bucket[%s]失败 err: %v", bucketName, err)
		return fmt.Errorf("创建bucket[%s]失败 err: %v", bucketName, err)
	}

	log.Infof("创建bucket[%s] 成功", bucketName)
	return nil
}
