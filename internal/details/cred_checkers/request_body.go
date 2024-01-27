package cred_checkers

type RequestBody struct {
	Filter             Filter   `json:"filter"`
	TreeFiltering      string   `json:"treeFiltering"`
	PageNum            int      `json:"pageNum"`
	PageSize           int      `json:"pageSize"`
	ParentRefItemValue string   `json:"parentRefItemValue"`
	SelectAttributes   []string `json:"selectAttributes"`
	Tx                 string   `json:"tx"`
}

type Value struct {
	AsString string `json:"asString"`
}

type Simple struct {
	AttributeName string `json:"attributeName"`
	Condition     string `json:"condition"`
	Value         Value  `json:"value"`
}

type Subs struct {
	Simple Simple `json:"simple"`
}

type Union struct {
	UnionKind string `json:"unionKind"`
	Subs      []Subs `json:"subs"`
}

type Filter struct {
	Union Union `json:"union"`
}
