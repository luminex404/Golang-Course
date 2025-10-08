package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"imageUrl" db:"img_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	db *sqlx.DB
}

// constructor or constructor function
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p Product) (*Product, error) {

	fmt.Println("product : ", p)

	query := `
	  INSERT INTO products (
       title,
	   description,
	   price,
	   img_url
     ) VALUES(
     $1,
	 $2,
	 $3,
	 $4
    )
	 RETURNING id
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID)

	if err != nil {
		return nil, err
	}

	return &p, nil

}

func (r *productRepo) Get(id int) (*Product, error) {
	var prd Product

	query := `
	 SEELCT 
	 id ,
	 title,
	 description,
	 price,
	 img_url
	 from products
	 where id = $1
	 `
	err := r.db.Get(&prd, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &prd, nil

}

func (r *productRepo) List() ([]*Product, error) {
	var prdList []*Product

	query := `
	 SELECT 
	  id ,
	  title,
	  description,
	  price,
	  img_url
	  from products
	 `
	err := r.db.Select(&prdList, query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return prdList, nil
}

func (r *productRepo) Update(p Product) (*Product, error) {
	query := `
	 UPDATE products
	   SET 
		title = $2,
		description = $3,
		price = $4,
		img_url = $5
	   WHERE id = $1
	   RETURNING id, title, description, price, img_url
	`
	row := r.db.QueryRow(query, p.ID, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.ImgUrl)

	if err != nil {
		fmt.Println("show the error update not successfull : ", err)
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Delete(id int) error {
	query := `
	 DELETE FROM products WHERE id =$1 
	`
	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
