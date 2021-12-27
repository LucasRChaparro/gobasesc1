package productos

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

type Repository interface {
	GetAll() Product
	Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() Product {
	producto1 := Product{
		ID:    25,
		Name:  "hgjhg",
		Type:  "100",
		Count: 122,
		Price: 12.45,
	}
	return producto1
}
