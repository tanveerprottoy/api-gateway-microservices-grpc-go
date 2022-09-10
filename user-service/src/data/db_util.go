package data

import (
	"database/sql"
	"fmt"
	"log"
)

func GetRowsAffected(result sql.Result) int64 {
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	return rows
}

func GetEntities[T any](
	rows *sql.Rows,
	entity *T,
) ([]*T, error) {
	var entities []*T
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		if err := rows.Scan(&entity); err != nil {
			return nil, fmt.Errorf("GetEntities %v", err)
		}
		entities = append(entities, entity)
	}
	return entities, nil
}
