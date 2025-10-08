package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

// constructor or constructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)

	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}
func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productID {
			return product, nil
		}
	}

	return nil, nil
}
func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}
func (r *productRepo) Delete(productID int) error {
	var tempList []*Product

	for _, p := range r.productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}

	r.productList = tempList

	return nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}
	return &product, nil
}

func generateInitialProducts(r *productRepo) {

	pro1 := &Product{
		ID:          1,
		Title:       "Orange",
		Description: "Fresh Orange",
		Price:       1.50,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQsa3K1PkfEgVzc6JeayE-uGwejpsBDBbVBUw&s",
	}
	pro2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "Green Apple",
		Price:       2.00,
		ImgUrl:      "https://static.vecteezy.com/system/resources/previews/012/086/172/non_2x/green-apple-with-green-leaf-isolated-on-white-background-vector.jpg",
	}
	pro3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "Yellow Banana",
		Price:       0.50,
		ImgUrl:      "https://cdn.jwplayer.com/v2/media/0QcqxCCL/poster.jpg?width=1280",
	}
	pro4 := &Product{
		ID:          4,
		Title:       "Mango",
		Description: "Sweet Mango",
		Price:       1.20,
		ImgUrl:      "https://img.freepik.com/free-photo/mango-still-life_23-2151542114.jpg?w=360",
	}
	pro5 := &Product{
		ID:          5,
		Title:       "Grapes",
		Description: "Fresh Grapes",
		Price:       2.50,
		ImgUrl:      "https://ds.rokomari.store/rokomari110/ProductNew20190903/260X372/Graps_Fruit_Seed_Black_-Non_Brand-96fe8-359969.png",
	}
	pro6 := &Product{
		ID:          6,
		Title:       "Pineapple",
		Description: "Sweet Pineapple",
		Price:       3.00,
		ImgUrl:      "https://www.healthxchange.sg/adobe/dynamicmedia/deliver/dm-aid--c06c2aed-90cf-4360-a423-7f053b2a44d9/pineapple-health-benefits-and-ways-to-enjoy.jpg?preferwebp=true",
	}
	r.productList = append(r.productList, pro1)
	r.productList = append(r.productList, pro2, pro3, pro4, pro5, pro6)

}
