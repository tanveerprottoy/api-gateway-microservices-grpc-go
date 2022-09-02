package dto

type CreateUpdateUserBody struct {
	Name string `db:"name" json:"name"`
}
