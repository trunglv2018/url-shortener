package model

import "url-shortener/helpers/db"

var linkTable = db.NewTable("link", "sl")
var logTable = db.NewTable("log", "lg")
