package request

type List struct {
	Search  string `json:"search"`
	Sort    string `json:"sort"`
	PerPage int    `json:"per_page"`
	Page    int    `json:"page"`
}
