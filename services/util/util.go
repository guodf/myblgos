package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/uuid"
)

// NewUUID 生成新的UUID
func NewUUID(separated ...string) string {
	uID, _ := uuid.NewUUID()
	uIDStr := uID.String()
	if separated != nil {
		return strings.ReplaceAll(uIDStr, "-", separated[0])
	}
	return uIDStr
}

// MD5 获取md5值
func MD5(str string) string {
	// h := md5.New()
	// h.Sum([]byte(str))
	// bytes:=h.Sum(nil)
	bytes := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", bytes)
}

// FileToken 计算文件唯一标识
func FileToken(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	defer f.Close()
	if err != nil {
		return "", err
	}

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	bytes := h.Sum(nil)
	return hex.EncodeToString(bytes), nil
}

// RootPath 获取程序启动目录
func RootPath() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(file))
}
