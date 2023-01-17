package tmdb

import "net/url"

type TrendingService struct {
	client *Client
}

type TrendingType struct {
	url    *url.URL
	client *Client
}

type TrendingTime struct {
	url    *url.URL
	client *Client
}

func (ts *TrendingService) All() *TrendingType {
	rel := &url.URL{Path: "/all"}

	trendingType := &TrendingType{
		url:    ts.client.BaseURL.ResolveReference(rel),
		client: ts.client,
	}

	return trendingType
}

func (ts *TrendingService) Movie() *TrendingType {
	rel := &url.URL{Path: "/movie"}

	trendingType := &TrendingType{
		url:    ts.client.BaseURL.ResolveReference(rel),
		client: ts.client,
	}

	return trendingType
}

func (ts *TrendingService) TV() *TrendingType {
	rel := &url.URL{Path: "/tv"}

	trendingType := &TrendingType{
		url:    ts.client.BaseURL.ResolveReference(rel),
		client: ts.client,
	}

	return trendingType
}

func (ts *TrendingService) Person() *TrendingType {
	rel := &url.URL{Path: "/person"}

	trendingType := &TrendingType{
		url:    ts.client.BaseURL.ResolveReference(rel),
		client: ts.client,
	}

	return trendingType
}

func (t *TrendingType) Day() *TrendingTime {
	rel := &url.URL{Path: "/day"}

	trendingTime := &TrendingTime{
		url:    t.url.ResolveReference(rel),
		client: t.client,
	}

	return trendingTime
}

func (t *TrendingType) Week() *TrendingTime {
	rel := &url.URL{Path: "/week"}

	trendingTime := &TrendingTime{
		url:    t.url.ResolveReference(rel),
		client: t.client,
	}

	return trendingTime
}

func (t *TrendingTime) Get() (interface{}, error) {
	req, err := t.client.NewRequest("GET", t.url, nil)
	if err != nil {
		return nil, err
	}

	var request interface{}
	_, err = t.client.Do(req, request)

	return request, err
}
