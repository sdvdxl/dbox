package service

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/getlantern/errors"
	"github.com/gohugoio/hugo/helpers"
	"github.com/sdvdxl/dbox/dbox/db"
	"github.com/sdvdxl/dbox/dbox/ex"
	. "github.com/sdvdxl/dbox/dbox/log"
	"github.com/sdvdxl/dbox/dbox/model"
	"os"
	"strings"
)

var (
	cfm       CloudFileManager
	ossClient *oss.Client
)

// CloudFileManager 文件管理
type CloudFileManager interface {
	// Upload 上传文件
	Upload(file, path string) error

	// Delete 删除文件
	Delete(path string) error

	// Move 移动文件，重命名文件
	Move(oldPath, newPath string) error
}

// AliOssFileManager 实现了 CloudFileManager 接口
type AliOssFileManager struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	Options         oss.ClientOption
	Bucket          string
}

func initAliOss(fm AliOssFileManager) error {
	var err error
	ossClient, err = oss.New(fm.Endpoint, fm.AccessKeyID, fm.AccessKeySecret)
	if err != nil {
		return err
	}
	bucket, err := ossClient.GetBucketInfo(fm.Bucket)
	if err != nil {
		return err
	}

	Log.Info(fm.Bucket, " info ", bucket.BucketInfo)
	return nil
}

// Upload 上传文件
func (fm *AliOssFileManager) Upload(file, path string) error {
	fmt.Println("oss upload file,file:", file)

	bucket, err := ossClient.Bucket(fm.Bucket)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(path, file)
	if err != nil {
		return err
	}
	return nil
}

func (oss *AliOssFileManager) Delete(path string) error {
	fmt.Println("oss upload file,path:", path)
	return nil
}

func (oss *AliOssFileManager) Move(path, p string) error {
	fmt.Println("oss upload file,path:", path)
	return nil
}

// QiNiuFileManager 实现了 CloudFileManager 接口
type QiNiuFileManager struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	Options         oss.ClientOption
}

// Upload 上传文件
func (oss *QiNiuFileManager) Upload(file, path string) error {
	fmt.Println("QiNiuFileManager upload file,path:", path)
	return nil
}

func (oss *QiNiuFileManager) Delete(path string) error {
	fmt.Println("QiNiuFileManager upload file,path:", path)
	return nil
}

func (oss *QiNiuFileManager) Move(path, p string) error {
	fmt.Println("QiNiuFileManager upload file,path:", path)
	return nil
}

func UseCloudFileManager(fm CloudFileManager) error {
	if fm == nil {
		panic("CloudFileManager can not be null")
	}

	cfm = fm
	switch fm.(type) {
	case *AliOssFileManager:
		return initAliOss(*fm.(*AliOssFileManager))
	}
	return errors.New("not support")
}

// UploadLocalFile 上传本地文件
func UploadLocalFile(file, category string, tags ...string) error {
	Log.Infow("upload file info", "file", file, "category", category, "tags", tags)
	stat, err := os.Stat(file)
	if err != nil {
		return ex.FileNotExistErr.Arg(file)
	}

	if stat.IsDir() {
		return ex.FileErr.Arg("not support dir, ", file)
	}

	f, err := os.Open(file)
	if err != nil {
		return ex.FileErr.Arg(err)
	}

	md5Sum, err := helpers.MD5FromFile(f)
	if err != nil {
		return ex.FileErr.Arg(err)
	}

	category = strings.TrimSpace(category)
	if category == "" {
		category = model.CatetoryRoot
	}
	c := db.Category.Save(category)
	db.File.Save(&model.File{Name: file, CategoryID: c.ID, MD5: md5Sum, Path: category})

	return cfm.Upload(file, category)

}
