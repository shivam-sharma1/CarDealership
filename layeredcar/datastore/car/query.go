package car

var InsertQuery = "INSERT INTO car VALUES (?,?,?,?,?,?)"
var UpdateQuery = "UPDATE car SET name=?,year=?,brand=?,fueltype=? WHERE id=?"
var SelectByIdQuery = "SELECT * FROM car WHERE id= ?"
var SelectByBrandQuery = "SELECT * FROM car WHERE brand=?"
var DeleteQuery = "DELETE FROM car WHERE id= ?"
