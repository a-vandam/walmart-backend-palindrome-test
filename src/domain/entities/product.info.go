package entities

type ProductInfo struct {
	Id                 uint
	Title              string
	Description        string
	ImageURL           string
	FullPrice          float32
	FinalPrice         float32
	PriceModifications float32
}
