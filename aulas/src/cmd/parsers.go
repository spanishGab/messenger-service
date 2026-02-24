package cmd

import (
	"fmt"
	"slices"
	"spanishGab/aula_camada_model/src/handlers"
)

const (
	PageOption     = "--page"
	PageSizeOption = "--page_size"
	FormatOption   = "--format"
)

func parsePaginationCommandData(commandData []string) (*handlers.CommandData, error) {
	pageIndex := slices.Index(commandData, PageOption)
	pageSizeIndex := slices.Index(commandData, PageSizeOption)
	if pageIndex == -1 || pageSizeIndex == -1 {
		return nil, fmt.Errorf("you should provide --page and --page_size arguments")
	}
	return &handlers.CommandData{
		"limit":  commandData[pageSizeIndex+1],
		"offset": commandData[pageIndex+1],
	}, nil
}

func parseFormatCommandData(commandData []string) (*handlers.CommandData, error) {
	formatIndex := slices.Index(commandData, FormatOption)
	if formatIndex == -1 {
		return nil, nil
	}
	format := handlers.OutputFormat(commandData[formatIndex+1])
	if !format.IsValid() {
		return nil, fmt.Errorf("the given format output is not valid")
	}
	return &handlers.CommandData{
		"format": format.String(),
	}, nil
}
