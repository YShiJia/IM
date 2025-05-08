/**
 * @author ysj
 * @email 2239831438@qq.com
 * @createTime: 2025-04-19 21:18:04
 */

package fs

import "strings"

const (
	FileCommonBucket  = "im.file"
	FilePictureBucket = "im.picture"
	FileVideoBucket   = "im.video"
)

// 定义后缀和文件类型的映射关系
var extToFileType = map[string]string{
	".jpg":  FilePictureBucket,
	".jpeg": FilePictureBucket,
	".png":  FilePictureBucket,
	".gif":  FilePictureBucket,
	".bmp":  FilePictureBucket,
	".webp": FilePictureBucket,
	".mp4":  FileVideoBucket,
	".mkv":  FileVideoBucket,
	".avi":  FileVideoBucket,
	".mov":  FileVideoBucket,
	".wmv":  FileVideoBucket,
}

// GetBucketName 根据文件名后缀获取文件存储桶
func GetBucketName(filename string) string {
	// 获取文件后缀（小写）
	ext := strings.ToLower(getFileExtension(filename))

	// 查找后缀对应的文件类型，如果没有找到，默认返回 FileCommonBucket
	if fileType, exists := extToFileType[ext]; exists {
		return fileType
	}
	return FileCommonBucket
}

// 获取文件后缀
func getFileExtension(filename string) string {
	// 查找文件名中最后一个点的位置，并返回点后面的部分作为后缀
	index := strings.LastIndex(filename, ".")
	if index == -1 {
		return ""
	}
	return filename[index:]
}
