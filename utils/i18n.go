package utils

import (
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

// "zh-CN"
func tfuncWithFallback(pref string) i18n.TranslateFunc {
	//try prefer language to translate
	t, _ := i18n.Tfunc(pref)

	return func(translationID string, args ...interface{}) string {
		if translated := t(translationID, args...); translated != translationID {
			return translated
		}
		//don't support prefer language,use default
		l4g.Warn("don't support prefer language %s ", pref)
		t, _ := i18n.Tfunc(DEFAULT_LOCALE)
		return t(translationID, args...)
	}
}

func InitTranslationsWithDir(dir string) (locales map[string]string) {
	i18nDirectory := FindDir(dir)
	files, _ := ioutil.ReadDir(i18nDirectory)
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".json" {
			filename := f.Name()
			locales[strings.Split(filename, ".")[0]] = i18nDirectory + filename
			i18n.MustLoadTranslationFile(i18nDirectory + filename)
		}
	}
	return locales
}

func GetUserTranslations(locale string, locales map[string]string) i18n.TranslateFunc {
	if _, ok := locales[locale]; !ok {
		l4g.Warn("don't support locale %s ", locale)
		locale = DEFAULT_LOCALE
	}
	translations := tfuncWithFallback(locale)
	return translations
}

func GetTranslationsAndLocale(w http.ResponseWriter, r *http.Request, locales map[string]string) (i18n.TranslateFunc, string) {
	// This is for checking against locales like pt_BR or zn_CN
	headerLocaleFull := strings.Split(r.Header.Get("Accept-Language"), ",")[0]
	// This is for checking agains locales like en, es
	headerLocale := strings.Split(strings.Split(r.Header.Get("Accept-Language"), ",")[0], "-")[0]
	if locales[headerLocaleFull] != "" {
		translations := TfuncWithFallback(headerLocaleFull)
		return translations, headerLocaleFull
	} else if locales[headerLocale] != "" {
		translations := TfuncWithFallback(headerLocale)
		return translations, headerLocale
	}

	translations := TfuncWithFallback(DEFAULT_LOCALE)
	return translations, DEFAULT_LOCALE
}
