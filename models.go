package AnilistAPIClientGo

import (
	"encoding/json"
)

type QueryID struct {
	ID int `json:"id"`
}

type APIRequestID struct {
	Query     string  `json:"query"`
	Variables QueryID `json:"variables"`
}

func (a *APIRequestID) Marshall() ([]byte, error) {
	return json.Marshal(a)
}

func NewAPIRequestID(query string, id int) *APIRequestID {
	return &APIRequestID{
		Query: query,
		Variables: QueryID{
			ID: id,
		},
	}
}

type QuerySearch struct {
	Search string `json:"search"`
}

type APIRequestSearch struct {
	Query     string      `json:"query"`
	Variables QuerySearch `json:"variables"`
}

func (a *APIRequestSearch) Marshall() ([]byte, error) {
	return json.Marshal(a)
}

func NewAPIRequestSearch(query string, search string) *APIRequestSearch {
	return &APIRequestSearch{
		Query: query,
		Variables: QuerySearch{
			Search: search,
		},
	}
}

type APIAnimeResponse struct {
	Data struct {
		Media struct {
			ID      int  `json:"id"`
			IsAdult bool `json:"isAdult"`
			Title   struct {
				Romaji  string `json:"romaji"`
				English string `json:"english"`
				Native  string `json:"native"`
			} `json:"title"`
			Description string `json:"description"`
			StartDate   struct {
				Year  int `json:"year"`
				Month int `json:"month"`
				Day   int `json:"day"`
			} `json:"startDate"`
			Episodes    int    `json:"episodes"`
			Season      string `json:"season"`
			Type        string `json:"type"`
			Format      string `json:"format"`
			Status      string `json:"status"`
			Duration    int    `json:"duration"`
			IsFavourite bool   `json:"isFavourite"`
			SiteURL     string `json:"siteUrl"`
			Studios     struct {
				Nodes []struct {
					Name string `json:"name"`
				} `json:"nodes"`
			} `json:"studios"`
			Tags []struct {
				Name string `json:"name"`
			} `json:"tags"`
			Trailer      interface{} `json:"trailer"`
			AverageScore int         `json:"averageScore"`
			Genres       []string    `json:"genres"`
			BannerImage  string      `json:"bannerImage"`
		} `json:"Media"`
	} `json:"data"`
}

func (a *APIAnimeResponse) Unmarshall(data []byte) error {
	return json.Unmarshal(data, a)
}

func NewAPIAnimeResponse() *APIAnimeResponse {
	return &APIAnimeResponse{}
}

type APICharacterResponse struct {
	Data struct {
		Character struct {
			ID   int `json:"id"`
			Name struct {
				First  string `json:"first"`
				Last   string `json:"last"`
				Full   string `json:"full"`
				Native string `json:"native"`
			} `json:"name"`
			SiteURL string `json:"siteUrl"`
			Image   struct {
				Large string `json:"large"`
			} `json:"image"`
			Media struct {
				Edges []struct {
					CharacterRole string `json:"characterRole"`
					Node          struct {
						ID    int `json:"id"`
						Title struct {
							Romaji string `json:"romaji"`
						} `json:"title"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"media"`
			Description string `json:"description"`
		} `json:"Character"`
	} `json:"data"`
}

func (a *APICharacterResponse) Unmarshall(data []byte) error {
	return json.Unmarshal(data, a)
}

func NewAPICharacterResponse() *APICharacterResponse {
	return &APICharacterResponse{}
}

type APIMangaResponse struct {
	Data struct {
		Media struct {
			ID    int `json:"id"`
			Title struct {
				Romaji  string `json:"romaji"`
				English string `json:"english"`
				Native  string `json:"native"`
			} `json:"title"`
			Description string `json:"description"`
			StartDate   struct {
				Year int `json:"year"`
			} `json:"startDate"`
			Type         string   `json:"type"`
			Format       string   `json:"format"`
			Status       string   `json:"status"`
			SiteURL      string   `json:"siteUrl"`
			AverageScore int      `json:"averageScore"`
			Genres       []string `json:"genres"`
			BannerImage  string   `json:"bannerImage"`
			IsAdult      bool     `json:"isAdult"`
			IsFavourite  bool     `json:"isFavourite"`
		} `json:"Media"`
	} `json:"data"`
}

func (a *APIMangaResponse) Unmarshall(data []byte) error {
	return json.Unmarshal(data, a)
}

func NewAPIMangaResponse() *APIMangaResponse {
	return &APIMangaResponse{}
}

type APIAnimeAiringResponse struct {
	Data struct {
		Media struct {
			ID       int    `json:"id"`
			Status   string `json:"status"`
			Episodes int    `json:"episodes"`
			Title    struct {
				Romaji  string `json:"romaji"`
				English string `json:"english"`
				Native  string `json:"native"`
			} `json:"title"`
			NextAiringEpisode struct {
				AiringAt        int `json:"airingAt"`
				TimeUntilAiring int `json:"timeUntilAiring"`
				Episode         int `json:"episode"`
			} `json:"nextAiringEpisode"`
		} `json:"Media"`
	} `json:"data"`
}

func (a *APIAnimeAiringResponse) Unmarshall(data []byte) error {
	return json.Unmarshal(data, a)
}

func NewAPIAnimeAiringResponse() *APIAnimeAiringResponse {
	return &APIAnimeAiringResponse{}
}

type APIErrorResponse struct {
	Errors []struct {
		Message   string `json:"message"`
		Status    int    `json:"status"`
		Locations []struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"locations"`
	} `json:"errors"`
	Data struct {
		Media interface{} `json:"Media"`
	} `json:"data"`
}

func (a *APIErrorResponse) Unmarshall(data []byte) error {
	return json.Unmarshal(data, a)
}

func NewAPIErrorResponse(err error) *APIErrorResponse {
	resp := &APIErrorResponse{}
	resp.Unmarshall([]byte(err.Error()))
	return resp
}
