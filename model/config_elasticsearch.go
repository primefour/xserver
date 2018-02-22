package model

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/utils"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
)

const (
	ELASTICSEARCH_SETTINGS_DEFAULT_CONNECTION_URL                    = ""
	ELASTICSEARCH_SETTINGS_DEFAULT_USERNAME                          = ""
	ELASTICSEARCH_SETTINGS_DEFAULT_PASSWORD                          = ""
	ELASTICSEARCH_SETTINGS_DEFAULT_POST_INDEX_REPLICAS               = 1
	ELASTICSEARCH_SETTINGS_DEFAULT_POST_INDEX_SHARDS                 = 1
	ELASTICSEARCH_SETTINGS_DEFAULT_AGGREGATE_POSTS_AFTER_DAYS        = 365
	ELASTICSEARCH_SETTINGS_DEFAULT_POSTS_AGGREGATOR_JOB_START_TIME   = "03:00"
	ELASTICSEARCH_SETTINGS_DEFAULT_INDEX_PREFIX                      = ""
	ELASTICSEARCH_SETTINGS_DEFAULT_LIVE_INDEXING_BATCH_SIZE          = 1
	ELASTICSEARCH_SETTINGS_DEFAULT_BULK_INDEXING_TIME_WINDOW_SECONDS = 3600
	ELASTICSEARCH_SETTINGS_DEFAULT_REQUEST_TIMEOUT_SECONDS           = 30
	ES_CONFIG_FILE_PATH                                              = "./config/es_config.json"
	ES_CONFIG_NAME                                                   = "ES_SETTINGS"
)

type ElasticsearchSettings struct {
	ConnectionUrl                 *string
	Username                      *string
	Password                      *string
	EnableIndexing                *bool
	EnableSearching               *bool
	Sniff                         *bool
	PostIndexReplicas             *int
	PostIndexShards               *int
	AggregatePostsAfterDays       *int
	PostsAggregatorJobStartTime   *string
	IndexPrefix                   *string
	LiveIndexingBatchSize         *int
	BulkIndexingTimeWindowSeconds *int
	RequestTimeoutSeconds         *int
}

func (s *ElasticsearchSettings) SetDefaults() {
	if s.ConnectionUrl == nil {
		s.ConnectionUrl = NewString(ELASTICSEARCH_SETTINGS_DEFAULT_CONNECTION_URL)
	}

	if s.Username == nil {
		s.Username = NewString(ELASTICSEARCH_SETTINGS_DEFAULT_USERNAME)
	}

	if s.Password == nil {
		s.Password = NewString(ELASTICSEARCH_SETTINGS_DEFAULT_PASSWORD)
	}

	if s.EnableIndexing == nil {
		s.EnableIndexing = NewBool(false)
	}

	if s.EnableSearching == nil {
		s.EnableSearching = NewBool(false)
	}

	if s.Sniff == nil {
		s.Sniff = NewBool(true)
	}

	if s.PostIndexReplicas == nil {
		s.PostIndexReplicas = NewInt(ELASTICSEARCH_SETTINGS_DEFAULT_POST_INDEX_REPLICAS)
	}

	if s.PostIndexShards == nil {
		s.PostIndexShards = NewInt(ELASTICSEARCH_SETTINGS_DEFAULT_POST_INDEX_SHARDS)
	}

	if s.AggregatePostsAfterDays == nil {
		s.AggregatePostsAfterDays = NewInt(ELASTICSEARCH_SETTINGS_DEFAULT_AGGREGATE_POSTS_AFTER_DAYS)
	}

	if s.PostsAggregatorJobStartTime == nil {
		s.PostsAggregatorJobStartTime = NewString(ELASTICSEARCH_SETTINGS_DEFAULT_POSTS_AGGREGATOR_JOB_START_TIME)
	}

	if s.IndexPrefix == nil {
		s.IndexPrefix = NewString(ELASTICSEARCH_SETTINGS_DEFAULT_INDEX_PREFIX)
	}

	if s.LiveIndexingBatchSize == nil {
		s.LiveIndexingBatchSize = NewInt(ELASTICSEARCH_SETTINGS_DEFAULT_LIVE_INDEXING_BATCH_SIZE)
	}

	if s.BulkIndexingTimeWindowSeconds == nil {
		s.BulkIndexingTimeWindowSeconds = NewInt(ELASTICSEARCH_SETTINGS_DEFAULT_BULK_INDEXING_TIME_WINDOW_SECONDS)
	}

	if s.RequestTimeoutSeconds == nil {
		s.RequestTimeoutSeconds = NewInt(ELASTICSEARCH_SETTINGS_DEFAULT_REQUEST_TIMEOUT_SECONDS)
	}
}

func (ess *ElasticsearchSettings) isValid() *AppError {
	if *ess.EnableIndexing {
		if len(*ess.ConnectionUrl) == 0 {
			return NewAppError("Config.IsValid", "model.config.is_valid.elastic_search.connection_url.app_error", nil, "", http.StatusBadRequest)
		}
	}

	if *ess.EnableSearching && !*ess.EnableIndexing {
		return NewAppError("Config.IsValid", "model.config.is_valid.elastic_search.enable_searching.app_error", nil, "", http.StatusBadRequest)
	}

	if *ess.AggregatePostsAfterDays < 1 {
		return NewAppError("Config.IsValid", "model.config.is_valid.elastic_search.aggregate_posts_after_days.app_error", nil, "", http.StatusBadRequest)
	}

	if _, err := time.Parse("15:04", *ess.PostsAggregatorJobStartTime); err != nil {
		return NewAppError("Config.IsValid", "model.config.is_valid.elastic_search.posts_aggregator_job_start_time.app_error", nil, err.Error(), http.StatusBadRequest)
	}

	if *ess.LiveIndexingBatchSize < 1 {
		return NewAppError("Config.IsValid", "model.config.is_valid.elastic_search.live_indexing_batch_size.app_error", nil, "", http.StatusBadRequest)
	}

	if *ess.BulkIndexingTimeWindowSeconds < 1 {
		return NewAppError("Config.IsValid", "model.config.is_valid.elastic_search.bulk_indexing_time_window_seconds.app_error", nil, "", http.StatusBadRequest)
	}

	if *ess.RequestTimeoutSeconds < 1 {
		return NewAppError("Config.IsValid", "model.config.is_valid.elastic_search.request_timeout_seconds.app_error", nil, "", http.StatusBadRequest)
	}

	return nil
}

func (ess *ElasticsearchSettings) Sanitize() {
	*ess.Password = FAKE_SETTING
}

func esConfigParser(f *os.File) (interface{}, error) {
	settings := &ElasticsearchSettings{}
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(f); err != nil {
		return nil, err
	}
	unmarshalErr := v.Unmarshal(settings)
	settings.SetDefaults()
	l4g.Debug("ES settings is:%v  ", *settings)
	return settings, unmarshalErr
}

func GetESSettings() *ElasticsearchSettings {
	settings := utils.GetSettings(ES_CONFIG_NAME)
	if settings != nil {
		tmp := settings.(*ElasticsearchSettings)
		return tmp
	}
	return nil
}

func init() {
	_, err := utils.AddConfigEntry(ES_CONFIG_NAME, ES_CONFIG_FILE_PATH, true, esConfigParser)
	if err != nil {
		return
	}
}
