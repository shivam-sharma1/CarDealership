package engine

var SelectByIdQuery = "SELECT * FROM engine WHERE engineid= ?"
var InsertQuery = "INSERT INTO engine VALUES (?,?,?,?)"
var UpdateQuery = "UPDATE engine SET displacement=?,noc=?,rng=? WHERE engineid=?"
var DeleteQuery = "DELETE FROM engine WHERE engineid= ?"
