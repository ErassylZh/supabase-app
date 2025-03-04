package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
	"work-project/internal/airtable"
)

type AirTable interface {
	GetProducts(ctx context.Context) ([]airtable.BaseObject[airtable.ProductListResponse], error)
	GetPosts(ctx context.Context) ([]airtable.BaseObject[airtable.Post], error)
	GetStories(ctx context.Context) ([]airtable.BaseObject[airtable.Stories], error)
	GetHashtags(ctx context.Context) ([]airtable.BaseObject[airtable.Hashtag], error)
	GetCollections(ctx context.Context) ([]airtable.BaseObject[airtable.Collection], error)
	GetProductTags(ctx context.Context) ([]airtable.BaseObject[airtable.ProductTag], error)
	GetContests(ctx context.Context) ([]airtable.BaseObject[airtable.Contest], error)
	GetContestBooks(ctx context.Context) ([]airtable.BaseObject[airtable.ContestBook], error)
	GetContestPrizes(ctx context.Context) ([]airtable.BaseObject[airtable.ContestPrize], error)
}

type AirTableClient struct {
	client  *http.Client
	baseURL *url.URL
	apiKey  string
}

func NewAirTableClient(baseUrl string, apiKey string) (*AirTableClient, error) {
	baseURL, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	return &AirTableClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: baseURL,
		apiKey:  apiKey,
	}, nil
}

func (r *AirTableClient) GetProducts(ctx context.Context) ([]airtable.BaseObject[airtable.ProductListResponse], error) {
	var allRecords []airtable.BaseObject[airtable.ProductListResponse]
	var offset *string

	for {
		select {
		case <-ctx.Done(): // ✅ Прерываем цикл, если контекст отменен
			log.Println("⏳ GetProducts отменен:", ctx.Err())
			return nil, ctx.Err()
		default:
		}
		requestURL := r.baseURL.JoinPath("/Store")
		if offset != nil {
			query := requestURL.Query()
			query.Set("offset", *offset)
			requestURL.RawQuery = query.Encode()
		}

		req, err := r.newRequest(ctx, http.MethodGet, requestURL, nil)
		if err != nil {
			return nil, err
		}

		resp, err := r.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error while getting Airtable products. Response code: %d", resp.StatusCode)
		}

		rawResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var response airtable.BaseResponse[airtable.ProductListResponse]
		if err := json.Unmarshal(rawResponse, &response); err != nil {
			return nil, err
		}

		allRecords = append(allRecords, response.Records...)

		if response.Offset == nil {
			break
		}

		offset = response.Offset
	}

	return allRecords, nil
}

func (r *AirTableClient) GetPosts(ctx context.Context) ([]airtable.BaseObject[airtable.Post], error) {
	var allRecords []airtable.BaseObject[airtable.Post]
	var offset *string

	for {
		select {
		case <-ctx.Done(): // ✅ Прерываем цикл, если контекст отменен
			log.Println("⏳ GetPosts отменен:", ctx.Err())
			return nil, ctx.Err()
		default:
		}
		requestURL := r.baseURL.JoinPath("/Post")
		if offset != nil {
			query := requestURL.Query()
			query.Set("offset", *offset)
			requestURL.RawQuery = query.Encode()
		}

		req, err := r.newRequest(ctx, http.MethodGet, requestURL, nil)
		if err != nil {
			return nil, err
		}

		resp, err := r.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error while getting Airtable posts. Response code: %d", resp.StatusCode)
		}

		rawResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var response airtable.BaseResponse[airtable.Post]
		if err := json.Unmarshal(rawResponse, &response); err != nil {
			return nil, err
		}

		allRecords = append(allRecords, response.Records...)

		if response.Offset == nil {
			break
		}

		offset = response.Offset
	}

	return allRecords, nil
}

