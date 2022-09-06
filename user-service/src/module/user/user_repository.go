package user

import (
	"fmt"
	"log"
	"time"
	"txp/userservice/src/data"
	"txp/userservice/src/module/user/dto"
	"txp/userservice/src/module/user/entity"
)

type UserRepository struct {
}

func (r *UserRepository) Create(d *dto.CreateUpdateUserBody) (
	int,
	error,
) {
	lastId := -1
	stmt, err := data.Db.PrepareNamed(
		"INSERT INTO users (name)" +
			"VALUES (:name) RETURNING id",
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	err = stmt.Get(&lastId, d)
	if err != nil {
		log.Println(err)
	}
	err = stmt.Close()
	if err != nil {
		log.Println(err)
	}
	return lastId, nil
}

func (r *UserRepository) FindMany() (
	[]entity.User,
	error,
) {
	users := []entity.User{}
	err := data.Db.Select(
		&users,
		"SELECT * FROM users WHERE id > $1",
		fmt.Sprintf("%d", 0),
	)
	if err != nil {
		log.Println(err)
		return users, err
	}
	return users, nil
}

func (r *UserRepository) FindOne(id string) (
	entity.User,
	error,
) {
	u := entity.User{}
	err := data.Db.Get(
		&u,
		"SELECT * FROM users WHERE id = $1 LIMIT 1",
		id,
	)
	return u, err
}

func (r *UserRepository) Update(
	id string,
	u *entity.User,
) (int64, error) {
	u.UpdatedAt = time.Now()
	q := "UPDATE users SET name=:name," +
		" updated_at = :updated_at WHERE id = " + id
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
