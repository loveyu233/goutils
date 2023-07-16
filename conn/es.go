package conn

import (
	"fmt"
	"github.com/olivere/elastic/v7"
)

func EsConnection(addr, username, password string) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("%s://%s", "http", addr)),
		elastic.SetBasicAuth(username, password),
		elastic.SetSniff(false),
	)
	return client, err
}
