package documentdb

import "github.com/jen20/riviera/azure"

type GetDatabaseResponse struct {
	ID               *string            `mapstructure:"id"`
	Name             *string            `mapstructure:"name"`
	Type             *string            `mapstructure:"type"`
	FailoverPolicies []FailoverPolicies `mapstructure:"failoverPolicies"`
}

type GetDocumentDBDatabase struct {
	DatabaseName      string `json:"-"`
	ResourceGroupName string `json:"-"`
}

func (s GetDocumentDBDatabase) APIInfo() azure.APIInfo {
	return azure.APIInfo{
		APIVersion:  apiVersion,
		Method:      "GET",
		URLPathFunc: documentDBPath(s.ResourceGroupName, s.DatabaseName),
		ResponseTypeFunc: func() interface{} {
			return &GetDatabaseResponse{}
		},
	}
}
