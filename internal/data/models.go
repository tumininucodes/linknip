package data

type LinkRequest struct {
	Url 		string `json:"url"`
	CustomSlug 	string `json:"customSlug"`
}

type LinkResponse struct {
	ShortenedUrl string `json:"shortenedUrl"`
}

type Link struct {
	Id	string `json:"id"`
	Slug string `json:"slug"`
	Url string `json:"url"`		
}