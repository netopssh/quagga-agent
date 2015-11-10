package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/netopssh/quagga-agent/models"
	"github.com/netopssh/quagga-agent/scraper"
)

const BGPSummaryCommand = "cl-bgp summary show json"
const BGPRouteCommand = "cl-bgp route show json"

// Cache
type Cache struct {
	BGPSummary models.BGPSummary
}

// New
func New() *Cache {
	c := &Cache{}
	c.PopulateCache()

	return c
}

// Run
func (c *Cache) Run() {
	go func() {
		for {
			time.Sleep(30 * time.Second)
			c.PopulateCache()
		}
	}()
}

// PopulateCache
func (c *Cache) PopulateCache() {
	b := scraper.GetCommandOutput(BGPSummaryCommand)

	err := json.Unmarshal(b, &c.BGPSummary)
	if err != nil {
		fmt.Println("error:", err)
	}

}
