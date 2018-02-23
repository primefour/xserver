package model

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/primefour/xserver/utils"
	"github.com/spf13/viper"
	"os"
)

const (
	//cluster config
	CLUSTER_CONFIG_FILE_PATH = "./config/cluster_config.json"
	CLUSTER_CONFIG_NAME      = "CLUSTER_SETTINGS"
)

type ClusterSettings struct {
	Enable                *bool
	ClusterName           *string
	OverrideHostname      *string
	UseIpAddress          *bool
	UseExperimentalGossip *bool
	ReadOnlyConfig        *bool
	GossipPort            *int
	StreamingPort         *int
}

func (s *ClusterSettings) setDefaults() {
	if s.Enable == nil {
		s.Enable = NewBool(false)
	}

	if s.ClusterName == nil {
		s.ClusterName = NewString("")
	}

	if s.OverrideHostname == nil {
		s.OverrideHostname = NewString("")
	}

	if s.UseIpAddress == nil {
		s.UseIpAddress = NewBool(true)
	}

	if s.UseExperimentalGossip == nil {
		s.UseExperimentalGossip = NewBool(false)
	}

	if s.ReadOnlyConfig == nil {
		s.ReadOnlyConfig = NewBool(true)
	}

	if s.GossipPort == nil {
		s.GossipPort = NewInt(8074)
	}

	if s.StreamingPort == nil {
		s.StreamingPort = NewInt(8075)
	}
}

func clusterConfigParser(f *os.File) (interface{}, error) {
	settings := &ClusterSettings{}
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(f); err != nil {
		return nil, err
	}
	unmarshalErr := v.Unmarshal(settings)
	settings.setDefaults()
	l4g.Debug("cluster settings is:%v  ", *settings)
	return settings, unmarshalErr
}

func GetClusterSettings() *ClusterSettings {
	settings := utils.GetSettings(CLUSTER_CONFIG_NAME)
	if settings != nil {
		tmp := settings.(*ClusterSettings)
		return tmp
	}
	return nil
}

func init() {
	_, err := utils.AddConfigEntry(CLUSTER_CONFIG_NAME, CLUSTER_CONFIG_FILE_PATH, true, clusterConfigParser)
	if err != nil {
		return
	}
}
