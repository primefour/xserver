package model

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/utils"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

const (
	//metrics config
	FILES_CONFIG_FILE_PATH = "./config/files_config.json"
	FILES_CONFIG_NAME      = "FILES_SETTINGS"
)

type FileSettings struct {
	EnableFileAttachments   *bool
	EnableMobileUpload      *bool
	EnableMobileDownload    *bool
	MaxFileSize             *int64
	DriverName              *string
	Directory               string
	EnablePublicLink        bool
	PublicLinkSalt          *string
	InitialFont             string
	AmazonS3AccessKeyId     string
	AmazonS3SecretAccessKey string
	AmazonS3Bucket          string
	AmazonS3Region          string
	AmazonS3Endpoint        string
	AmazonS3SSL             *bool
	AmazonS3SignV2          *bool
	AmazonS3SSE             *bool
	AmazonS3Trace           *bool
}

func (fs *FileSettings) isValid() *AppError {
	if *fs.MaxFileSize <= 0 {
		return NewAppError("Config.IsValid", "model.config.is_valid.max_file_size.app_error", nil, "", http.StatusBadRequest)
	}

	if !(*fs.DriverName == IMAGE_DRIVER_LOCAL || *fs.DriverName == IMAGE_DRIVER_S3) {
		return NewAppError("Config.IsValid", "model.config.is_valid.file_driver.app_error", nil, "", http.StatusBadRequest)
	}

	if len(*fs.PublicLinkSalt) < 32 {
		return NewAppError("Config.IsValid", "model.config.is_valid.file_salt.app_error", nil, "", http.StatusBadRequest)
	}

	return nil
}

func (s *FileSettings) SetDefaults() {
	if s.DriverName == nil {
		s.DriverName = NewString(IMAGE_DRIVER_LOCAL)
	}

	if s.AmazonS3Endpoint == "" {
		// Defaults to "s3.amazonaws.com"
		s.AmazonS3Endpoint = "s3.amazonaws.com"
	}

	if s.AmazonS3SSL == nil {
		s.AmazonS3SSL = NewBool(true) // Secure by default.
	}

	if s.AmazonS3SignV2 == nil {
		s.AmazonS3SignV2 = new(bool)
		// Signature v2 is not enabled by default.
	}

	if s.AmazonS3SSE == nil {
		s.AmazonS3SSE = NewBool(false) // Not Encrypted by default.
	}

	if s.AmazonS3Trace == nil {
		s.AmazonS3Trace = NewBool(false)
	}

	if s.EnableFileAttachments == nil {
		s.EnableFileAttachments = NewBool(true)
	}

	if s.EnableMobileUpload == nil {
		s.EnableMobileUpload = NewBool(true)
	}

	if s.EnableMobileDownload == nil {
		s.EnableMobileDownload = NewBool(true)
	}

	if s.MaxFileSize == nil {
		s.MaxFileSize = NewInt64(52428800) // 50 MB
	}

	if s.PublicLinkSalt == nil || len(*s.PublicLinkSalt) == 0 {
		s.PublicLinkSalt = NewString(NewRandomString(32))
	}

	if s.InitialFont == "" {
		// Defaults to "luximbi.ttf"
		s.InitialFont = "luximbi.ttf"
	}

	if s.Directory == "" {
		s.Directory = "./data/"
	}
}

func filesConfigParser(f *os.File) (interface{}, error) {
	settings := &FileSettings{}
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(f); err != nil {
		return nil, err
	}
	unmarshalErr := v.Unmarshal(settings)
	settings.SetDefaults()
	l4g.Debug("files settings is:%v  ", *settings)
	return settings, unmarshalErr
}

func GetfilesSettings() *FileSettings {
	settings := utils.GetSettings(FILES_CONFIG_NAME)
	if settings != nil {
		tmp := settings.(*FileSettings)
		return tmp
	}
	return nil
}

func init() {
	_, err := utils.AddConfigEntry(FILES_CONFIG_NAME, FILES_CONFIG_FILE_PATH, true, filesConfigParser)
	if err != nil {
		return
	}
}
