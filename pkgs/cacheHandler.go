package pkgs

import (
	"crypto/md5"
	"fmt"
	"github.com/robfig/cron/v3"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type CacheHandler struct {
	Dir      string
	CacheDay int
}

func (ch *CacheHandler) GetCacheFilepath(url string, contentType string) string {
	h := md5.New()
	_, err := io.WriteString(h, url)
	if err != nil {
		panic("md5失败")
	}
	md5Ent := h.Sum(nil)
	ext := "jpg"
	if strings.Contains(contentType, "png") {
		ext = "png"
	}
	filename := fmt.Sprintf("%x", md5Ent) + "." + ext
	return strings.TrimRight(ch.Dir, "/") + "/" + filename
}

func (ch *CacheHandler) GetFileBytes(path string) []byte {
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

func (ch *CacheHandler) SaveFile(path string, buf []byte) {
	err := os.WriteFile(path, buf, 0660)
	if err != nil {
		panic("保存缓存文件失败")
	}
}

func (ch *CacheHandler) Init() {
	mErr := os.MkdirAll(ch.Dir, 0750)
	if mErr != nil {
		panic("创建缓存目录失败")
	}

	ch.StartCron()
}

func (ch *CacheHandler) StartCron() {
	cr := cron.New(cron.WithSeconds())
	cr.AddFunc("0 0 4 * * *", ch.CleanOld)
	cr.Start()
}

func (ch *CacheHandler) CleanOld() {
	entries, err := os.ReadDir(ch.Dir)
	if err != nil {
		return
	}

	currentTime := time.Now()
	cutoffTime := currentTime.AddDate(0, 0, -ch.CacheDay)

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}

		if info.ModTime().Before(cutoffTime) {
			os.Remove(filepath.Join(ch.Dir, info.Name()))
		}
	}
}
