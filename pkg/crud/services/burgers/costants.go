package burgers

const insertAddBurger = "INSERT INTO burgers(name, price) VALUES ($1, $2)"
const updateDeleteBurgerID = `UPDATE burgers SET removed=TRUE WHERE id=$1`
const selectAllBurgersRemovedFalse = "SELECT id, name, price FROM burgers WHERE removed = FALSE"
const createTableDDL = `CREATE TABLE  IF NOT EXISTS  burgers (
id BIGSERIAL PRIMARY KEY,
name TEXT NOT NULL,
price INT NOT NULL CHECK ( price > 0 ),
removed BOOLEAN DEFAULT FALSE
);`