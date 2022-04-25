package model

import "url-shortener/helpers/db"

type ShortLink struct {
	db.BaseModel  `json:",inline"`
	LongLink      string `json:"long_link"`
	ShortLinkCode string `json:"short_link_code"`
}

func (s *ShortLink) Create() error {
	return shortLinkTable.Create(s)
}
