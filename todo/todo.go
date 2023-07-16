package todo

type Item struct {
	ID   uint `json:"id"`
	Name uint `json:"name"`
	Done bool `json:"done"`
	List List `json:"list"`
}

type List struct {
	ID    uint   `json:"id"`
	Name  uint   `json:"name"`
	Items []Item `json:"items"`
}
