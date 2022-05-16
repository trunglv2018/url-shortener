package db

import (
	"context"
	"fmt"
	"log"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

var client driver.Client
var conn driver.Connection
var db driver.Database
var col driver.Collection

func ConnectDB(endpoint, uname, password, dbName string) {
	var err error
	conn, err = http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{endpoint},
	})
	if err != nil {
		log.Fatal(err)
	}
	client, err = driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(uname, password),
	})
	if err != nil {
		log.Fatal(err)
	}
	db, err = client.Database(context.TODO(), dbName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect to arrangodb successfully")

}

type Table struct {
	Name   string
	Prefix string
	driver.Collection
}

const ErrCollectionNotFound = "collection or view not found"

func NewTable(name string, prefix string) *Table {
	// db.CreateCollection(context.TODO(), name, &driver.CreateCollectionOptions{
	// 	KeyOptions: &driver.CollectionKeyOptions{
	// 		Type: &driver.KeyGeneratorAutoIncrement,
	// 	},
	// })
	col, err := db.Collection(context.TODO(), name)
	if err != nil {
		if err.Error() == ErrCollectionNotFound {
			col, err = db.CreateCollection(context.TODO(), name, &driver.CreateCollectionOptions{
				KeyOptions: &driver.CollectionKeyOptions{
					Type: driver.KeyGeneratorAutoIncrement,
				},
			})
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}

	}
	return &Table{
		Name:       name,
		Prefix:     prefix,
		Collection: col,
	}
}

func (table *Table) Create(model IModel) error {
	model.BeforeCreate(table.Prefix)
	_, err := table.Collection.CreateDocument(context.TODO(), model)
	return err
}

func (table *Table) FindWhere(filter map[string]string, result interface{}) error {
	var filterStr = ""
	if len(filter) > 0 {
		filterStr = " FILTER "
		for key, value := range filter {
			filterStr += "d." + key + " == " + "'" + value + "'"
		}
	}
	query := fmt.Sprintf("FOR d IN %s %s  RETURN d", table.Name, filterStr)
	fmt.Println(query)
	cursor, err := db.Query(context.TODO(), query, nil)
	if err != nil {
		return err
	}
	_, err = cursor.ReadDocument(context.TODO(), &result)
	if err != nil {
		return err
	}
	defer cursor.Close()
	return nil
}

func (table *Table) Increase(docKey, field string) error {
	var let = fmt.Sprintf(`
		LET doc = DOCUMENT("%s/%s")
	`, table.Name, docKey)
	var update = fmt.Sprintf(`
		UPDATE doc WITH {%s: doc.%s + 1 } IN %s
	`, field, field, table.Name)
	var query = fmt.Sprintf(`
		%s
		%s
	`, let, update)
	_, err := db.Query(context.TODO(), query, nil)
	return err

}
