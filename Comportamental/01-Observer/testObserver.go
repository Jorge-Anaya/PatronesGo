package _1_Observer

func TestObserver() {
	m := Message{}
	nameObs := "slack"
	obs := observerFactory(nameObs)
	m.AddObserver(nameObs, obs)
	nameObs1 := "email"
	obs1 := observerFactory(nameObs1)
	m.AddObserver(nameObs1, obs1)
	m.Msg = "My messageeeeee X"
	m.NotifyObservers()
}

func observerFactory(name string) Observer {
	switch name {
	case "slack":
		return &Slack{}
	case "email":
		return &Email{}
	}

	return nil
}
