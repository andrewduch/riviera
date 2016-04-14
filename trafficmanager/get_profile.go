package trafficmanager

import "github.com/jen20/riviera/azure"

type GetTrafficManagerProfileResponse struct {
	ID                   *string     `mapstructure:"id"`
	Name                 *string     `mapstructure:"name"`
	Type                 *string     `mapstructure:"type"`
	TrafficRoutingMethod *string     `mapstructure:"trafficRoutingMethod"`
	ProfileStatus        *string     `mapstructure:"profileStatus"`
	Endpoints            []*Endpoint `mapstructure:"endpoints"`
}

type Endpoint struct {
	ID         *string             `mapstructure:"id"`
	Name       *string             `mapstructure:"name"`
	Type       *string             `mapstructure:"type"`
	Properties *EndpointProperties `mapstructure:"properties"`
}

type EndpointProperties struct {
	EndpointStatus        *string `mapstructure:"endpointStatus"`
	EndpointMonitorStatus *string `mapstructure:"endpointMonitorStatus"`
	Weight                int     `mapstructure:"weight"`
	Priority              int     `mapstructure:"priority"`
	Target                *string `mapstructure:"target"`
	EndpointLocation      *string `mapstructure:"endpointLocation"`
	MinChildEndpoints     int     `mapstructure:"minChildEndpoints"`
}

type GetTrafficManagerProfile struct {
	ProfileName       string `json:"-"`
	ResourceGroupName string `json:"-"`
}

func (s GetTrafficManagerProfile) APIInfo() azure.APIInfo {
	return azure.APIInfo{
		APIVersion:  apiVersion,
		Method:      "GET",
		URLPathFunc: trafficManagerProfileURLPath(s.ResourceGroupName, s.ProfileName),
		ResponseTypeFunc: func() interface{} {
			return &GetTrafficManagerProfileResponse{}
		},
	}
}
