package database

type Query struct {
	Page   int    `form:"page,default=1"`
	Size   int    `form:"size,default=10"`
	Sort   string `form:"sort,default=id"`
	Order  string `form:"order,default=desc"`
	Search string `form:"search"`
}

type Response struct {
	Message string `json:"message"`
}

type Error struct {
	Response Response
	Error    error
}

type GetByIDArgs struct {
	Include map[string][]string
	Exclude []string
}

type GetArgs[T any] struct {
	Include map[string][]string
	Exclude []string
	Filter  T
}

type GetAllArgs[T any] struct {
	Include    map[string][]string
	Exclude    []string
	Search     bool
	Pagination bool
	Filter     T
	MapFilter  map[string]any
}

type GetResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type GetAllResponse[T any] struct {
	Message     string  `json:"message"`
	Data        []T     `json:"data"`
	Length      int     `json:"length"`
	Count       int64   `json:"count,omitempty"`
	CurrentPage int     `json:"currentPage,omitempty"`
	TotalPage   float64 `json:"totalPage,omitempty"`
}
