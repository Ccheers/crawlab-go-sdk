package interfaces

import "github.com/Ccheers/crawlab-go-sdk/entity"

type ResultService interface {
	SaveItem(item ...entity.Result)
	SaveItems(item []entity.Result)
}
