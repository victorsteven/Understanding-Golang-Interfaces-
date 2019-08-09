package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type ShopModel interface {
	CountCustomers(time.Time) (int, error)
	CountSales(time.Time) (int, error)
}

// The ShopDb type satisfies our new custom ShopModel interface, because it has the two necessary methods -- CountCustomers() and CountSales()

type ShopDB struct {
	*sql.DB
}

func (sdb *ShopDB) CountCustomers(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM customers WHERE timestamp > $1", since).Scan(&count)
	return count, err
}

func (sdb *ShopDB) CountSales(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM sales WHERE timestamp > $1", since).Scan(&count)

	return count, err
}

func main() {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// shopDB := &ShopDB{db}

	// sr := calculateSalesRate(shopDB)
	// fmt.Printf(sr)
}

// Swap this to use the ShopModel interface type as the parameter, instead of the concrete *ShopDB type

// func calculateSalesRate(sm ShopModel) string {
// 	// since := time.Now().Sub(24 * time.Hour)
// 	since := time.Now().Add(-24 * time.Hour)

// 	sales, err := sm.CountSales(since)

// 	if err != nil {
// 		return "", err
// 	}

// 	customers, err := sm.CountCustomers(since)
// 	if err != nil {
// 		return "", err
// 	}

// 	rate := float64(sales) / float64(customers)

// 	return fmt.Sprintf("%.2f", rate), nil

// }
