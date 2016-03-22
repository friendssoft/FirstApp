package main
import (
"database/sql"

)
func main() {

db, err := sql.Open("mysql", "root:@/ravikiran")
if err != nil {
     panic(err.Error())
}
defer db.Close()
stmtIns, err :=db.Prepare("INSERT INTO go_ex VALUES(?, ?)")

for i :=0;i<25;i++ {
    _, err = stmtIns.Exec(i, (i*i))
}
if err != nil {
     panic(err.Error())
}

}