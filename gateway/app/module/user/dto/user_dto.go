package dto

type CreateUpdateUserDto struct {
	Name string `db:"name" json:"name"`
}
