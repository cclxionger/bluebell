package mysql

import (
	"bluebell/global"
	"bluebell/models"
	"database/sql"
	"errors"
)

func CheckUserExist(username string) (bool, error) {
	sqlStr := "select count(user_id)  from `user` where username = ?"
	var count int
	err := global.DB.QueryRow(sqlStr, username).Scan(&count)
	if err != nil {
		// 没有查询到错误
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return count > 0, nil
}
func InsertUser(u *models.User) error {
	sqlStr := "INSERT INTO `user` (user_id,username,password) VALUES (?,?,?)"
	var count int64
	res, err := global.DB.Exec(sqlStr, u.UserID, u.Username, u.Password)
	if err != nil {
		return err
	}
	count, err = res.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		return errors.New("影响数据不为一条，插入有误")
	}
	return nil
}
func QueryPassword(u *models.User) (string, error) {
	var encryPassword string
	sqlStr := "select user_id,password  from `user` where username = ?"
	err := global.DB.QueryRow(sqlStr, u.Username).Scan(&u.UserID, &encryPassword)
	if err != nil {
		return "", err
	}
	return encryPassword, nil
}
func GetUserDetail(id int) (*models.UserDetail, error) {
	ud := new(models.UserDetail)
	sqlStr := `select user_id, username from user where user_id = ?`
	err := global.DB.QueryRow(sqlStr, id).Scan(&ud.UserID, &ud.Username)
	return ud, err
}
