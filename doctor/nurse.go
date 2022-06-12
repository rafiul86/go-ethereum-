package doctor

type Nurse struct {
	Name string
	Age  int
}

func (n *Nurse) Talk() (int){
	return n.Age
}