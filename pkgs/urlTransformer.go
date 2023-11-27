package pkgs

import (
	"regexp"
	"strings"
)

type UrlTransformer struct {
}

type optionsMapItem struct {
	reg string
}

func (t *UrlTransformer) Handle(fullUrl string) (map[string]string, string) {
	if strings.Trim(fullUrl, "/") == "" {
		panic("url 为 空")
	}

	reg := regexp.MustCompile(`^(/[\w,]+)?/(https?://?.+)$`)
	if !reg.MatchString(fullUrl) {
		panic("url不合法")
	}

	options := t.TransOptions(reg.FindStringSubmatch(fullUrl)[1])

	// 截取url，因代理服务器可能存在双斜线变单斜线，导致url不一致
	urlReg := regexp.MustCompile(`^(https?)://?(.+)$`)
	full := reg.FindStringSubmatch(fullUrl)[2]
	url := urlReg.FindStringSubmatch(full)[1] + "://" + urlReg.FindStringSubmatch(full)[2]

	return options, url
}

func (t UrlTransformer) TransOptions(optionsStr string) map[string]string {
	options := map[string]string{
		"width":   "1920",
		"height":  "1080",
		"quality": "90",
	}

	optionsMaps := map[string]optionsMapItem{
		"width":   {reg: `(\d+)x`},
		"height":  {reg: `x(\d+)`},
		"quality": {reg: `q_(\d+)`},
	}

	for name, item := range optionsMaps {
		reg := regexp.MustCompile(item.reg)
		if !reg.MatchString(optionsStr) {
			continue
		}
		options[name] = reg.FindStringSubmatch(optionsStr)[1]
	}

	return options
}
