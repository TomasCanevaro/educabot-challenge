package services

import (
	"context"
	"slices"

	"educabot.com/bookshop/models"
	"educabot.com/bookshop/providers"
)

type MetricsService struct {
	booksProvider providers.BooksProvider
}

func NewMetricsService(booksProvider providers.BooksProvider) MetricsService {
	return MetricsService{booksProvider}
}

func (s MetricsService) GetMetrics(ctx context.Context, author string) (map[string]interface{}, error) {
	books, err := s.booksProvider.GetBooks(ctx)
	if err != nil {
		return nil, err
	}

	meanUnitsSold := calculateMeanUnitsSold(books)
	cheapestBook := findCheapestBook(books).Name
	booksByAuthor := countBooksByAuthor(books, author)

	return map[string]interface{}{
		"mean_units_sold":         meanUnitsSold,
		"cheapest_book":           cheapestBook,
		"books_written_by_author": booksByAuthor,
	}, nil
}

func calculateMeanUnitsSold(books []models.Book) uint {
	var sum uint
	for _, book := range books {
		sum += book.UnitsSold
	}
	return sum / uint(len(books))
}

func findCheapestBook(books []models.Book) models.Book {
	return slices.MinFunc(books, func(a, b models.Book) int {
		return int(a.Price - b.Price)
	})
}

func countBooksByAuthor(books []models.Book, author string) uint {
	var count uint
	for _, book := range books {
		if book.Author == author {
			count++
		}
	}
	return count
}
