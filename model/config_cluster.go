package model

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

func (s *ClusterSettings) SetDefaults() {
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

type MetricsSettings struct {
	Enable           *bool
	BlockProfileRate *int
	ListenAddress    *string
}

func (s *MetricsSettings) SetDefaults() {
	if s.ListenAddress == nil {
		s.ListenAddress = NewString(":8067")
	}

	if s.Enable == nil {
		s.Enable = NewBool(false)
	}

	if s.BlockProfileRate == nil {
		s.BlockProfileRate = NewInt(0)
	}
}
