// Code generated by goctl. DO NOT EDIT.
package types

type GetBooksReq struct {
	Page   int    `form:"page,default=1"`
	Limit  int    `form:"limit,default=10"`
	Name   string `form:"name,optional"`
	Author string `form:"author,optional"`
}

type AddBookReq struct {
	Name   string `json:"name"`
	Image  string `json:"image,optional"`
	Author string `json:"author,optional"`
	Info   string `json:"info,optional"`
}

type UpdateBookReq struct {
	ID     uint   `path:"id"`
	Name   string `json:"name"`
	Image  string `json:"image,optional"`
	Author string `json:"author,optional"`
	Info   string `json:"info,optional"`
}

type DeleteBookReq struct {
	ID uint `path:"id"`
}

type BookData struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Author string `json:"author"`
	Info   string `json:"info"`
}

type Reply struct {
	Code int                    `json:"code"`
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"msg"`
}