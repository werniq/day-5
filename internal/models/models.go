package models

import (
	"context"
	"database/sql"
	"time"
)

// Db model is the type for db connection
type DBModel struct {
	DB *sql.DB
}


// Models is the wrapper for all models
type Models struct {
	DB DBModel
}

// returns a model type with db connection pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// 
type Widget struct {
	Id 				int `json:"id"`
	Name 			string `json:"name"`
	Description 	string `json:"description"`
	InventoryLevel  int `json:"inventory_level"`
	Price 			int `json:"price"`
	CreatedAt 		time.Time `json:"-"`
	UpdatedAt 		time.Time `json:"-"`
}


func (m *DBModel) GetWidget(id int) (Widget, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second) 
	defer cancel()

	var widget Widget

	row := m.DB.QueryRowContext(ctx, "select id, name from Widgets where id = ?", id)
	err := row.Scan(&widget.Id, &widget.Name)
	if err != nil {
		return widget, err
	}
	return widget, nil
}