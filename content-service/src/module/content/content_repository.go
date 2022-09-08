package content

import (
	"fmt"
	"log"
	"txp/contentservice/src/data"
	"txp/contentservice/src/module/content/proto"
)

type ContentRepository struct {
}

func (r *ContentRepository) Create(
	u *proto.Content,
) (*string, error) {
	var lastId *string = nil
	stmt, err := data.Db.PrepareNamed(
		"INSERT INTO Contents (name)" +
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

func (r *ContentRepository) ReadMany() (
	*proto.Contents,
	error,
) {
	Contents := &proto.Contents{}
	err := data.Db.Select(
		&Contents,
		"SELECT * FROM Contents WHERE id > $1",
		fmt.Sprintf("%d", 0),
	)
	return Contents, err
}

func (r *ContentRepository) ReadOne(id string) (
	*proto.Content,
	error,
) {
	u := &proto.Content{}
	err := data.Db.Get(
		&u,
		"SELECT * FROM Contents WHERE id = $1 LIMIT 1",
		id,
	)
	return u, err
}

func (r *ContentRepository) Update(
	id string,
	u *proto.Content,
) (int64, error) {
	q := "UPDATE Contents SET name=:name WHERE id = " + id
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

func (r *ContentRepository) Delete(id string) (
	int64,
	error,
) {
	q := "DELETE FROM Contents WHERE id = $1"
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
