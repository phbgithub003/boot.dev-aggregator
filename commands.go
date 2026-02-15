package main

import (
	"errors"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) Register(name string, fn func(*state, command) error) {
	c.registeredCommands[name] = fn
}

func (c *commands) Run(s *state, cmd command) error {
	fn, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return fn(s, cmd)
}
