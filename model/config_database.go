package model

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/utils"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

const (
	DATABASE_DRIVER_MYSQL            = "mysql"
	DATABASE_DRIVER_POSTGRES         = "postgres"
	SQL_SETTINGS_DEFAULT_DATA_SOURCE = "xserver_user:xserver_password@tcp(dockerhost:3306)/xserver_dev_database?charset=utf8mb4,utf8&readTimeout=30s&writeTimeout=30s"
	DATABASE_CONFIG_FILE_PATH        = "./config/database_config.json"
	DATABASE_CONFIG_NAME             = "DATABASE_SETTINGS"
)

type SqlSettings struct {
	DriverName               *string
	DataSource               *string
	DataSourceReplicas       []string
	DataSourceSearchReplicas []string
	MaxIdleConns             *int
	MaxOpenConns             *int
	Trace                    bool
	AtRestEncryptKey         string
	QueryTimeout             *int
}

func (ss *SqlSettings) isValid() *AppError {
	if len(ss.AtRestEncryptKey) < 32 {
		return NewAppError("Config.IsValid", "model.config.is_valid.encrypt_sql.app_error", nil, "", http.StatusBadRequest)
	}

	if !(*ss.DriverName == DATABASE_DRIVER_MYSQL || *ss.DriverName == DATABASE_DRIVER_POSTGRES) {
		return NewAppError("Config.IsValid", "model.config.is_valid.sql_driver.app_error", nil, "", http.StatusBadRequest)
	}

	if *ss.MaxIdleConns <= 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.sql_idle.app_error", nil, "", http.StatusBadRequest)
	}

	if *ss.QueryTimeout <= 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.sql_query_timeout.app_error", nil, "", http.StatusBadRequest)
	}

	if len(*ss.DataSource) == 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.sql_data_src.app_error", nil, "", http.StatusBadRequest)
	}

	if *ss.MaxOpenConns <= 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.sql_max_conn.app_error", nil, "", http.StatusBadRequest)
	}

	return nil
}

func (s *SqlSettings) setDefaults() {
	if s.DriverName == nil {
		s.DriverName = NewString(DATABASE_DRIVER_MYSQL)
	}

	if s.DataSource == nil {
		s.DataSource = NewString(SQL_SETTINGS_DEFAULT_DATA_SOURCE)
	}

	if len(s.AtRestEncryptKey) == 0 {
		s.AtRestEncryptKey = NewRandomString(32)
	}

	if s.MaxIdleConns == nil {
		s.MaxIdleConns = NewInt(20)
	}

	if s.MaxOpenConns == nil {
		s.MaxOpenConns = NewInt(300)
	}

	if s.QueryTimeout == nil {
		s.QueryTimeout = NewInt(30)
	}
}

func dbConfigParser(f *os.File) (interface{}, error) {
	settings := &SqlSettings{}
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(f); err != nil {
		return nil, err
	}
	unmarshalErr := v.Unmarshal(settings)
	settings.setDefaults()
	l4g.Debug("database settings is DriverName:%s  ", *(settings.DriverName))
	l4g.Debug("database settings is DataSource:%s  ", *(settings.DataSource))
	l4g.Debug("database settings is:%v  ", *settings)
	return settings, unmarshalErr
}

func GetDBSettings() *SqlSettings {
	settings := utils.GetSettings(DATABASE_CONFIG_NAME)
	if settings != nil {
		tmp := settings.(*SqlSettings)
		return tmp
	}
	return nil
}

func init() {
	_, err := utils.AddConfigEntry(DATABASE_CONFIG_NAME, DATABASE_CONFIG_FILE_PATH, true, dbConfigParser)
	if err != nil {
		return
	}
}
