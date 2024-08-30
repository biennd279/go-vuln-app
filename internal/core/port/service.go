package port

import "context"

type CommandStruct struct {
	Command string
	Args    []string
}

func (c CommandStruct) GetCommand() string {
	return c.Command
}

func (c CommandStruct) GetArgs() []string {
	return c.Args
}

type CommandInterface interface {
	GetCommand() string
	GetArgs() []string
}

type CommandInjectionService interface {
	VulnerableCommand(command string, params []string) (string, error)
	VulnerableContextCommand(ctx context.Context, command string, params []string) (string, error)
	VulnerableCommandWithShell(command string, params []string) (string, error)
	NonVulnerableCommand(command string, params []string) (string, error)
	VulnerableCommandViaStruct(command CommandStruct) (string, error)
	VulnerableCommandViaInterface(command CommandInterface) (string, error)
	VulnerableCommandTaint(command string, params []string) (string, error)
	VulnerableCommandTaintViaStruct(command string, params []string) (string, error)
}
