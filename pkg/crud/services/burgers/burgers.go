package burgers

import (
	"context"
	"crud/pkg/crud/crudErrors"
	"crud/pkg/crud/models"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type BurgersSvc struct {
	pool *pgxpool.Pool // dependency
}

func NewBurgersSvc(pool *pgxpool.Pool) *BurgersSvc {
	if pool == nil {
		panic(errors.New("pool can't be nil")) // <- be accurate
	}
	return &BurgersSvc{pool: pool}
}

func (service *BurgersSvc) BurgersList() (list []models.Burger, err error) {
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		return nil, crudErrors.NewDbError(err)
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), createTableDDL)
	if err != nil {
		return nil, crudErrors.NewQueryError(createTableDDL, err)
	}
	list = make([]models.Burger, 0)
	conn, err = service.pool.Acquire(context.Background())
	if err != nil {
		return nil, crudErrors.NewDbError(err)
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), selectAllBurgersRemovedFalse)
	if err != nil {
		return nil, crudErrors.NewQueryError(selectAllBurgersRemovedFalse, err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Burger{}
		err := rows.Scan(&item.Id, &item.Name, &item.Price)
		if err != nil {
			return nil, crudErrors.NewDbError(err)
		}
		list = append(list, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, crudErrors.NewDbError(err)
	}

	return list, nil
}

func (service *BurgersSvc) Save(model models.Burger) (err error) {
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		return crudErrors.NewDbError(err)
	}
	defer conn.Release()
	if model.Name == "" {
		return crudErrors.NewModelError("name= not found", err)
	}
	name := model.Name
	if model.Price <= 0 {
		return crudErrors.NewModelError("value= is not more than zero", err)
	}
	price := model.Price
	_, err = conn.Exec(context.Background(), insertAddBurger, name, price)
	if err != nil {
		return crudErrors.NewQueryError(insertAddBurger, err)
	}
	return nil
}

func (service *BurgersSvc) RemoveById(id int) (err error) {
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		return crudErrors.NewDbError(err)
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), updateDeleteBurgerID, id)
	if err != nil {
		return crudErrors.NewQueryError(updateDeleteBurgerID, err)
	}
	return nil
}
