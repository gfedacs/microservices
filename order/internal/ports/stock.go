package ports

type StockPort interface {

	
	ExistsStockItems(productCodes []string) (bool, []string, error)



}