package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func Store(p Product) Product{
	p.ID = len(productList)+1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productID int) *Product{

      for _, product := range productList {
		if product.ID == productID {
			 return &product
		}
	}

	return nil
}

func Update(product Product){
	for idx,p := range productList{
		if p.ID == product.ID{
			productList[idx]=product
		}
	}
}

func Delete(productID int){
	var tempList []Product

	for _,p := range productList{
		if p.ID != productID{
			tempList = append(tempList, p)
		}
	}

	productList=tempList
}

func init() {

	pro1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Fresh Orange",
		Price:       1.50,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQsa3K1PkfEgVzc6JeayE-uGwejpsBDBbVBUw&s",
	}
	pro2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Green Apple",
		Price:       2.00,
		ImgUrl:      "https://static.vecteezy.com/system/resources/previews/012/086/172/non_2x/green-apple-with-green-leaf-isolated-on-white-background-vector.jpg",
	}

	pro3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Yellow Banana",
		Price:       0.50,
		ImgUrl:      "https://cdn.jwplayer.com/v2/media/0QcqxCCL/poster.jpg?width=1280",
	}
	pro4 := Product{
		ID:         4,
		Title:       "Mango",
		Description: "Sweet Mango",
		Price:       1.20,
		ImgUrl:      "https://img.freepik.com/free-photo/mango-still-life_23-2151542114.jpg?w=360",
	}
	pro5 := Product{
		ID:          5,
		Title:       "Grapes",
		Description: "Fresh Grapes",
		Price:       2.50,
		ImgUrl:      "https://ds.rokomari.store/rokomari110/ProductNew20190903/260X372/Graps_Fruit_Seed_Black_-Non_Brand-96fe8-359969.png",
	}
	pro6 := Product{
		ID:          6,
		Title:       "Pineapple",
		Description: "Sweet Pineapple",
		Price:       3.00,
		ImgUrl:      "https://www.healthxchange.sg/adobe/dynamicmedia/deliver/dm-aid--c06c2aed-90cf-4360-a423-7f053b2a44d9/pineapple-health-benefits-and-ways-to-enjoy.jpg?preferwebp=true",
	}
	productList = append(productList, pro1)
	productList = append(productList, pro2, pro3, pro4, pro5, pro6)

 
}
