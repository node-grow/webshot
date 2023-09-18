package main

import (
	"encoding/json"
	"flag"
	"main/pkgs"
	"net/http"
	"os"
	"strings"
)

var cacher pkgs.CacheHandler
var trans pkgs.UrlTransformer
var webs pkgs.Webshoter

func main() {
	cacheDir := flag.String("cache-dir", strings.TrimRight(os.TempDir(), "/")+"/webshot", "图片缓存路径")
	flag.Parse()

	cacher = pkgs.CacheHandler{
		Dir: *cacheDir,
	}
	cacher.Init()
	println(cacher.Dir)
	trans = pkgs.UrlTransformer{}
	webs = pkgs.Webshoter{}

	http.HandleFunc("/", handle)

	hErr := http.ListenAndServe("0.0.0.0:8080", nil)
	if hErr != nil {
		println("http listen failed")
	}
}

type ResponseJson struct {
	Message any `json:"message"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			println(err)

			b, _ := json.Marshal(ResponseJson{Message: err})
			w.Write(b)
		}
	}()

	allUrl := r.URL.String()
	options, url := trans.Handle(allUrl)
	println(url)

	contentType := "image/png"
	if options["quality"] != "100" {
		contentType = "image/jpg"
	}

	fPath := cacher.GetCacheFilepath(allUrl, contentType)

	buf := cacher.GetFileBytes(fPath)
	if buf == nil {
		webs.Shot(url, options, &buf)
		cacher.SaveFile(fPath, buf)
	}
	w.Header().Add("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
