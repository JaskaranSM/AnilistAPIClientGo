// AnilistAPIClientGo project AnilistAPIClientGo.go
package AnilistAPIClientGo

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type AnilistClient struct {
	client  *http.Client
	BaseUrl string
}

func (a *AnilistClient) SetClient(client *http.Client) {
	a.client = client
}

func (a *AnilistClient) QueryAnimeByID(id int) (*APIAnimeResponse, error) {
	anime := NewAPIAnimeResponse()
	data := NewAPIRequestID(AnimeQuery, id)
	rawBytes, err := data.Marshall()
	if err != nil {
		return anime, err
	}
	respBytes, err := a.DoRequest(rawBytes)
	if err != nil {
		return anime, err
	}
	if IsAPIErrorBytes(respBytes) {
		return anime, errors.New(string(respBytes))
	}
	err = anime.Unmarshall(respBytes)
	if err != nil {
		return anime, err
	}
	return anime, nil
}

func (a *AnilistClient) QueryAnimeBySearch(search string) (*APIAnimeResponse, error) {
	anime := NewAPIAnimeResponse()
	data := NewAPIRequestSearch(AnimeQuery, search)
	rawBytes, err := data.Marshall()
	if err != nil {
		return anime, err
	}
	respBytes, err := a.DoRequest(rawBytes)
	if err != nil {
		return anime, err
	}
	if IsAPIErrorBytes(respBytes) {
		return anime, errors.New(string(respBytes))
	}
	err = anime.Unmarshall(respBytes)
	if err != nil {
		return anime, err
	}
	return anime, nil
}

func (a *AnilistClient) QueryPagedAnimeBySearch(search string) (*APIAnimePagedResponse, error) {
	anime := NewAPIAnimePagedResponse()
	data := NewAPIRequestSearch(AnimeQueryPaged, search)
	rawBytes, err := data.Marshall()
	if err != nil {
		return anime, err
	}
	respBytes, err := a.DoRequest(rawBytes)
	if err != nil {
		return anime, err
	}
	if IsAPIErrorBytes(respBytes) {
		return anime, errors.New(string(respBytes))
	}
	err = anime.Unmarshall(respBytes)
	if err != nil {
		return anime, err
	}
	return anime, nil
}

func (a *AnilistClient) QueryPagedMangaBySearch(search string) (*APIMangaPagedResponse, error) {
	manga := NewAPIMangaPagedResponse()
	data := NewAPIRequestSearch(MangaQueryPaged, search)
	rawBytes, err := data.Marshall()
	if err != nil {
		return manga, err
	}
	respBytes, err := a.DoRequest(rawBytes)
	if err != nil {
		return manga, err
	}
	if IsAPIErrorBytes(respBytes) {
		return manga, errors.New(string(respBytes))
	}
	err = manga.Unmarshall(respBytes)
	if err != nil {
		return manga, err
	}
	return manga, nil
}

func (a *AnilistClient) QueryCharacterByID(id int) (*APICharacterResponse, error) {
	character := NewAPICharacterResponse()
	data := NewAPIRequestID(CharacterQuery, id)
	rawBytes, err := data.Marshall()
	if err != nil {
		return character, err
	}
	respBytes, err := a.DoRequest(rawBytes)
	if err != nil {
		return character, err
	}
	if IsAPIErrorBytes(respBytes) {
		return character, errors.New(string(respBytes))
	}
	err = character.Unmarshall(respBytes)
	if err != nil {
		return character, err
	}
	return character, nil
}

func (a *AnilistClient) QueryCharacterBySearch(search string) (*APICharacterResponse, error) {
	character := NewAPICharacterResponse()
	data := NewAPIRequestSearch(CharacterQuery, search)
	rawBytes, err := data.Marshall()
	if err != nil {
		return character, err
	}
	respBytes, err := a.DoRequest(rawBytes)
	if err != nil {
		return character, err
	}
	if IsAPIErrorBytes(respBytes) {
		return character, errors.New(string(respBytes))
	}
	err = character.Unmarshall(respBytes)
	if err != nil {
		return character, err
	}
	return character, nil
}

func (a *AnilistClient) QueryMangaByID(id int) (*APIMangaResponse, error) {
	manga := NewAPIMangaResponse()
	data := NewAPIRequestID(MangaQuery, id)
	rawBytes, err := data.Marshall()
	if err != nil {
		return manga, err
	}
	respBytes, err := a.DoRequest(rawBytes)
	if err != nil {
		return manga, err
	}
	if IsAPIErrorBytes(respBytes) {
		return manga, errors.New(string(respBytes))
	}
	err = manga.Unmarshall(respBytes)
	if err != nil {
		return manga, err
	}
	return manga, nil
}

func (a *AnilistClient) QueryMangaBySearch(search string) (*APIMangaResponse, error) {
	manga := NewAPIMangaResponse()
	data := NewAPIRequestSearch(MangaQuery, search)
	rawBytes, err := data.Marshall()
	if err != nil {
		return manga, err
	}
	respBytes, err := a.DoRequest(rawBytes)
	if err != nil {
		return manga, err
	}
	if IsAPIErrorBytes(respBytes) {
		return manga, errors.New(string(respBytes))
	}
	err = manga.Unmarshall(respBytes)
	if err != nil {
		return manga, err
	}
	return manga, nil
}

func (a *AnilistClient) QueryAnimeAiringByID(id int) (*APIAnimeAiringResponse, error) {
	airing := NewAPIAnimeAiringResponse()
	data := NewAPIRequestID(AnimeAiringQuery, id)
	rawBytes, err := data.Marshall()
	if err != nil {
		return airing, err
	}
	respBytes, err := a.DoRequest(rawBytes)
	if err != nil {
		return airing, err
	}
	if IsAPIErrorBytes(respBytes) {
		return airing, errors.New(string(respBytes))
	}
	err = airing.Unmarshall(respBytes)
	if err != nil {
		return airing, err
	}
	return airing, nil
}

func (a *AnilistClient) QueryAnimeAiringBySearch(search string) (*APIAnimeAiringResponse, error) {
	airing := NewAPIAnimeAiringResponse()
	data := NewAPIRequestSearch(AnimeAiringQuery, search)
	rawBytes, err := data.Marshall()
	if err != nil {
		return airing, err
	}
	respBytes, err := a.DoRequest(rawBytes)
	if err != nil {
		return airing, err
	}
	if IsAPIErrorBytes(respBytes) {
		return airing, errors.New(string(respBytes))
	}
	err = airing.Unmarshall(respBytes)
	if err != nil {
		return airing, err
	}
	return airing, nil
}

func (a *AnilistClient) DoRequest(data []byte) ([]byte, error) {
	request, err := a.CookRequest(data)
	if err != nil {
		return nil, err
	}
	resp, err := a.client.Do(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func (a *AnilistClient) CookRequest(data []byte) (*http.Request, error) {
	request, err := http.NewRequest("POST", a.BaseUrl, GetBytesReader(data))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	return request, nil
}

func NewAnlistClient() *AnilistClient {
	return &AnilistClient{
		BaseUrl: "https://graphql.anilist.co/",
		client:  &http.Client{},
	}
}
