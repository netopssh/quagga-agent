package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// BGPSummaryHandle
func (ctrl *Controller) BGPSummaryHandle(c *echo.Context) error {
	return c.JSON(http.StatusOK, ctrl.Cache.BGPSummary)
}

// BGPPeerHandle
func (ctrl *Controller) BGPPeerHandle(c *echo.Context) error {
	return c.JSON(http.StatusOK, ctrl.Cache.BGPSummary.Peers[c.Param("peer")])
}

// BGPPeerStateHandle
func (ctrl *Controller) BGPPeerStateHandle(c *echo.Context) error {
	return c.JSON(http.StatusOK, ctrl.Cache.BGPSummary.Peers[c.Param("peer")].State)
}
