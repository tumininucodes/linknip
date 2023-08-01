package data

type LinkRequest struct {
	Url 			string `json:"url"`
	CustomSlug 		string `json:"customSlug"`
}

type LinkResponse struct {
	ShortenedUrl	string `json:"shortenedUrl"`
}

