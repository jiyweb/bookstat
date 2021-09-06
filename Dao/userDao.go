package Dao

import (
	"bookstat/model"
	"bookstat/utils"
)

func CheckUserNameAndPassword(username string, password string) (*model.User, error) {

	sql := "select id,username,password,email from users where username=?and password=?"
	row := utils.Db.QueryRow(sql, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)

	return user, nil
}

//查询用户名
func CheckUserName(username string) (*model.User, error) {

	sql := "select id,username,password,email from users where username=?"
	row := utils.Db.QueryRow(sql, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)

	return user, nil
}

//保存添加
func SaveUser(username, password, email string) error {

	sql := "insert into users(username,password,email)values(?,?,?)"

	_, err := utils.Db.Exec(sql, username, password, email)

	if err != nil {

		return err

	}
	return nil
}
