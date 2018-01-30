package sqlstore

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/mattermost/gorp"

	"github.com/primefour/xserver/store"
)

/*
 *this is unique stucture for database manage
 */
type SqlStore interface {
	DriverName() string              //get current driver name
	GetCurrentSchemaVersion() string //get database schemea version
	GetMaster() *gorp.DbMap          //connect to master
	GetSearchReplica() *gorp.DbMap
	GetReplica() *gorp.DbMap
	TotalMasterDbConnections() int
	TotalReadDbConnections() int
	TotalSearchDbConnections() int
	MarkSystemRanUnitTests()
	DoesTableExist(tablename string) bool
	DoesColumnExist(tableName string, columName string) bool
	CreateColumnIfNotExists(tableName string, columnName string, mySqlColType string, postgresColType string, defaultValue string) bool
	RemoveColumnIfExists(tableName string, columnName string) bool
	RemoveTableIfExists(tableName string) bool
	RenameColumnIfExists(tableName string, oldColumnName string, newColumnName string, colType string) bool
	GetMaxLengthOfColumnIfExists(tableName string, columnName string) string
	AlterColumnTypeIfExists(tableName string, columnName string, mySqlColType string, postgresColType string) bool
	CreateUniqueIndexIfNotExists(indexName string, tableName string, columnName string) bool
	CreateIndexIfNotExists(indexName string, tableName string, columnName string) bool
	CreateCompositeIndexIfNotExists(indexName string, tableName string, columnNames []string) bool
	CreateFullTextIndexIfNotExists(indexName string, tableName string, columnName string) bool
	RemoveIndexIfExists(indexName string, tableName string) bool
	GetAllConns() []*gorp.DbMap //get all database instance
	Close()
	Team() store.TeamStore       //get team database interface
	Channel() store.ChannelStore //get channel database interface
	Post() store.PostStore       //get post databawse interface
	User() store.UserStore
	Audit() store.AuditStore
	ClusterDiscovery() store.ClusterDiscoveryStore
	Compliance() store.ComplianceStore
	Session() store.SessionStore
	OAuth() store.OAuthStore
	System() store.SystemStore
	Webhook() store.WebhookStore
	Command() store.CommandStore
	CommandWebhook() store.CommandWebhookStore
	Preference() store.PreferenceStore
	License() store.LicenseStore
	Token() store.TokenStore
	Emoji() store.EmojiStore
	Status() store.StatusStore
	FileInfo() store.FileInfoStore
	Reaction() store.ReactionStore
	Job() store.JobStore
	Plugin() store.PluginStore
	UserAccessToken() store.UserAccessTokenStore
}
