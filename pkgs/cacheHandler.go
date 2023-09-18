package pkgs

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"
)

type CacheHandler struct {
	Dir string
}

func (c *CacheHandler) GetCacheFilepath(url string, contentType string) string {
	h := md5.New()
	_, err := io.WriteString(h, url)
	if err != nil {
		panic("md5失败")
	}
	md5 := h.Sum(nil)
	ext := "jpg"
	if strings.Contains(contentType, "png") {
		ext = "png"
	}
	filename := fmt.Sprintf("%x", md5) + "." + ext
	return strings.TrimRight(c.Dir, "/") + "/" + filename
}

func (c *CacheHandler) GetFileBytes(path string) []byte {
	_, err := os.Lstat(path)
	if err != nil {
		return nil
	}
	buf, rErr := os.ReadFile(path)
	if rErr != nil {
		return nil
	}
	return buf
}

func (c *CacheHandler) SaveFile(path string, buf []byte) {
	os.WriteFile(path, buf, 0660)
}

func (c *CacheHandler) Init() {
	mErr := os.MkdirAll(c.Dir, 0750)
	if mErr != nil {
		panic("创建缓存目录失败")
	}
}
