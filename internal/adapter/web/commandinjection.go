package web

import (
	"gitea.homelab.d3s34.me/d3s34/vuln-app/internal/core/port"
	"github.com/gin-gonic/gin"
)

type CommandInjectionAdapter struct {
	service port.CommandInjectionService
}

func NewCommandInjectionAdapter(service port.CommandInjectionService) *CommandInjectionAdapter {
	return &CommandInjectionAdapter{
		service: service,
	}
}

func (a *CommandInjectionAdapter) vulnerableCommand(c *gin.Context) {
	command := c.Query("command")
	params := c.QueryArray("params")

	out, _ := a.service.VulnerableCommand(command, params)
	c.JSON(200, gin.H{"output": out})
}

func (a *CommandInjectionAdapter) vulnerableContextCommand(c *gin.Context) {
	command := c.Query("command")
	params := c.QueryArray("params")

	out, _ := a.service.VulnerableContextCommand(c, command, params)
	c.JSON(200, gin.H{"output": out})
}

func (a *CommandInjectionAdapter) vulnerableCommandWithShell(c *gin.Context) {
	command := c.Query("command")
	params := c.QueryArray("params")

	out, _ := a.service.VulnerableCommandWithShell(command, params)
	c.JSON(200, gin.H{"output": out})
}

func (a *CommandInjectionAdapter) nonVulnerableCommand(c *gin.Context) {
	command := c.Query("command")
	params := c.QueryArray("params")

	out, _ := a.service.NonVulnerableCommand(command, params)
	c.JSON(200, gin.H{"output": out})
}

func (a *CommandInjectionAdapter) vulnerableCommandViaStruct(c *gin.Context) {
	command := c.Query("command")
	params := c.QueryArray("params")

	out, _ := a.service.VulnerableCommandViaStruct(port.CommandStruct{
		Command: command,
		Args:    params,
	})
	c.JSON(200, gin.H{"output": out})
}

func (a *CommandInjectionAdapter) vulnerableCommandViaInterface(c *gin.Context) {
	command := c.Query("command")
	params := c.QueryArray("params")

	out, _ := a.service.VulnerableCommandViaInterface(&port.CommandStruct{
		Command: command,
		Args:    params,
	})
	c.JSON(200, gin.H{"output": out})
}

func (a *CommandInjectionAdapter) RegisterRoutes(rg *gin.RouterGroup) {
	commandInjection := rg.Group("/command-injection")
	{
		commandInjection.GET("/", a.vulnerableCommand)
		commandInjection.GET("/context", a.vulnerableContextCommand)
		commandInjection.GET("/shell", a.vulnerableCommandWithShell)
		commandInjection.GET("/non-vulnerable", a.nonVulnerableCommand)
		commandInjection.GET("/struct", a.vulnerableCommandViaStruct)
		commandInjection.GET("/interface", a.vulnerableCommandViaInterface)
	}
}
