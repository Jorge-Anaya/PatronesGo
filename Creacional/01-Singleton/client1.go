package _1_Singleton

func IncrementAgeClient1() {
	person:= GetInstance()
	person.IncrementAge()
}
