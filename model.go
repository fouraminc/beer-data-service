package main

import "database/sql"

type beer struct {

	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

const (
	createBeerQuery = "INSERT INTO beer (name, description) VALUES($1, $2) RETURNING beer_id"
	getBeerQuery = "SELECT name, description FROM beer WHERE beer_id = $1"
	getBeersQuery = "SELECT beer_id, name, description from beer"
	deleteBeerQuery = "DELETE from beer where beer_id = $1"
	updateBeerQuery = "UPDATE beer SET name = $1, description = $2 WHERE beer_id = $3"
)

func (b *beer) getBeer(db *sql.DB) error {
	return db.QueryRow(getBeerQuery, b.ID).Scan(&b.Name, &b.Description)
}

func (b *beer) createBeer(db *sql.DB) error {
	err := db.QueryRow(createBeerQuery, &b.Name, &b.Description).Scan(&b.ID)

	if err != nil {
		return err
	}

	return nil
}

func getBeers(db *sql.DB) ([]beer, error) {
	rows, err := db.Query(getBeersQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var beers []beer

	for rows.Next() {
		var b beer
		if err := rows.Scan(&b.ID, &b.Name, &b.Description); err != nil {
			return nil, err
		}
		beers = append(beers, b)
	}
	return beers, nil

}

func (b *beer) deleteBeer(db *sql.DB) error {

	_, err := db.Exec(deleteBeerQuery, b.ID)
	if err != nil {
		return err
	}

	return nil
}

func (b *beer) updateBeer(db *sql.DB) error {
	_, err := db.Exec(updateBeerQuery, b.Name, b.Description, b.ID)
	return err
}