package bookstore

type Book struct {
	Title           string
	Author          []string
	Copies          int
	PriceCents      int
	DiscountPercent int
	Featured        bool
	Description     string
}

func (b *Book) NetPrice() float64 {
	discountDecimal := 1 - (float64(b.DiscountPercent) / 100)
	res := float64(b.PriceCents) * discountDecimal
	return res
}

func AddToCatalog(b Book) []Book {
	return append(Catalog, b)
}

func FeaturedBooks() (ret []Book) {
	for _, book := range Catalog {
		if book.Featured {
			ret = append(ret, book)
		}
	}
	return ret
}

type Series struct {
	Books []Book
}

var Catalog = []Book{}
