package model

import "url-shortener/helpers/db"

var shortLinkTable = db.NewTable("shorten_link", "sl")
