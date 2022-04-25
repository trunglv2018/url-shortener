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

func ConnectDB(endpoint, uname, password string) {
	var err error
	conn, err = http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{endpoint},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uname, password)
	client, err = driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(uname, password),
	})
	if err != nil {
		log.Fatal(err)
	}
	db, err = client.Database(context.TODO(), "shortenlink")
	if err != nil {
		log.Fatal(err)
	}
	// Root password
	// Base64 encoded CA certificate
	// encodedCA := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURHVENDQWdHZ0F3SUJBZ0lSQU1YZnhmci9lSHR3QUQrMWV1SzlVMlV3RFFZSktvWklodmNOQVFFTEJRQXcKSmpFUk1BOEdBMVVFQ2hNSVFYSmhibWR2UkVJeEVUQVBCZ05WQkFNVENFRnlZVzVuYjBSQ01CNFhEVEl5TURReQpNVEV3TXprd01Wb1hEVEkzTURReU1ERXdNemt3TVZvd0pqRVJNQThHQTFVRUNoTUlRWEpoYm1kdlJFSXhFVEFQCkJnTlZCQU1UQ0VGeVlXNW5iMFJDTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUEKd3BaaTdVUmdpb0NJRzdFOEovdDZ1VzU3cHRyU2FyVE1YTWpycTVwbHU4YVVzd1dNazhGUmFHTjlCVDZvTCthcgpyOGVrQ0VLVjZibFM5K3c1ZXFnZ0tpM2dwVjRxaGxvWG5WQ2NYRXQyeTNOTUtMLzRzeExaa3FtTnFjSFlLVjhnCmpXaXRJM0VOUjJPeXlrNU1URzNCc3FKSUFjUzRLaW1JKzFPUXpDdU1LUkF0cGt0Yk5mcnY2YUo3dGJCVWY5eWoKOVZSOG9JanhVREo0SWI4OUc3b0QvWHk2dVFHUmlqM2MybER4U1NIMENMeWRORE8vM0VqNm9KczI3Wnd2UWVIUQpaeUNBVGZnbWdHYm5MVjFnb3JtNUR4d0s2b1FFdTBGOHZxejRqN1gzRGRjc0wyTW13ekV2aGgxMVJleUJMUjZxCkhReEdwODJsRGlrZzJONmtaK0FYK1FJREFRQUJvMEl3UURBT0JnTlZIUThCQWY4RUJBTUNBcVF3RHdZRFZSMFQKQVFIL0JBVXdBd0VCL3pBZEJnTlZIUTRFRmdRVWx0SGlRZENmUXp5RXFVaEZXS21BNnRSRkVUVXdEUVlKS29aSQpodmNOQVFFTEJRQURnZ0VCQURiQ2VuSUFlZ2grSXVqK01hMjBpMHBndHI2dlJGZnAvQVU2UmlNbUFMemJ4SWtPCjV5bVMwVDRRU0ZGVkhpNDhvL21GbjdzM1JnK2dxbUZUUE9RS1V3QU1hQ2h2OEw5VXM5YkNCREdZdEZRMVFSZlAKT2F6UUUraXpkbW9ZSTF3cDBzenFBc0VLWUxtL3AvS3ZXeUE1aDBvbncwbDlmbGFnOGpCc0xZRVlSUzBXWmtmdAovQlNDYnhVSlIrOHpBOFdMLzZkUFNnNHNaQmlZUmFpUEFhampnUDRyTE5xcWRpWnF3OU9rblZJWWRQWVpPRXlMCitoTmtRcGVVeldOcTZ2NkVvOXJIYUdsWnNNdTVpc082ZzVCZGViNFZRYW1aTVI1bXdKdEdEUjlFdE9UUTFCNmsKTmlHL1hHUGttT256dUJlYm41VHZuYjdXSWl1dy9KTjIvajIzbzRvPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="

	// // Decode CA certificate
	// caCertificate, err := base64.StdEncoding.DecodeString(encodedCA)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Prepare TLS configuration
	// tlsConfig := &tls.Config{}
	// certpool := x509.NewCertPool()
	// if success := certpool.AppendCertsFromPEM(caCertificate); !success {
	// 	log.Fatal("Invalid certificate")
	// }
	// tlsConfig.RootCAs = certpool

	// // Prepare HTTPS connection
	// conn, err = http.NewConnection(http.ConnectionConfig{
	// 	Endpoints: []string{endpoint},
	// 	TLSConfig: tlsConfig,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Create client
	// opts := driver.ClientConfig{
	// 	Connection:     conn,
	// 	Authentication: driver.BasicAuthentication(uname, password),
	// }
	// client, err = driver.NewClient(opts)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // You're now connected
	// // Note that ArangoDB Oasis runs deployments in a cluster configuration.
	// // To achieve the best possible availability, your client application has to handle
	// // connection failures by retrying operations if needed.
	// ctx := context.Background()
	// version, err := client.Version(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Deployment is using version %s\n", version)

}

type Table struct {
	Name   string
	Prefix string
	driver.Collection
}

func NewTable(name string, prefix string) *Table {
	col, err := db.Collection(context.TODO(), name)
	if err != nil {
		log.Fatal(err)
	}
	return &Table{
		Name:       name,
		Prefix:     prefix,
		Collection: col,
	}
}

func (table *Table) Create(model IModel) error {
	model.BeforeCreate(table.Prefix)
	_, err := col.CreateDocument(context.TODO(), model)
	return err
}
