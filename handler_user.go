package main

import (
	"context"
	"fmt"
	"time"

	"boot.dev-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Printf("Logged in as %s\n", name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	user, err := s.db.CreateUser(context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			Name:      name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}
	fmt.Println("User created successfully:")
	fmt.Printf("%+v\n", user)
	return nil
}
