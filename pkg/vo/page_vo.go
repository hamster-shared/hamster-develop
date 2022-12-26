package vo

type Page[T any] struct {
	Data     []T `json:"data"`
	Total    int `json:"total"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func NewPage[T any](data []T, total int, page int, pageSize int) Page[T] {
	return Page[T]{
		Data:     data,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
}
