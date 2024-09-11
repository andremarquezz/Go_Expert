package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// product1 := NewProduct("Iphone 14 Pro Max", 3650.90)
	// err = insertProduct(db, product1)
	// if err != nil {
	// 	panic(err)
	// }

	// product2 := NewProduct("Samsungzin Pro Max", 3650.90)
	// err = insertProduct(db, product2)
	// if err != nil {
	// 	panic(err)
	// }

	// product2.Price = 4000.00
	// err = updateProduct(db, product2)
	// if err != nil {
	// 	panic(err)
	// }

	// p, err := selectOneProduct(db, product2.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Product: %v, possui o preço de: %v\n", p.Name, p.Price)

	// products, err := selectAllProducts(db)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, p := range products {
	// 	fmt.Printf("Product: %v, possui o preço de: %v\n", p.Name, p.Price)
	// }

	err = deleteProduct(db, "db198fb5-e07c-4b41-9ad1-284ffd594ac9")
	if err != nil {
		panic(err)
	}

}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func insertProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.ID, p.Name, p.Price)
	return err
}

func updateProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Name, p.Price, p.ID)
	return nil
}

func selectOneProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	p := &Product{}
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		p := Product{}
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
