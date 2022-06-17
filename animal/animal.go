package animal

type Dog struct {
	Sound string
	Leg   int
}

type Cat struct {
	Sound string
	Leg   int
}

func (d *Dog) Bark() {
	println(d.Sound)
}

func (c *Cat) Mew() {
	println(c.Sound)
}
