package entities

type ProductInfo struct {
	Id                 int
	Title              string
	Description        string
	ImageURL           string
	FullPrice          float32
	FinalPrice         float32
	PriceModifications float32
}
