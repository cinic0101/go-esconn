package template

type SearchTemplate struct {
	Body map[string]interface{}
}

func New() *SearchTemplate {
	return &SearchTemplate{
		Body: make(map[string]interface{}),
	}
}

func (s *SearchTemplate) WithFromSize(from int, size int) *SearchTemplate {
	s.Body["from"] = from
	s.Body["size"] = size

	return s
}

func (s *SearchTemplate) MatchQuery(field string, keyword string) *SearchTemplate {
	s.Body["query"] = map[string]interface{} {
		"match": map[string]interface{} {
			field: keyword,
		},
	}

	return s
}