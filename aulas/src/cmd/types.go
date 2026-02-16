package cmd

import "spanishGab/aula_camada_model/src/handlers"

type CLIParser func([]string) (*handlers.Command, error)
