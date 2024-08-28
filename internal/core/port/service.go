package port

import "context"

type CommandInjectionService interface {
	VulnerableCommand(command string, params []string) (string, error)
	VulnerableContextCommand(ctx context.Context, command string, params []string) (string, error)
	VulnerableCommandWithShell(command string, params []string) (string, error)
	NonVulnerableCommand(command string, params []string) (string, error)
}
