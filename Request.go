package gmodel

type FieldValue[CType any] struct {
	Field *CType `json:"value" validate:"required"`
}
