package burgers

const insertAddBurger = "INSERT INTO burgers(name, price) VALUES ($1, $2)"
const updateDeleteBurgerID = `UPDATE burgers SET removed=TRUE WHERE id=$1`
const selectAllBurgersRemovedFalse = "SELECT id, name, price FROM burgers WHERE removed = FALSE"