func (r *AirTableClient) GetStories(ctx context.Context) ([]airtable.BaseObject[airtable.Stories], error) {
	var allRecords []airtable.BaseObject[airtable.Stories]
	var offset *string

	for {
		select {
		case <-ctx.Done(): // ✅ Прерываем цикл, если контекст отменен
			log.Println("⏳ GetStories отменен:", ctx.Err())
			return nil, ctx.Err()
		default:
		}
		requestURL := r.baseURL.JoinPath("/Stories")
		if offset != nil {
			query := requestURL.Query()
			query.Set("offset", *offset)
			requestURL.RawQuery = query.Encode()
		}

		req, err := r.newRequest(ctx, http.MethodGet, requestURL, nil)
		if err != nil {
			return nil, err
		}

		resp, err := r.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error while getting Airtable stories. Response code: %d", resp.StatusCode)
		}

		rawResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var response airtable.BaseResponse[airtable.Stories]
		if err := json.Unmarshal(rawResponse, &response); err != nil {
			return nil, err
		}

		allRecords = append(allRecords, response.Records...)

		if response.Offset == nil {
			break
		}

		offset = response.Offset
	}

	return allRecords, nil
}

func (r *AirTableClient) GetHashtags(ctx context.Context) ([]airtable.BaseObject[airtable.Hashtag], error) {
	var allRecords []airtable.BaseObject[airtable.Hashtag]
	var offset *string

	for {
		select {
		case <-ctx.Done(): // ✅ Прерываем цикл, если контекст отменен
			log.Println("⏳ GetHashtags отменен:", ctx.Err())
			return nil, ctx.Err()
		default:
		}
		requestURL := r.baseURL.JoinPath("/Hashtags")
		if offset != nil {
			query := requestURL.Query()
			query.Set("offset", *offset)
			requestURL.RawQuery = query.Encode()
		}

		req, err := r.newRequest(ctx, http.MethodGet, requestURL, nil)
		if err != nil {
			return nil, err
		}

		resp, err := r.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error while getting Airtable hashtag. Response code: %d", resp.StatusCode)
		}

		rawResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var response airtable.BaseResponse[airtable.Hashtag]
		if err := json.Unmarshal(rawResponse, &response); err != nil {
			return nil, err
		}

		allRecords = append(allRecords, response.Records...)

		if response.Offset == nil {
			break
		}

		offset = response.Offset
	}

	return allRecords, nil
}

func (r *AirTableClient) GetCollections(ctx context.Context) ([]airtable.BaseObject[airtable.Collection], error) {
	var allRecords []airtable.BaseObject[airtable.Collection]
	var offset *string

	for {
		select {
		case <-ctx.Done(): // ✅ Прерываем цикл, если контекст отменен
			log.Println("⏳ GetCollections отменен:", ctx.Err())
			return nil, ctx.Err()
		default:
		}
		requestURL := r.baseURL.JoinPath("/Collections")
		if offset != nil {
			query := requestURL.Query()
			query.Set("offset", *offset)
			requestURL.RawQuery = query.Encode()
		}

		req, err := r.newRequest(ctx, http.MethodGet, requestURL, nil)
		if err != nil {
			return nil, err
		}

		resp, err := r.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error while getting Airtable hashtag. Response code: %d", resp.StatusCode)
		}

		rawResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var response airtable.BaseResponse[airtable.Collection]
		if err := json.Unmarshal(rawResponse, &response); err != nil {
			return nil, err
		}

		allRecords = append(allRecords, response.Records...)

		if response.Offset == nil {
			break
		}

		offset = response.Offset
	}

	return allRecords, nil
}

func (r *AirTableClient) GetProductTags(ctx context.Context) ([]airtable.BaseObject[airtable.ProductTag], error) {
	var allRecords []airtable.BaseObject[airtable.ProductTag]
	var offset *string

	for {
		select {
		case <-ctx.Done(): // ✅ Прерываем цикл, если контекст отменен
			log.Println("⏳ GetProductTags отменен:", ctx.Err())
			return nil, ctx.Err()
		default:
		}
		requestURL := r.baseURL.JoinPath("/StoreTag")
		if offset != nil {
			query := requestURL.Query()
			query.Set("offset", *offset)
			requestURL.RawQuery = query.Encode()
		}

		req, err := r.newRequest(ctx, http.MethodGet, requestURL, nil)
		if err != nil {
			return nil, err
		}

		resp, err := r.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error while getting Airtable hashtag. Response code: %d", resp.StatusCode)
		}

		rawResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var response airtable.BaseResponse[airtable.ProductTag]
		if err := json.Unmarshal(rawResponse, &response); err != nil {
			return nil, err
		}

		allRecords = append(allRecords, response.Records...)

		if response.Offset == nil {
			break
		}

		offset = response.Offset
	}

	return allRecords, nil
}

