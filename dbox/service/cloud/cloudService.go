package cloudService

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sdvdxl/dbox/dbox/config"
	. "github.com/sdvdxl/dbox/dbox/log"
	"github.com/sdvdxl/dbox/dbox/model"
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
}

func initAliOss(aliOss model.AliOss) error {
	var err error
	ossClient, err = oss.New(aliOss.Endpoint, aliOss.AccessKeyID, aliOss.AccessKeySecret)
	if err != nil {
		return err
	}
	bucket, err := ossClient.GetBucketInfo(aliOss.Bucket)
	if err != nil {
		return err
	}

	Log.Info(aliOss.Bucket, " info ", bucket.BucketInfo)
	return nil
}

// Upload 上传文件
func (fm *AliOssFileManager) Upload(file, path string) error {
	Log.Info("oss upload file,file:", file)

	bucket, err := ossClient.Bucket(config.Cfg.AliOss.Bucket)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(path, file)
	if err != nil {
		return err
	}
	return nil
}

func (fm *AliOssFileManager) Delete(path string) error {
	Log.Info("oss upload file,path:", path)
	return nil
}

func (fm *AliOssFileManager) Move(path, p string) error {
	Log.Info("oss upload file,path:", path)
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
	Log.Info("QiNiuFileManager upload file,path:", path)
	return nil
}

func (oss *QiNiuFileManager) Delete(path string) error {
	Log.Info("QiNiuFileManager upload file,path:", path)
	return nil
}

func (oss *QiNiuFileManager) Move(path, p string) error {
	Log.Info("QiNiuFileManager upload file,path:", path)
	return nil
}

func UseCloudFileManager(fm CloudFileManager) error {
	if fm == nil {
		panic("CloudFileManager can not be null")
	}

	cfm = fm
	switch fm.(type) {
	case *AliOssFileManager:
		return initAliOss(config.Cfg.AliOss)
	}
	return errors.New("not support")
}

func Upload(file, category string) error {
	return cfm.Upload(file, category)
}
