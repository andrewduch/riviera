package sql

import "github.com/jen20/riviera/azure"

type GetReplicationLinkResponse struct {
	ID              *string            `mapstructure:"id"`
	Name            *string            `mapstructure:"name"`
	Location        *string            `mapstructure:"location"`
	Tags            *map[string]string `mapstructure:"tags"`
	PartnerServer   *string            `mapstructure:"partnerServer"`
	PartnerDatabase *string            `mapstructure:"partnerDatabase"`
	Role            *string            `mapstructure:"role"`
}

type GetReplicationLink struct {
	DatabaseName      string `json:"-"`
	ServerName        string `json:"-"`
	ResourceGroupName string `json:"-"`
	LinkID            string `json:"-"`
}

func (s GetReplicationLink) APIInfo() azure.APIInfo {
	return azure.APIInfo{
		APIVersion:  apiVersion,
		Method:      "GET",
		URLPathFunc: sqlDatabaseReplicationLinkURLPath(s.ResourceGroupName, s.ServerName, s.DatabaseName, s.LinkID),
		ResponseTypeFunc: func() interface{} {
			return &GetReplicationLinkResponse{}
		},
	}
}
