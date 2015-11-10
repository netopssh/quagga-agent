package main

import (
	"github.com/labstack/echo"
	"github.com/netopssh/quagga-agent/cache"
)

var ctrl *Controller

func main() {

	c := cache.New()
	ctrl = New(c)
	c.Run()

	e := echo.New()

	// Routes
	e.Get("/bgp/summary", ctrl.BGPSummaryHandle)
	e.Get("/bgp/:peer", ctrl.BGPPeerHandle)
	e.Get("/bgp/:peer/state", ctrl.BGPPeerStateHandle)

	// Start server
	e.Run(":3211")
}
