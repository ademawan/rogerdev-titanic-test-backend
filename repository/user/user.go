package user

import (
	"database/sql"
	"errors"
	"fmt"
	"rogerdev-titanic-test-backend/entities"
	"rogerdev-titanic-test-backend/middlewares"
	"strings"
	"time"

	"github.com/lithammer/shortuuid"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Register(user entities.User) (entities.User, error) {

	user.Password, _ = middlewares.HashPassword(user.Password)
	uid := shortuuid.New()
	user.UserUid = uid
	_, err := ur.db.Exec("insert into user (uid,name,email,password,address,gender,created_at,updated_at,deleted_at) values (?, ?, ?, ?,?,?,?,?,?)", user.UserUid, user.Name, user.Email, user.Password, user.Address, user.Gender, user.CreatedAt.(int64), user.UpdatedAt.(int64), user.DeletedAt.(int64))
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) GetByUid(userUid string) (entities.User, error) {
	user := entities.User{}

	err := ur.db.QueryRow("SELECT * FROM user WHERE uid=?", userUid).Scan(&user.UserUid, &user.Name, &user.Email, &user.Password, &user.Address, &user.Gender, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return user, err
	}
	if user.UserUid == "" {
		return user, errors.New("record not found")
	}

	return user, nil
}

func (ur *UserRepository) Update(userUid string, newUser entities.User) (entities.User, error) {

	var user entities.User

	queryPrepareField := make([]string, 0)
	queryExec := make([]interface{}, 0)
	if newUser.Name != "" {
		queryPrepareField = append(queryPrepareField, "name")
		queryExec = append(queryExec, newUser.Name)
	}
	if newUser.Email != "" {
		queryPrepareField = append(queryPrepareField, "email")
		queryExec = append(queryExec, newUser.Email)
	}
	if newUser.Address != "" {
		queryPrepareField = append(queryPrepareField, "address")
		queryExec = append(queryExec, newUser.Address)
	}
	if newUser.Gender != "" {
		queryPrepareField = append(queryPrepareField, "gender")
		queryExec = append(queryExec, newUser.Gender)
	}
	timeLocation, _ := time.LoadLocation("Asia/Jakarta")
	timeNow := time.Now().In(timeLocation).Unix()
	queryPrepareField = append(queryPrepareField, "updated_at")
	queryExec = append(queryExec, timeNow)
	queryExec = append(queryExec, userUid)

	queryPrepareString := strings.Join(queryPrepareField, "= ? ,")
	queryPrepareString += "= ? "

	stmt, err := ur.db.Prepare("UPDATE user SET " + queryPrepareString + " where uid=?")

	if err != nil {
		return user, err
	}
	// execute
	if _, err := stmt.Exec(queryExec...); err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Delete(userUid string) error {

	_, err := ur.db.Exec("delete from user where uid = ?", userUid)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil

}
