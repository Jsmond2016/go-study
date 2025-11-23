---
title: 文件上传
difficulty: intermediate
duration: "2-3小时"
prerequisites: ["文章管理"]
tags: ["文件上传", "图片", "存储", "OSS"]
---

# 文件上传

本章节将实现文件上传功能，包括图片上传、文件管理、文件验证和云存储集成。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 实现文件上传功能
- [ ] 验证文件类型和大小
- [ ] 处理图片压缩和裁剪
- [ ] 集成云存储（OSS/S3）
- [ ] 实现文件管理功能
- [ ] 处理文件访问权限

## 📤 上传服务

创建 `pkg/upload/upload.go`:

```go
package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type UploadService interface {
	SaveFile(file *multipart.FileHeader, uploadPath string) (string, error)
	ValidateFile(file *multipart.FileHeader, maxSize int64, allowedTypes []string) error
	DeleteFile(filepath string) error
}

type LocalUploadService struct {
	BasePath string
}

func NewLocalUploadService(basePath string) UploadService {
	// 确保目录存在
	os.MkdirAll(basePath, 0755)
	return &LocalUploadService{BasePath: basePath}
}

func (s *LocalUploadService) SaveFile(file *multipart.FileHeader, subPath string) (string, error) {
	// 打开文件
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("打开文件失败: %w", err)
	}
	defer src.Close()

	// 生成文件名
	filename := generateFilename(file.Filename)
	fullPath := filepath.Join(s.BasePath, subPath, filename)

	// 确保目录存在
	os.MkdirAll(filepath.Dir(fullPath), 0755)

	// 创建目标文件
	dst, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("创建文件失败: %w", err)
	}
	defer dst.Close()

	// 复制文件内容
	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("保存文件失败: %w", err)
	}

	// 返回相对路径
	return filepath.Join(subPath, filename), nil
}

func (s *LocalUploadService) ValidateFile(file *multipart.FileHeader, maxSize int64, allowedTypes []string) error {
	// 检查文件大小
	if file.Size > maxSize {
		return fmt.Errorf("文件大小超过限制: %d bytes", maxSize)
	}

	// 检查文件类型
	contentType := file.Header.Get("Content-Type")
	if !contains(allowedTypes, contentType) {
		return fmt.Errorf("不支持的文件类型: %s", contentType)
	}

	return nil
}

func (s *LocalUploadService) DeleteFile(filepath string) error {
	fullPath := filepath.Join(s.BasePath, filepath)
	return os.Remove(fullPath)
}

func generateFilename(originalName string) string {
	ext := filepath.Ext(originalName)
	timestamp := time.Now().UnixNano()
	randomStr := generateRandomString(8)
	return fmt.Sprintf("%d_%s%s", timestamp, randomStr, ext)
}

func generateRandomString(length int) string {
	// 简单的随机字符串生成
	return fmt.Sprintf("%x", time.Now().UnixNano())[:length]
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
```

## 🖼️ 图片处理

创建 `pkg/upload/image.go`:

```go
package upload

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"github.com/nfnt/resize"
)

func ResizeImage(srcPath, dstPath string, width, height uint) error {
	// 打开源文件
	file, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 解码图片
	var img image.Image
	ext := filepath.Ext(srcPath)
	switch ext {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	default:
		return fmt.Errorf("不支持的图片格式: %s", ext)
	}
	if err != nil {
		return err
	}

	// 调整大小
	resized := resize.Resize(width, height, img, resize.Lanczos3)

	// 保存调整后的图片
	out, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer out.Close()

	switch ext {
	case ".jpg", ".jpeg":
		return jpeg.Encode(out, resized, nil)
	case ".png":
		return png.Encode(out, resized)
	}

	return nil
}

func GenerateThumbnail(srcPath, dstPath string) error {
	return ResizeImage(srcPath, dstPath, 200, 200)
}
```

## 📝 上传处理器

创建 `internal/handler/upload.go`:

```go
package handler

import (
	"net/http"
	"blog-system/pkg/upload"
	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	uploadService upload.UploadService
	config        UploadConfig
}

type UploadConfig struct {
	MaxSize      int64
	AllowedTypes []string
	BasePath     string
}

func NewUploadHandler(uploadService upload.UploadService, config UploadConfig) *UploadHandler {
	return &UploadHandler{
		uploadService: uploadService,
		config:        config,
	}
}

func (h *UploadHandler) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "获取文件失败",
		})
		return
	}

	// 验证文件
	if err := h.uploadService.ValidateFile(file, h.config.MaxSize, h.config.AllowedTypes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// 保存文件
	filepath, err := h.uploadService.SaveFile(file, "images")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "上传文件失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "上传成功",
		"data": gin.H{
			"path": filepath,
			"url":  "/uploads/" + filepath,
		},
	})
}
```

## ☁️ 云存储集成

### OSS 上传服务

创建 `pkg/upload/oss.go`:

```go
package upload

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
)

type OSSUploadService struct {
	client   *oss.Client
	bucket   *oss.Bucket
	endpoint string
	bucketName string
}

func NewOSSUploadService(endpoint, accessKeyID, accessKeySecret, bucketName string) (UploadService, error) {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	return &OSSUploadService{
		client:     client,
		bucket:     bucket,
		endpoint:   endpoint,
		bucketName: bucketName,
	}, nil
}

func (s *OSSUploadService) SaveFile(file *multipart.FileHeader, subPath string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	filename := generateFilename(file.Filename)
	objectKey := fmt.Sprintf("%s/%s", subPath, filename)

	err = s.bucket.PutObject(objectKey, src)
	if err != nil {
		return "", err
	}

	return objectKey, nil
}

func (s *OSSUploadService) ValidateFile(file *multipart.FileHeader, maxSize int64, allowedTypes []string) error {
	if file.Size > maxSize {
		return fmt.Errorf("文件大小超过限制")
	}
	// 其他验证逻辑
	return nil
}

func (s *OSSUploadService) DeleteFile(filepath string) error {
	return s.bucket.DeleteObject(filepath)
}
```

## 🔧 路由设置

```go
func setupUploadRoutes(r *gin.RouterGroup, uploadHandler *handler.UploadHandler) {
	upload := r.Group("/upload")
	upload.Use(auth.AuthMiddleware())
	{
		upload.POST("/image", uploadHandler.UploadImage)
		upload.POST("/file", uploadHandler.UploadFile)
		upload.DELETE("/:path", uploadHandler.DeleteFile)
	}
}
```

## 📝 API 使用示例

### 上传图片

```bash
curl -X POST http://localhost:8080/api/upload/image \
  -H "Authorization: Bearer <token>" \
  -F "file=@/path/to/image.jpg"
```

### 响应

```json
{
  "success": true,
  "message": "上传成功",
  "data": {
    "path": "images/1234567890_abc123.jpg",
    "url": "/uploads/images/1234567890_abc123.jpg"
  }
}
```

## 💡 最佳实践

### 1. 文件验证

- **类型验证**: 检查文件MIME类型
- **大小限制**: 限制文件大小
- **文件名验证**: 防止路径遍历攻击

### 2. 存储策略

- **本地存储**: 适合小规模应用
- **云存储**: 适合大规模应用
- **CDN加速**: 使用CDN加速文件访问

### 3. 安全考虑

- **权限控制**: 限制上传权限
- **文件扫描**: 扫描恶意文件
- **访问控制**: 控制文件访问权限

## ⏭️ 下一步

文件上传完成后，下一步是：

- [搜索功能](./07-search.md) - 实现全文搜索功能

---

**🎉 文件上传完成！** 现在你可以开始实现搜索功能了。

