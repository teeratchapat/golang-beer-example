package logic

import (
	"golang-beer-example/models"
	"strings"
)

const DEFAULT_PAGINATION = 5

type InputPagination struct {
	Page int    // Current page number
	Size int    // Number of items per page
	Name string // Filter by name (optional)
}

type OutputPagination[T any] struct {
	Data  []T // Array of items (generic type)
	Total int // Total number of items in the collection
	Page  int // Current page number
	Size  int // Number of items per page
}

func calculatePagination(pagination InputPagination) InputPagination {
	if pagination.Page == 0 || pagination.Size == 0 {
		return InputPagination{
			Page: 1,
			Size: DEFAULT_PAGINATION,
		}
	}
	return InputPagination{
		Page: pagination.Page,
		Size: pagination.Size,
		Name: pagination.Name,
	}
}

func filterBeersByName(beers []models.Beer, name string) []models.Beer {
	var filteredBeers []models.Beer
	for _, beer := range beers {
		if strings.Contains(strings.ToLower(beer.Name), strings.ToLower(name)) {
			filteredBeers = append(filteredBeers, beer)
		}
	}
	return filteredBeers
}

func mapToOutputPagination(beers []models.Beer, pagination InputPagination) OutputPagination[models.Beer] {
	start := (pagination.Page - 1) * pagination.Size
	end := start + pagination.Size

	if start > len(beers) {
		start = len(beers)
	}
	if end > len(beers) {
		end = len(beers)
	}
	data := beers[start:end]
	return OutputPagination[models.Beer]{
		Data:  data,
		Total: len(beers),
		Page:  pagination.Page,
		Size:  pagination.Size,
	}
}

func NumberBasePaginate(inputPagination InputPagination, beers []models.Beer) OutputPagination[models.Beer] {
	pagination := calculatePagination(inputPagination)

	if pagination.Name != "" {
		beers = filterBeersByName(beers, pagination.Name)
	}

	output := mapToOutputPagination(beers, pagination)
	return output
}
