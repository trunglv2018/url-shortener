package model

import (
	"url-shortener/helpers/db"
)

type Link struct {
	db.BaseModel `json:",inline"`
	LongLink     string                 `json:"long_link"`
	VisitCount   int                    `json:"visit_count"`
	Payload      map[string]interface{} `json:"payload"`
}

func (s *Link) Create() error {
	return linkTable.Create(s)
}

func (s *Link) GetByCode(code string) (*Link, error) {
	var links *Link
	return links, linkTable.FindWhere(map[string]string{
		"_key": code,
	}, &links)
}

func (s *Link) GetAll() ([]*Link, error) {
	var links []*Link
	return links, linkTable.FindWhere(map[string]string{}, &links)
}

func (s *Link) Visit() error {
	return linkTable.Increase(s.Key, "visit_count")
}
