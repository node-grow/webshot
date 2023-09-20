package main

import (
	"flag"
	"log"
	"main/pkgs"
	"net/http"
	"os"
	"strings"
)

func main() {
	cacheDir := flag.String(
		"cache-dir",
		strings.TrimRight(os.TempDir(), "/")+"/webshot",
		"图片缓存路径",
	)
	cacheDay := flag.Int("cache-day", 7, "图片缓存天数")
	flag.Parse()

	cacheHandler := pkgs.CacheHandler{
		Dir:      *cacheDir,
		CacheDay: *cacheDay,
	}
	cacheHandler.Init()

	serverHandler := &pkgs.ServerHandler{
		Trans:  pkgs.UrlTransformer{},
		Webs:   pkgs.Webshoter{},
		Cacher: cacheHandler,
	}

	log.Fatal(http.ListenAndServe(":8080", serverHandler))
}
