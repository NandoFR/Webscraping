package data

var CellphoneList = []string{
	"iPhone", "iPhone 3G", "iPhone 3GS", "iPhone 4", "iPhone 4s", "iPhone 5", "iPhone 5c", "iPhone 5s", "iPhone 6", "iPhone 6", "iPhone 6 Plus", "iPhone 6s", "iPhone 6s Plus", "iPhone SE 1ª Geração", "iPhone SE 2ª Geração", "iPhone SE 3ª Geração", "iPhone 7", "iPhone 7 Plus", "iPhone 8", "iPhone 8 Plus", "iPhone X", "iPhone XR", "iPhone XS", "iPhone XS Max", "iPhone 11", "iPhone 11 Pro", "iPhone 11 Pro Max", "iPhone 12", "iPhone 12 mini", "iPhone 12 Pro", "iPhone 12 Pro Max", "iPhone 13", "iPhone 13 mini", "iPhone 13 Pro", "iPhone 13 Pro Max", "iPhone 14", "iPhone 14 Plus", "iPhone 14 Pro", "iPhone 14 Pro Max",
}

var CategoryList = map[string]string{
	"Celular": "Celular",
}
var ProductList = map[string][]string{
	CategoryList["Celular"]: CellphoneList,
}
