package pkgs

import (
	"regexp"
	"strings"
)

type UrlTransformer struct {
}

type OptionsMapItem struct {
	reg string
}

func (t *UrlTransformer) Handle(fullUrl string) (map[string]string, string) {
	if strings.Trim(fullUrl, "/") == "" {
		panic("url 为 空")
	}

	reg := regexp.MustCompile(`^(/[\w,]+)?/(.+)$`)
	if !reg.MatchString(fullUrl) {
		panic("url不合法")
	}

	options := t.TransOptions(reg.FindStringSubmatch(fullUrl)[1])
	url := reg.FindStringSubmatch(fullUrl)[2]

	return options, url
}

func (t UrlTransformer) TransOptions(optionsStr string) map[string]string {
	options := map[string]string{
		"width":   "1920",
		"height":  "1080",
		"quality": "90",
	}

	optionsMaps := map[string]OptionsMapItem{
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
