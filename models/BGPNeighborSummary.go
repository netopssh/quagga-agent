package models

// BGPNeighborSummary
type BGPNeighborSummary struct {
	RemoteAs              int    `json:"remoteAs"`
	Version               int    `json:"version"`
	MsgRcvd               int    `json:"msgRcvd"`
	MsgSent               int    `json:"msgSent"`
	TableVersion          int    `json:"tableVersion"`
	Outq                  int    `json:"outq"`
	Inq                   int    `json:"inq"`
	Uptime                string `json:"uptime"`
	PrefixReceivedCount   int    `json:"prefixReceivedCount"`
	PrefixAdvertisedCount int    `json:"prefixAdvertisedCount"`
	State                 string `json:"state"`
	IDType                string `json:"idType"`
}
