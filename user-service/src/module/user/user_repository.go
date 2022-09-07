package user

import (
	"fmt"
	"log"
	"txp/userservice/src/data"
	"txp/userservice/src/module/user/proto"
)

type UserRepository struct {
}

func (r *UserRepository) Create(
	u *proto.User,
) (*string, error) {
	var lastId *string = nil
	stmt, err := data.Db.PrepareNamed(
		"INSERT INTO users (name)" +
			"VALUES (:name) RETURNING id",
	)
	if err != nil {
		log.Println(err)
		return lastId, err
	}
	err = stmt.Get(&lastId, u)
	if err != nil {
		log.Println(err)
	}
	err = stmt.Close()
	if err != nil {
		log.Println(err)
	}
	return lastId, nil
}

func (r *UserRepository) ReadMany() (
	*proto.Users,
	error,
) {
	users := &proto.Users{}
	err := data.Db.Select(
		&users,
		"SELECT * FROM users WHERE id > $1",
		fmt.Sprintf("%d", 0),
	)
	return users, err
}

func (r *UserRepository) ReadOne(id string) (
	*proto.User,
	error,
) {
	u := &proto.User{}
	err := data.Db.Get(
		&u,
		"SELECT * FROM users WHERE id = $1 LIMIT 1",
		id,
	)
	return u, err
}

func (r *UserRepository) Update(
	id string,
	u *proto.User,
) (int64, error) {
	q := "UPDATE users SET name=:name WHERE id = " + id
	res, err := data.Db.NamedExec(
		q,
		u,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return data.GetRowsAffected(res), nil
}

func (r *UserRepository) Delete(id string) (
	int64,
	error,
) {
	q := "DELETE FROM users WHERE id = $1"
	res, err := data.Db.Exec(
		q,
		id,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return data.GetRowsAffected(res), nil
}
