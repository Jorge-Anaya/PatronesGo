package main

import (
	"fmt"
	"github.com/Jorge-Anaya/PatronesGo/Creacionales/01-Singleton"
	"github.com/Jorge-Anaya/PatronesGo/Creacionales/01-Singleton/client"
	"github.com/Jorge-Anaya/PatronesGo/Creacionales/02-Factory"
	"github.com/Jorge-Anaya/PatronesGo/Creacionales/03-Builder"
	"os"
	"sync"
)

func main() {
	//testSingleton()
	//testFactory()
	testBuilder()
}

func testSingleton() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(200)
	for i := 0; i < 100; i++ {
		go func() {
			defer waitGroup.Done()
			client.IncrementAgeClient1()
		}()
		go func() {
			defer waitGroup.Done()
			client.IncrementAgeClient2()
		}()
	}
	waitGroup.Wait()
	person := _1_Singleton.GetInstance()
	fmt.Println(person.GetAge())
}

func testFactory() {
	connection := _2_Factory.Factory(1)
	if connection == nil {
		fmt.Println("Motor no valido")
		os.Exit(1)
	}

	err := connection.Connection()
	if err != nil {
		fmt.Printf("no se pudo crear la conexión %v", err)
		os.Exit(1)
	}

	now, err := connection.GetNow()
	if err != nil {
		fmt.Printf("no se pudo consultar la fecha: %v", err)
		os.Exit(1)
	}

	fmt.Println(now.Format("2006-01-02"))

	err = connection.Close()
	if err != nil {
		fmt.Printf("No se pudo cerrar la conexión %v", err)
	}
}

func testBuilder() {
	json := &_3_Builder.JSONMessageBuilder{}
	xml := &_3_Builder.XMLMessageBuilder{}
	director := &_3_Builder.Director{}
	director.SetBuilder(json)
	jsonMsg, err := director.BuildMessage("Gopher", "Hola mundo!")
	if err != nil {
		fmt.Printf("BuildMesage(): JSON: no se esperaba error, se recibió: %v\n", err)
	}
	fmt.Printf(string(jsonMsg.Body) + "\n")
	director.SetBuilder(xml)
	xmlMsg, err := director.BuildMessage("Gopher", "Hola mundo")
	if err != nil {
		fmt.Printf("BuildMesage(): XML: no se esperaba error, se recibió: %v", err)
	}
	fmt.Printf(string(xmlMsg.Body))
}
