package _1_Singleton

func IncrementAgeClient2() {
	person:= GetInstance()
	person.IncrementAge()
}
