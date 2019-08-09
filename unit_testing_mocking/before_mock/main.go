package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type ShopDB struct {
	*sql.DB
}

func (sdb *ShopDB) CountCustomers(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM sales WHERE timestamp > $1", since).Scan(&count)

	return count, err
}

func (sdb *ShopDB) CountSales(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM sales WHERE timestamp > $1", since).Scan(&count)
	return count, err
}

func main() {
	//db, err := sql.Open("postgres", "postgres://user:pass@localhost/db")

	db, err := sql.Open("postgres", "postgres://steven:@localhost/shop")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	shopDB := &ShopDB{db}
	sr, err := calculateSalesRate(shopDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(sr)
}

func calculateSalesRate(sdb *ShopDB) (string, error) {
	// since := time.Now().Sub(24 * time.Hour)
	since := time.Now().Add(-24 * time.Hour)
	sales, err := sdb.CountSales(since)
	if err != nil {
		return "", err
	}
	customers, err := sdb.CountCustomers(since)
	if err != nil {
		return "", err
	}

	rate := float64(sales) / float64(customers)

	return fmt.Sprintf("%.2f", rate), nil
}
