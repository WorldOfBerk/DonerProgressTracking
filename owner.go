package main

type Owner struct {
	Storage *Storage
}

func NewOwner(storage *Storage) *Owner {
	return &Owner{
		Storage: storage,
	}
}
