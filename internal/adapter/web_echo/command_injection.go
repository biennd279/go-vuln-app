package web_echo

import (
	"gitea.homelab.d3s34.me/d3s34/vuln-app/internal/core/port"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CommandInjectionAdapter struct {
	service port.CommandInjectionService
}

func NewCommandInjectionAdapter(service port.CommandInjectionService) *CommandInjectionAdapter {
	return &CommandInjectionAdapter{
		service: service,
	}
}

func (a *CommandInjectionAdapter) RegisterRoutes(e *echo.Echo) {
	e.GET("/vulnerable-command", a.vulnerableCommand)
	e.GET("/vulnerable-context-command", a.vulnerableContextCommand)
	e.GET("/vulnerable-command-with-shell", a.vulnerableCommandWithShell)
	e.GET("/non-vulnerable-command", a.nonVulnerableCommand)
	e.GET("/vulnerable-command-via-struct", a.vulnerableCommandViaStruct)
	e.GET("/vulnerable-command-via-interface", a.vulnerableCommandViaInterface)
	e.GET("/vulnerable-command-taint", a.vulnerableCommandTaint)
	e.GET("/vulnerable-command-taint-via-struct", a.vulnerableCommandTaintViaStruct)
}

func (a *CommandInjectionAdapter) vulnerableCommand(c echo.Context) error {
	command := c.QueryParam("command")
	params := c.QueryParams()["params"]

	out, _ := a.service.VulnerableCommand(command, params)
	return c.String(http.StatusOK, out)
}

func (a *CommandInjectionAdapter) vulnerableContextCommand(c echo.Context) error {
	command := c.QueryParam("command")
	params := c.QueryParams()["params"]

	out, _ := a.service.VulnerableContextCommand(c.Request().Context(), command, params)
	return c.String(http.StatusOK, out)
}

func (a *CommandInjectionAdapter) vulnerableCommandWithShell(c echo.Context) error {
	command := c.QueryParam("command")
	params := c.QueryParams()["params"]

	out, _ := a.service.VulnerableCommandWithShell(command, params)
	return c.String(http.StatusOK, out)
}

func (a *CommandInjectionAdapter) nonVulnerableCommand(c echo.Context) error {
	command := c.QueryParam("command")
	params := c.QueryParams()["params"]

	out, _ := a.service.NonVulnerableCommand(command, params)
	return c.String(http.StatusOK, out)
}

func (a *CommandInjectionAdapter) vulnerableCommandViaStruct(c echo.Context) error {
	command := c.QueryParam("command")
	params := c.QueryParams()["params"]

	out, _ := a.service.VulnerableCommandViaStruct(port.CommandStruct{
		Command: command,
		Args:    params,
	})
	return c.String(http.StatusOK, out)
}

func (a *CommandInjectionAdapter) vulnerableCommandViaInterface(c echo.Context) error {
	command := c.QueryParam("command")
	params := c.QueryParams()["params"]

	out, _ := a.service.VulnerableCommandViaInterface(&port.CommandStruct{
		Command: command,
		Args:    params,
	})
	return c.String(http.StatusOK, out)
}

func (a *CommandInjectionAdapter) vulnerableCommandTaint(c echo.Context) error {
	command := c.QueryParam("command")
	params := c.QueryParams()["params"]

	out, _ := a.service.VulnerableCommandTaint(command, params)
	return c.String(http.StatusOK, out)
}

func (a *CommandInjectionAdapter) vulnerableCommandTaintViaStruct(c echo.Context) error {
	command := c.QueryParam("command")
	params := c.QueryParams()["params"]

	out, _ := a.service.VulnerableCommandTaintViaStruct(command, params)
	return c.String(http.StatusOK, out)
}
