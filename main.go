package main

func main() {
	p1 := CreatePersonStack()
	p2 := CreatePersonHeap()

	IncreaseAge(&p1, 10)
	IncreaseAge(p2, 15)

	/*fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("p1 = %p\np2 = %p", p1, p2)*/
}

type Person struct {
	Name string
	Age  int
}

func CreatePersonStack() Person {
	pStack := Person{Name: "Auber", Age: 35}
	return pStack
}

func CreatePersonHeap() *Person {
	pHeap := Person{Name: "Auber", Age: 35}
	return &pHeap
}

func IncreaseAge(p *Person, years int) {
	p.Age = p.Age + years
}
