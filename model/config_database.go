package model

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

func (s *SqlSettings) SetDefaults() {
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

func dbConfigParser(buff []byte) (interface{}, error) {
	settings = &SqlSettings{}
	settings.SetDefault()
	return settings, nil
}

func GetDBSettings() *SqlSettings {
	settings := GetSettings(DATABASE_CONFIG_NAME)
	if settings != nil {
		tmp := settings.(*SqlSettings)
		return tmp
	}
	return nil
}

func init() {
	_, err := AddConfigEntry(DATABASE_CONFIG_NAME, DATABASE_CONFIG_FILE_PATH, true, dbConfigParser)
	if err != nil {
		l4g.Error(fmt.Sprintf("%v ", err))
	}
}
