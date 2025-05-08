/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-19 21:48:22
 */

package fs

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
)

var MinioClient *minio.Client

type FileSystem struct{}

var Minio = &FileSystem{}

func (*FileSystem) UploadFile(ctx context.Context, bucketName, objectName string, size int64, reader io.Reader) error {
	_, err := MinioClient.PutObject(ctx, bucketName, objectName, reader, size, minio.PutObjectOptions{})
	return err
}

func (*FileSystem) GetFile(ctx context.Context, bucketName, objectName string) ([]byte, error) {
	file, err := MinioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return io.ReadAll(file)
}
