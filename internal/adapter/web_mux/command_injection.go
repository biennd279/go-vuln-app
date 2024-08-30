// Package web_mux internal/adapter/web_mux/command_injection.go
package web_mux

import (
	"gitea.homelab.d3s34.me/d3s34/vuln-app/internal/core/port"
	"github.com/gorilla/mux"
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

func (a *CommandInjectionAdapter) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/vulnerable-command", a.vulnerableCommand).Methods("GET")
	r.HandleFunc("/vulnerable-context-command", a.vulnerableContextCommand).Methods("GET")
	r.HandleFunc("/vulnerable-command-with-shell", a.vulnerableCommandWithShell).Methods("GET")
	r.HandleFunc("/non-vulnerable-command", a.nonVulnerableCommand).Methods("GET")
	r.HandleFunc("/vulnerable-command-via-struct", a.vulnerableCommandViaStruct).Methods("GET")
	r.HandleFunc("/vulnerable-command-via-interface", a.vulnerableCommandViaInterface).Methods("GET")
	r.HandleFunc("/vulnerable-command-taint", a.vulnerableCommandTaint).Methods("GET")
	r.HandleFunc("/vulnerable-command-taint-via-struct", a.vulnerableCommandTaintViaStruct).Methods("GET")
	r.HandleFunc("/vulnerable-command-taint-via-interface", a.vulnerableCommandTaintViaInterface).Methods("GET")
}

func (a *CommandInjectionAdapter) vulnerableCommand(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")
	params := r.URL.Query()["params"]

	out, _ := a.service.VulnerableCommand(command, params)
	w.Write([]byte(out))
}

func (a *CommandInjectionAdapter) vulnerableContextCommand(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")
	params := r.URL.Query()["params"]

	out, _ := a.service.VulnerableContextCommand(r.Context(), command, params)
	w.Write([]byte(out))
}

func (a *CommandInjectionAdapter) vulnerableCommandWithShell(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")
	params := r.URL.Query()["params"]

	out, _ := a.service.VulnerableCommandWithShell(command, params)
	w.Write([]byte(out))
}

func (a *CommandInjectionAdapter) nonVulnerableCommand(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")
	params := r.URL.Query()["params"]

	out, _ := a.service.NonVulnerableCommand(command, params)
	w.Write([]byte(out))
}

func (a *CommandInjectionAdapter) vulnerableCommandViaStruct(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")
	params := r.URL.Query()["params"]

	out, _ := a.service.VulnerableCommandViaStruct(port.CommandStruct{
		Command: command,
		Args:    params,
	})
	w.Write([]byte(out))
}

func (a *CommandInjectionAdapter) vulnerableCommandViaInterface(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")
	params := r.URL.Query()["params"]

	commandStruct := port.CommandStruct{
		Command: command,
		Args:    params,
	}

	out, _ := a.service.VulnerableCommandViaInterface(commandStruct)
	w.Write([]byte(out))
}

func (a *CommandInjectionAdapter) vulnerableCommandTaint(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")
	params := r.URL.Query()["params"]

	out, _ := a.service.VulnerableCommandTaint(command, params)
	w.Write([]byte(out))
}

func (a *CommandInjectionAdapter) vulnerableCommandTaintViaStruct(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")
	params := r.URL.Query()["params"]

	out, _ := a.service.VulnerableCommandTaintViaStruct(command, params)
	w.Write([]byte(out))
}

func (a *CommandInjectionAdapter) vulnerableCommandTaintViaInterface(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")
	params := r.URL.Query()["params"]

	commandStruct := port.CommandStruct{
		Command: command,
		Args:    params,
	}

	out, _ := a.service.VulnerableCommandViaInterface(commandStruct)
	w.Write([]byte(out))
}
