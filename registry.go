package control

// CommandRegistry holds a list of all the commands that are registered.
type CommandRegistry struct {
	cmds map[string]ControlCommand
}

// NewCommandRegistry returns a new commandregistry for storing actions to perform
func NewCommandRegistry() (CommandRegistry, error) {
	return CommandRegistry{
		cmds: make(map[string]ControlCommand),
	}, nil
}
