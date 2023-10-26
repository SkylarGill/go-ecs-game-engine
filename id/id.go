package id

var currentId uint64 = 0

type HasId struct {
	Id uint64
}

func GetNew() uint64 {
	newId := currentId
	currentId++
	return newId
}
