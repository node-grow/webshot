package pkgs

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"
)

type ServerHandler struct {
	Trans  UrlTransformer
	Webs   Webshoter
	Cacher CacheHandler
}

func (h *ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
		}
	}()

	fullURL := r.URL.String()
	options, url := h.Trans.Handle(fullURL)

	contentType := "image/png"
	if options["quality"] != "100" {
		contentType = "image/jpeg"
	}

	cacheFilePath := h.Cacher.GetCacheFilepath(fullURL, contentType)
	fileBytes := h.Cacher.GetFileBytes(cacheFilePath)

	if fileBytes == nil {
		h.Webs.Shot(url, options, &fileBytes)
		h.Cacher.SaveFile(cacheFilePath, fileBytes)
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("X-Powered-By", "Webshot")
	w.Header().Set("ETag", fmt.Sprintf("%x", md5.Sum(fileBytes)))
	w.Header().Set("Expires", time.Now().AddDate(0, 0, 1).Format("Mon, 02 Jan 2006 15:04:05 GMT"))
	w.WriteHeader(http.StatusOK)
	w.Write(fileBytes)
}
