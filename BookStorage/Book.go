package BookStorage

// Book struct for reading from json and do operations.
type Book struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	PageNumber int    `json:"pageNumber"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
	StockCode  string `json:"stockCode"`
	ISBN       string `json:"isbn"`
	Author     Author `json:"author"`
}

// Author reading from json and do operations
type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
