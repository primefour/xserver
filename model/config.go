package model

const (
	CONN_SECURITY_NONE     = ""
	CONN_SECURITY_PLAIN    = "PLAIN"
	CONN_SECURITY_TLS      = "TLS"
	CONN_SECURITY_STARTTLS = "STARTTLS"

	IMAGE_DRIVER_LOCAL = "local"
	IMAGE_DRIVER_S3    = "amazons3"

	MINIO_ACCESS_KEY = "minioaccesskey"
	MINIO_SECRET_KEY = "miniosecretkey"
	MINIO_BUCKET     = "mattermost-test"

	PASSWORD_MAXIMUM_LENGTH = 64
	PASSWORD_MINIMUM_LENGTH = 5

	WEBSERVER_MODE_REGULAR  = "regular"
	WEBSERVER_MODE_GZIP     = "gzip"
	WEBSERVER_MODE_DISABLED = "disabled"

	FAKE_SETTING = "********************************"

	SITENAME_MAX_LENGTH = 30
)
