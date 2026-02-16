package handlers

type Handler func(Command) (string, error)
