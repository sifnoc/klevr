package agent

type Cluster struct{
	Primary Primary `json:"primary"`
	Secondary []Secondary `json:"secondary"`
}

type Primary struct{
	AgentKey       string `json:"agentKey"`
	IP             string `json:"ip"`
	Port           int    `json:"port"`
	IsActive       bool   `json:"isActive"`
	LastAccessTime int64  `json:"lastAccessTime"`
}
type Secondary struct{
	IP             string `json:"ip"`
	Port           int    `json:"port"`
	IsActive       bool   `json:"isActive"`
}

type AliveCheck struct{
	Time           int64  `json:"time"`
	IsActive       bool   `json:"isActive"`
}