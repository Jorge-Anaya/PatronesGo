package _2_Factory

func Factory(typeConnection int) DBConnection {
	switch typeConnection {
	case 1:
		return &MySql{}
	case 2:
		return &Postgres{}
	default:
		return nil
	}
}
