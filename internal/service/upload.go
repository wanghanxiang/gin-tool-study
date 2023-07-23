package service

import (
	"context"
	"mime/multipart"
	"product-mall/conf"
	"product-mall/pkg/e"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// 上传图片到七牛云 返回下载url
//https://developer.qiniu.com/kodo/1238/go
func Upload2QiNiu(file multipart.File, fileSize int64) (int, string) {
	var AccessKey = conf.AccessKey
	var SerectKey = conf.SerectKey
	var Bucket = conf.Bucket
	var ImgUrl = conf.QiniuServer
	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SerectKey)
	upToken := putPlicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, // 空间对应的机房
		UseCdnDomains: false,               // 上传是否使用CDN上传加速
		UseHTTPS:      false,               // 是否使用https域名
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		code := e.ErrorUploadFile
		return code, err.Error()
	}
	url := ImgUrl + ret.Key
	return 200, url
}
