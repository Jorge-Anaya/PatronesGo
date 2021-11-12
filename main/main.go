package main

import (
	"fmt"
	"github.com/Jorge-Anaya/PatronesGo/Creacional/01-Singleton"
	"github.com/Jorge-Anaya/PatronesGo/Creacional/02-Factory"
	"github.com/Jorge-Anaya/PatronesGo/Creacional/03-Builder"
	"github.com/Jorge-Anaya/PatronesGo/Creacional/04-Prototype"
	"github.com/Jorge-Anaya/PatronesGo/Estructural/01-Proxy"
	"github.com/Jorge-Anaya/PatronesGo/Estructural/02-Adapter"
	"github.com/Jorge-Anaya/PatronesGo/Estructural/03-Facade"
	"github.com/Jorge-Anaya/PatronesGo/Estructural/04-Bridge"
	"github.com/Jorge-Anaya/PatronesGo/Estructural/05-Composite"
	"github.com/Jorge-Anaya/PatronesGo/Estructural/06-Decorator"
	"github.com/Jorge-Anaya/PatronesGo/Comportamental/01-Observer"
	"os"
	"sync"
)

func main() {

	//*** Creacional
	//testSingleton()
	//testFactory()
	//testBuilder()
	//testPrototype()

	//*** Estructural
	//testProxy()
	//testAdapter()
	//testFacade()
	//testBridge()
	//testComposite()
	//testDecorator()

	//*** Comportamental
	testObserver()
}

func testSingleton() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(200)
	for i := 0; i < 100; i++ {
		go func() {
			defer waitGroup.Done()
			_1_Singleton.IncrementAgeClient1()
		}()
		go func() {
			defer waitGroup.Done()
			_1_Singleton.IncrementAgeClient2()
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

func testPrototype() {
	_4_Prototype.TestPrototype()
}

func testProxy() {
	_1_Proxy.TestProxy()
}

func testAdapter() {
	_2_Adapter.TestAdapter()
}

func testFacade() {
	_3_Facade.TestFacade()
}

func testBridge()  {
	_4_Bridge.TestBridge()
}

func testComposite()  {
	_5_Composite.TestComposite()
}

func testDecorator()  {
	_6_Decorator.TestDecorator()
}

func testObserver()  {
	_1_Observer.TestObserver()
}