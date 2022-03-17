package entities

/*ProductInfo is an entity used to pass around and store Product information*/
type ProductInfo struct {
	ID                 int
	Title              string
	Description        string
	ImageURL           string
	FullPrice          float32
	FinalPrice         float32
	PriceModifications float32
}
