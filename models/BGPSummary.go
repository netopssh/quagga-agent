package models

// BGPSummary
type BGPSummary struct {
	RouterID     string                        `json:"routerId"`
	AS           int                           `json:"as"`
	TableVersion int                           `json:"tableVersion"`
	RibCount     int                           `json:"ribCount"`
	RibMemory    int                           `json:"ribMemory"`
	PeerCount    int                           `json:"peerCount"`
	PeerMemory   int                           `json:"peerMemory"`
	Peers        map[string]BGPNeighborSummary `json:"peers"`
	TotalPeers   int                           `json:"totalPeers"`
	DynamicPeers int                           `json:"dynamicPeers"`
}
