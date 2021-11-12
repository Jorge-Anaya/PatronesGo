package _6_Decorator

func TestDecorator() {
	route := NewRoute()
	start(&route)
	route.Exec("/saludar")
}

func start(route *Route) {
	route.Add(NewLogRegistry(NewHandlerSayHello()), "/saludar")
	route.Add(NewLogRegistry(NewHandlerSayGoodbye()), "/despedirse")
}
