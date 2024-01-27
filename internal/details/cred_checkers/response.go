package cred_checkers

type ResponseBody struct {
	Total    int     `json:"total"`
	LastPage bool    `json:"lastPage"`
	Items    []Items `json:"items"`
}

type Items struct {
	Name     string `json:"name"`
	Inn      string `json:"inn"`
	Fullname string `json:"fullname"`
	Type     string `json:"type"`
}
