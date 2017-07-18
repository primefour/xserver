// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package utils

import (
	"os"
	"testing"
)

func SaveConfig(fileName string, config interface{}) *AppError {
	/*
		cfgMutex.Lock()
		defer cfgMutex.Unlock()

		b, err := json.MarshalIndent(config, "", "    ")
		if err != nil {
			return model.NewLocAppError("SaveConfig", "utils.config.save_config.saving.app_error",
				map[string]interface{}{"Filename": fileName}, err.Error())
		}

		err = ioutil.WriteFile(fileName, b, 0644)
		if err != nil {
			return model.NewLocAppError("SaveConfig", "utils.config.save_config.saving.app_error",
				map[string]interface{}{"Filename": fileName}, err.Error())
		}
	*/
	return nil
}

func LoadConfig(fileName string) interface{} {
	self.cfgMutex.Lock()
	defer self.cfgMutex.Unlock()

	fileNameWithExtension := filepath.Base(fileName)
	fileExtension := filepath.Ext(fileNameWithExtension)
	fileDir := filepath.Dir(fileName)

	if len(fileNameWithExtension) > 0 {
		fileNameOnly := fileNameWithExtension[:len(fileNameWithExtension)-len(fileExtension)]
		viper.SetConfigName(fileNameOnly)
	} else {
		viper.SetConfigName("config")
	}

	if len(fileDir) > 0 {
		viper.AddConfigPath(fileDir)
	}

	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")
	viper.AddConfigPath(".")

	configReadErr := viper.ReadInConfig()
	if configReadErr != nil {
		errMsg := T("utils.config.load_config.opening.panic", map[string]interface{}{"Filename": fileName, "Error": configReadErr.Error()})
		l4g.Error("%s ", errMsg)
		os.Exit(1)
	}

	var config model.Config
	unmarshalErr := viper.Unmarshal(&config)
	if unmarshalErr != nil {
		errMsg := T("utils.config.load_config.decoding.panic", map[string]interface{}{"Filename": fileName, "Error": unmarshalErr.Error()})
		l4g.Error("%s", errMsg)
		os.Exit(1)
	}

	CfgFileName = viper.ConfigFileUsed()

	l4g.Debug("use config file is %s ", CfgFileName)

	needSave := len(config.SqlSettings.AtRestEncryptKey) == 0 || len(*config.FileSettings.PublicLinkSalt) == 0 ||
		len(config.EmailSettings.InviteSalt) == 0

	config.SetDefaults()

	if err := config.IsValid(); err != nil {
		l4g.Debug("panic configure file fail")
		panic(T(err.Id))
	}

	if needSave {
		cfgMutex.Unlock()
		if err := SaveConfig(CfgFileName, &config); err != nil {
			l4g.Warn(T(err.Id))
		}
		cfgMutex.Lock()
	}

	if err := ValidateLocales(&config); err != nil {
		l4g.Debug("panic configure file fail")
		panic(T(err.Id))
	}

	configureLog(&config.LogSettings)

	if config.FileSettings.DriverName == model.IMAGE_DRIVER_LOCAL {
		dir := config.FileSettings.Directory
		if len(dir) > 0 && dir[len(dir)-1:] != "/" {
			config.FileSettings.Directory += "/"
		}
	}

	Cfg = &config
	CfgHash = fmt.Sprintf("%x", md5.Sum([]byte(Cfg.ToJson())))
	ClientCfg = getClientConfig(Cfg)
	clientCfgJson, _ := json.Marshal(ClientCfg)
	ClientCfgHash = fmt.Sprintf("%x", md5.Sum(clientCfgJson))

	SetSiteURL(*Cfg.ServiceSettings.SiteURL)
}

/*
	if configReadErr := viper.ReadInConfig(); configReadErr == nil {
		xj, err := self.Intf.LoadConfig(self)
		if err == nil {
			self.Intf.UpdateConfgi(xj)
		} else {
			l4g.Error(fmt.Sprintf("Failed to read while watching config file"))
		}
	} else {
		l4g.Error(fmt.Sprintf("Failed to read while watching config file at %v with err=%v", cfn, configReadErr.Error()))
	}
*/

func TestConfig(t *testing.T) {
	TranslationsPreInit()
	LoadConfig("config.json")
	InitTranslations(Cfg.LocalizationSettings)
}

func TestConfigFromEnviroVars(t *testing.T) {

	os.Setenv("MM_TEAMSETTINGS_SITENAME", "From Enviroment")
	os.Setenv("MM_TEAMSETTINGS_CUSTOMBRANDTEXT", "Custom Brand")
	os.Setenv("MM_SERVICESETTINGS_ENABLECOMMANDS", "false")
	os.Setenv("MM_SERVICESETTINGS_READTIMEOUT", "400")

	TranslationsPreInit()
	EnableConfigFromEnviromentVars()
	LoadConfig("config.json")

	if Cfg.TeamSettings.SiteName != "From Enviroment" {
		t.Fatal("Couldn't read config from enviroment var")
	}

	if *Cfg.TeamSettings.CustomBrandText != "Custom Brand" {
		t.Fatal("Couldn't read config from enviroment var")
	}

	if *Cfg.ServiceSettings.EnableCommands != false {
		t.Fatal("Couldn't read config from enviroment var")
	}

	if *Cfg.ServiceSettings.ReadTimeout != 400 {
		t.Fatal("Couldn't read config from enviroment var")
	}

	os.Unsetenv("MM_TEAMSETTINGS_SITENAME")
	os.Unsetenv("MM_TEAMSETTINGS_CUSTOMBRANDTEXT")
	os.Unsetenv("MM_SERVICESETTINGS_ENABLECOMMANDS")
	os.Unsetenv("MM_SERVICESETTINGS_READTIMEOUT")

	Cfg.TeamSettings.SiteName = "Mattermost"
	*Cfg.ServiceSettings.SiteURL = ""
	*Cfg.ServiceSettings.EnableCommands = true
	*Cfg.ServiceSettings.ReadTimeout = 300
	SaveConfig(CfgFileName, Cfg)

	LoadConfig("config.json")

	if Cfg.TeamSettings.SiteName != "Mattermost" {
		t.Fatal("should have been reset")
	}

}
