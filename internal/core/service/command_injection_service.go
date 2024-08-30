package service

import (
	"context"
	"gitea.homelab.d3s34.me/d3s34/vuln-app/internal/core/port"
	"os/exec"
)

type service struct{}

func (s *service) VulnerableCommand(command string, params []string) (string, error) {
	out, _ := exec.Command(command, params...).Output()
	return string(out), nil
}

func (s *service) VulnerableContextCommand(ctx context.Context, command string, params []string) (string, error) {
	out, _ := exec.CommandContext(ctx, command, params...).Output()
	return string(out), nil
}

func (s *service) VulnerableCommandWithShell(command string, params []string) (string, error) {
	out, _ := exec.Command("sh", "-c", command).Output()
	return string(out), nil
}

func (s *service) NonVulnerableCommand(_ string, params []string) (string, error) {
	out, _ := exec.Command("echo", params...).Output()
	return string(out), nil
}

func (s *service) VulnerableCommandViaStruct(command port.CommandStruct) (string, error) {
	out, _ := exec.Command(command.Command, command.Args...).Output()
	return string(out), nil
}

func (s *service) VulnerableCommandViaInterface(command port.CommandInterface) (string, error) {
	out, _ := exec.Command(command.GetCommand(), command.GetArgs()...).Output()
	return string(out), nil
}

func (s *service) VulnerableCommandTaint(command string, params []string) (string, error) {
	commandX := DoNothing(command)
	out, _ := exec.Command(commandX, params...).Output()
	return string(out), nil
}

func (s *service) VulnerableCommandTaintViaStruct(command string, params []string) (string, error) {
	commandY := port.CommandStruct{
		Command: command,
		Args:    params,
	}
	out, _ := exec.Command(commandY.GetCommand(), commandY.GetArgs()...).Output()
	return string(out), nil
}

func NewCommandInjectionService() port.CommandInjectionService {
	return &service{}
}
