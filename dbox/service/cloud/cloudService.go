package cloudService

import (
	"errors"
	"fmt"
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
	Upload(file, fileName, category string) error

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
func (fm *AliOssFileManager) Upload(file, fileName, category string) error {
	Log.Info("oss upload file,file:", file)

	bucket, err := ossClient.Bucket(config.Cfg.AliOss.Bucket)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(category+"/"+fileName, file, oss.Progress(&OssProgressListener{}))
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
func (oss *QiNiuFileManager) Upload(file, fileName, category string) error {
	Log.Info("QiNiuFileManager upload file,path:", category)
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

func Upload(file, fileName, category string) error {
	return cfm.Upload(file, fileName, category)
}

// 定义进度条监听器。
type OssProgressListener struct {
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		fmt.Printf("Transfer Started, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferDataEvent:
		fmt.Printf("\rTransfer Data, ConsumedBytes: %d, TotalBytes %d, %d%%.",
			event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
	case oss.TransferCompletedEvent:
		fmt.Printf("\nTransfer Completed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferFailedEvent:
		fmt.Printf("\nTransfer Failed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	default:
	}
}
