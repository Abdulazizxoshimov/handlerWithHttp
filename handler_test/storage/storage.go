package storage

import (
	"database/sql"
	"fmt"
	"postgresql/handler_test/models"

	_ "github.com/lib/pq"
)



func connect()(*sql.DB, error){
	dsn := "user=postgres password=4444 dbname=nt8 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil{
		return nil, err
	}
	return db, nil
}

func CreteUser(user *models.User)(*models.User, error){
	db, err := connect()
	if err != nil{
		return nil, err
	}

	defer db.Close()

	respUser := models.User{}

	resRow := db.QueryRow("insert into users (id, first_name, last_name) values($1, $2, $3)returning id, first_name, last_name",user.Id, user.FirstName, user.LastName)
	if err := resRow.Scan(&respUser.Id, &respUser.FirstName, &respUser.LastName); err != nil{
		fmt.Print("erro while inserting users", err)
		return nil, err
	}
	return &respUser, nil
}

func GetUser(userId string)( *models.User, error){
	db, err := connect()
	if err != nil{
		return nil, err
	}

	defer db.Close()

    var respuser models.User
	rows := db.QueryRow("select id, first_name, last_name from users where id = $1", userId)
		
		err = rows.Scan(&respuser.Id, &respuser.FirstName, &respuser.LastName)
		if err != nil{
			return nil, err
		}
	
	return &respuser, nil
}

func UpdateUser(userId string, newsName string, newLastName string)( *models.User, error){
	db, err := connect()
	if err != nil{
		return nil, err
	}

	defer db.Close()

    var respuser models.User
	rows := db.QueryRow("update users set  first_name = $1, last_name = $2  where id = $3 returning id, first_name, last_name",newsName, newLastName, userId)
		
		err = rows.Scan(&respuser.Id, &respuser.FirstName, &respuser.LastName)
		if err != nil{
			return nil, err
		}
	
	return &respuser, nil
}

func DeleteUser(userId string,){
	db, err := connect()
	if err != nil{
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec("delete from users where id = $1",userId)
	if err != nil{
		panic(err)
	}
}

func GetAll(page, limit int)([]*models.User, error){
	db, err := connect()
	if err != nil{
		return nil, err
	}

	defer db.Close()

	offset := limit * (page - 1)
	var users []*models.User
	rows, err := db.Query(`select id, first_name, last_name from users limit $1 offset $2`, limit, offset)
	if err != nil {
		panic(err)
	}
	for rows.Next(){
		var user models.User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName)
		if err != nil{
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}