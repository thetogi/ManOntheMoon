package models

type Model interface {
	Create()
	Retrieve()
	RetrieveAll()
}
