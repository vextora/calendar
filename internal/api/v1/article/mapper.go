package article

import "oncomapi/internal/api/v1/article/dto"

func DataList(data []*Article) []dto.ArticleResponse {
	result := make([]dto.ArticleResponse, 0, len(data))
	for _, a := range data {
		result = append(result, StandardMap(a))
	}
	return result
}

func StandardMap(data *Article) dto.ArticleResponse {
	return dto.ArticleResponse{
		ID:    data.ID,
		Title: data.Title,
		Slug:  data.Slug,
	}
}

func DetailResponse(data *Article) dto.ArticleResponse {
	return dto.ArticleResponse{
		ID:      data.ID,
		Title:   data.Title,
		Slug:    data.Slug,
		Content: data.Content,
	}
}
