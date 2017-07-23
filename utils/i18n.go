package utils

import (
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/nicksnyder/go-i18n/i18n"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

const (
	DEFAULT_LOCALE = "zh-CN"
)

var T i18n.TranslateFunc
var locales map[string]string = make(map[string]string)

// "zh-CN"
func tfuncWithFallback(pref string) i18n.TranslateFunc {
	//try prefer language to translate
	T, _ = i18n.Tfunc(pref)
	return func(translationID string, args ...interface{}) string {
		if translated := T(translationID, args...); translated != translationID {
			return translated
		}
		//don't support prefer language,use default
		l4g.Warn("don't support prefer language %s ", pref)
		t, _ := i18n.Tfunc(DEFAULT_LOCALE)
		return t(translationID, args...)
	}
}

func InitTranslationsWithDir(dir string) {
	i18nDirectory := FindDir(dir)
	files, _ := ioutil.ReadDir(i18nDirectory)
	for _, f := range files {
		l4g.Info(fmt.Sprintf("i18n f %s", f.Name()))
		if filepath.Ext(f.Name()) == ".json" {
			filename := f.Name()
			locales[strings.Split(filename, ".")[0]] = i18nDirectory + filename
			i18n.MustLoadTranslationFile(i18nDirectory + filename)
		}
	}

	GetUserTranslations(DEFAULT_LOCALE)
}

func GetUserTranslations(locale string) i18n.TranslateFunc {
	if _, ok := locales[locale]; !ok {
		l4g.Warn("don't support locale %s ", locale)
		locale = DEFAULT_LOCALE
	}
	translations := tfuncWithFallback(locale)
	return translations
}

func GetTranslationsAndLocale(w http.ResponseWriter, r *http.Request) (i18n.TranslateFunc, string) {
	// This is for checking against locales like pt_BR or zn_CN
	headerLocaleFull := strings.Split(r.Header.Get("Accept-Language"), ",")[0]
	// This is for checking agains locales like en, es
	headerLocale := strings.Split(strings.Split(r.Header.Get("Accept-Language"), ",")[0], "-")[0]
	if locales[headerLocaleFull] != "" {
		translations := tfuncWithFallback(headerLocaleFull)
		return translations, headerLocaleFull
	} else if locales[headerLocale] != "" {
		translations := tfuncWithFallback(headerLocale)
		return translations, headerLocale
	}

	translations := tfuncWithFallback(DEFAULT_LOCALE)
	return translations, DEFAULT_LOCALE
}
