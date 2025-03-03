package todo

type Item struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Done   bool   `json:"done"`
	ListID uint   `json:"listID"`
	List   List   `json:"list"`
}

type List struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
