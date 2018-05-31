package engine

import (
	"github.com/markbest/crawler/utils"
)

type Storage struct {
	client  *utils.ES
	esIndex string
	esType  string
}

func NewStorage(es *utils.ES, esIndex string, esType string) *Storage {
	return &Storage{client: es, esIndex: esIndex, esType: esType}
}

func (s *Storage) Save(item *interface{}) error {
	client := s.client.EsClient
	_, err := client.Index().
		Index(s.esIndex).
		Type(s.esType).
		BodyJson(*item).
		Do()
	if err != nil {
		return err
	}
	return nil
}
