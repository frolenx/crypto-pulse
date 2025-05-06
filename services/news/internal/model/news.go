package model

type GetNewsResponse struct {
	Results []*News `json:"results"`
}

type News struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PublishedAt string `json:"published_at"`
}
