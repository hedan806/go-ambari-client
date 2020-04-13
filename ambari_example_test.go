package ambari_test

import (
	"github.com/hedan806/go-ambari-client"
	"github.com/hedan806/go-ambari-client/transport"
	"log"
	"os"
	"testing"
)

func ExampleNewDefaultClient() {
	es, err := ambari.NewClient(ambari.Config{
		Address:           "http://10.20.1.21:8001",
		Username:          "admin",
		Password:          "admin",
		EnableDebugLogger: true,
		Logger:            &transport.ColorLogger{Output: os.Stdout},
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s\n", err)
	}

	res, err := es.Cluster.Health()
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}
	defer res.Body.Close()

	log.Print(res.StatusCode)
	log.Print(res.String())

	log.Print(es.Transport.(*transport.Client).URLs())
}

func TestClusterHealth(t *testing.T) {
	ExampleNewDefaultClient()
}