func (r *AirTableClient) GetContests(ctx context.Context) ([]airtable.BaseObject[airtable.Contest], error) {
	var allRecords []airtable.BaseObject[airtable.Contest]
	var offset *string

	for {
		select {
		case <-ctx.Done(): // ✅ Прерываем цикл, если контекст отменен
			log.Println("⏳ GetContests отменен:", ctx.Err())
			return nil, ctx.Err()
		default:
		}
		requestURL := r.baseURL.JoinPath("/Contest")
		if offset != nil {
			query := requestURL.Query()
			query.Set("offset", *offset)
			requestURL.RawQuery = query.Encode()
		}

		req, err := r.newRequest(ctx, http.MethodGet, requestURL, nil)
		if err != nil {
			return nil, err
		}

		resp, err := r.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error while getting Airtable hashtag. Response code: %d", resp.StatusCode)
		}

		rawResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var response airtable.BaseResponse[airtable.Contest]
		if err := json.Unmarshal(rawResponse, &response); err != nil {
			return nil, err
		}

		allRecords = append(allRecords, response.Records...)

		if response.Offset == nil {
			break
		}

		offset = response.Offset
	}

	return allRecords, nil
}

func (r *AirTableClient) GetContestBooks(ctx context.Context) ([]airtable.BaseObject[airtable.ContestBook], error) {
	var allRecords []airtable.BaseObject[airtable.ContestBook]
	var offset *string

	for {
		select {
		case <-ctx.Done(): // ✅ Прерываем цикл, если контекст отменен
			log.Println("⏳ GetContestBooks отменен:", ctx.Err())
			return nil, ctx.Err()
		default:
		}
		requestURL := r.baseURL.JoinPath("/ContestBook")
		if offset != nil {
			query := requestURL.Query()
			query.Set("offset", *offset)
			requestURL.RawQuery = query.Encode()
		}

		req, err := r.newRequest(ctx, http.MethodGet, requestURL, nil)
		if err != nil {
			return nil, err
		}

		resp, err := r.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error while getting Airtable ContestBook. Response code: %d", resp.StatusCode)
		}

		rawResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var response airtable.BaseResponse[airtable.ContestBook]
		if err := json.Unmarshal(rawResponse, &response); err != nil {
			return nil, err
		}

		allRecords = append(allRecords, response.Records...)

		if response.Offset == nil {
			break
		}

		offset = response.Offset
	}

	return allRecords, nil
}

func (r *AirTableClient) GetContestPrizes(ctx context.Context) ([]airtable.BaseObject[airtable.ContestPrize], error) {
	var allRecords []airtable.BaseObject[airtable.ContestPrize]
	var offset *string

	for {
		select {
		case <-ctx.Done(): // ✅ Прерываем цикл, если контекст отменен
			log.Println("⏳ GetContestPrizes отменен:", ctx.Err())
			return nil, ctx.Err()
		default:
		}
		requestURL := r.baseURL.JoinPath("/ContestPrize")
		if offset != nil {
			query := requestURL.Query()
			query.Set("offset", *offset)
			requestURL.RawQuery = query.Encode()
		}

		req, err := r.newRequest(ctx, http.MethodGet, requestURL, nil)
		if err != nil {
			return nil, err
		}

		resp, err := r.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error while getting Airtable ContestPrize. Response code: %d", resp.StatusCode)
		}

		rawResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var response airtable.BaseResponse[airtable.ContestPrize]
		if err := json.Unmarshal(rawResponse, &response); err != nil {
			return nil, err
		}

		allRecords = append(allRecords, response.Records...)

		if response.Offset == nil {
			break
		}

		offset = response.Offset
	}

	return allRecords, nil
}

func (r *AirTableClient) newRequest(
	ctx context.Context,
	method string,
	requestURL *url.URL,
	body interface{},
) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		rawBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(rawBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, requestURL.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	req.Header.Add("Authorization", "Bearer "+r.apiKey)

	return req, nil
}
