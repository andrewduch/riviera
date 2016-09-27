package documentdb

import "github.com/jen20/riviera/azure"

type FailoverDBDatabase struct {
	DatabaseName      string             `json:"-"`
	ResourceGroupName string             `json:"-"`
	FailoverPolicies  []FailoverPolicies `json:"failoverPolicies,omitempty"`
}

type FailoverPolicies struct {
	LocationName     string `json:"locationName"`
	FailoverPriority int    `json:"failoverPriority"`
}

func (s FailoverDBDatabase) APIInfo() azure.APIInfo {
	return azure.APIInfo{
		APIVersion:  apiVersion,
		Method:      "POST",
		URLPathFunc: documentDBFailoverPath(s.ResourceGroupName, s.DatabaseName),
		ResponseTypeFunc: func() interface{} {
			return &GetDatabaseResponse{}
		},
	}
}
