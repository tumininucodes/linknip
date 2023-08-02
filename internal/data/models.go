package data

type LinkRequest struct {
	Url 		string `json:"url"`
	CustomSlug 	string `json:"customSlug"`
}

type LinkResponse struct {
	ShortenedUrl string `json:"shortenedUrl"`
}

type Link struct {
	Id	uint64 `json:"id"`
	Url string `json:"url"`		
}