package dto

import "github.com/korzepadawid/go-ly/pkg/model"

type URLRequestDTO struct {
	URL string `json:"url"`
}

type URLResponseDTO struct {
	Base62     string `json:"base62"`
	URL        string `json:"url"`
	AuthorIPv4 string `json:"authorIpv4"`
}

func (dto *URLRequestDTO) Map() *model.URL {
	return &model.URL{
		Url: dto.URL,
	}
}

func (dto *URLResponseDTO) Map(u *model.URL, base62 string) {
	dto.AuthorIPv4 = u.AuthorIPv4
	dto.URL = u.Url
	dto.Base62 = base62
}
