package main

import (
	"configServices/controller"
	"configServices/domain"
	"configServices/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	config := domain.Config{}
	//data ,err := ioutil.ReadFile("config.yaml")
	//if err != nil {
	//	panic(err)
	//}
	//err = config.SetFromBytes(data)
	//if err != nil {
	//	panic(err)
	//}
	//reService, err := config.Get("service.registry.device")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(reService)
	Conservice := service.ConfigService{
		Config:   &config,
		Location: "config.yaml",
	}

	go Conservice.Watch(3 * time.Second)

	c := controller.Controller{
		Config: &config,
	}
	router := mux.NewRouter()
	//router.Get("/read/{serviceName}",c.ReadConfig())
	router.HandleFunc("/read/{serviceName}", c.ReadConfig).Methods("GET")
	//router.HandleFunc("/read/{serviceName}", c.ReadConfig)
	log.Fatal(http.ListenAndServe(":8000", router))

}
