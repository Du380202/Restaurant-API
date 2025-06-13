package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filte  interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, Filte: filter}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return NewSuccessResponse(data, nil, nil)
}
