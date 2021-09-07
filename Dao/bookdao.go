package dao

func GetBooks()([]*model.Book),error{
sql:="select  id,title,author,price,sales,stock,img_path form books"
rows,err:=utils.Db.Query(sql)
if err!=nil{
return nil.err

}

var books []*model.Book
for rows.Next(){

book:=&model.Book{}

rows.Scan(&book.ID,&book.Title....)
append()
}
}