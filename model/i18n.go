package model

import (
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/nicksnyder/go-i18n/i18n"
	"github.com/primefour/xserver/utils"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

const (
	DEFAULT_LOCALE = "zh-CN"
)

var localeDirName string = "i18n"
var locateDirPath string
var T i18n.TranslateFunc
var locales map[string]string = make(map[string]string)

// "zh-CN"
func tfuncWithFallback(pref string) i18n.TranslateFunc {
	//try prefer language to translate
	l4g.Debug("pref language is  %s ", pref)
	xt, err := i18n.Tfunc(pref)

	if err != nil {
		l4g.Error("get pref language translate failed %v ", err)
	}

	return func(translationID string, args ...interface{}) string {
		if translated := xt(translationID, args...); translated != translationID {
			return translated
		} else {
			//don't support prefer language,use default
			l4g.Warn("don't support prefer language %s ", pref)

			if T != nil {
				return T(translationID, args...)
			} else {
				t, _ := i18n.Tfunc(DEFAULT_LOCALE)
				return t(translationID, args...)
			}
		}
	}
}

func InitTranslations() {
	localeDir := utils.FindDir(localeDirName)
	locateDirPath, _ = filepath.Abs(localeDir)
	l4g.Debug(fmt.Sprintf("locale dir path is %s ", locateDirPath))
	initTranslationsWithDir(locateDirPath)
}

func initTranslationsWithDir(dir string) {
	dir += "/"
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		l4g.Info(fmt.Sprintf("i18n f %s", f.Name()))
		if filepath.Ext(f.Name()) == ".json" {
			filename := f.Name()
			locales[strings.Split(filename, ".")[0]] = dir + filename
			i18n.MustLoadTranslationFile(dir + filename)
		}
	}
	locale := GetAppBaseSettings().LocalizationSettings.DefaultServerLocale
	if locale == nil || len(*locale) == 0 {
		T = GetUserTranslations(DEFAULT_LOCALE)
	} else {
		T = GetUserTranslations(*locale)
	}
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

func init() {
	InitTranslations()
}
