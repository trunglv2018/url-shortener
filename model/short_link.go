package model

import (
	"url-shortener/helpers/db"
)

type Link struct {
	db.BaseModel `json:",inline"`
	LongLink     string `json:"long_link"`
	Code         string `json:"code"`
}

func (s *Link) Create() error {
	return linkTable.Create(s)
}

func (s *Link) GetByCode(code string) (*Link, error) {
	var links *Link
	return links, linkTable.FindWhere(map[string]string{
		"code": code,
	}, &links)
}

func (s *Link) GetAll() (*Link, error) {
	var links *Link
	return links, linkTable.FindWhere(map[string]string{}, &links)
}
