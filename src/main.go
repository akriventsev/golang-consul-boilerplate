package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hashicorp/consul/api"
)

var (
	port = 4456
	name = "fvd"
)

func init() {
	name = os.Getenv("SERVICE_NAME")
}

func hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}
func main() {

	registration := api.AgentServiceRegistration{
		ID:   fmt.Sprintf("id-%s-%d", name, port),
		Name: fmt.Sprintf("service-%s", name),
		Tags: []string{"tag4", "tag6"},
		Meta: map[string]string{"karamba": "first"},
	}

	address := hostname()
	registration.Address = address
	registration.Port = port
	registration.Check = new(api.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/healthcheck",
		address, port)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"

	config := api.DefaultConfig()
	config.Address = os.Getenv("CONSUL_SERVER_ADDRESS")
	log.Printf("config: %s\n", config.Address)
	// Get a new Consul client
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	agent := client.Agent()

	err = agent.ServiceRegister(&registration)

	http.HandleFunc("/healthcheck", homeRouterHandler)       // установим роутер
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil) // задаем  слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func homeRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Consul!") // отправляем данные на клиентскую сторону
}
