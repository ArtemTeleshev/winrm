package winrm

// Shell is the local view of a WinRM Shell of a given Client
type Shell struct {
	client  *Client
	ShellId string
}

// Execute command on the given Shell, returning either an error or a Command
func (shell *Shell) Execute(command string) (cmd *Command, err error) {
	request := NewExecuteCommandRequest(shell.client.Parameters.url, shell.ShellId, command, &shell.client.Parameters)
	defer request.Free()

	response, err := shell.client.sendRequest(request)
	if err == nil {
		var commandId string
		if commandId, err = ParseExecuteCommandResponse(response); err == nil {
			cmd = newCommand(shell, commandId)
		}
	}
	return
}

// Close will terminate this shell. No commands can be issued once the shell is closed.
func (shell *Shell) Close() (err error) {
	request := NewDeleteShellRequest(shell.client.Parameters.url, shell.ShellId, &shell.client.Parameters)
	defer request.Free()

	_, err = shell.client.sendRequest(request)
	return
}
