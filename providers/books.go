package providers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"educabot.com/bookshop/models"
)

type BooksProvider interface {
	GetBooks(ctx context.Context) ([]models.Book, error)
}

type ApiBooksProvider struct {
	endpoint string
}

func NewApiBooksProvider(endpoint string) ApiBooksProvider {
	return ApiBooksProvider{endpoint: endpoint}
}

func (p ApiBooksProvider) GetBooks(ctx context.Context) ([]models.Book, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.endpoint, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch books: " + resp.Status)
	}

	var books []models.Book
	if err := json.NewDecoder(resp.Body).Decode(&books); err != nil {
		return nil, err
	}

	return books, nil
}
