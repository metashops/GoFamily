package main

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

func (c *Cat) Quack() {
	println(c.Name + " meow")
}

//func main() {
//	var c Duck = &Cat{Name: "draven"}
//	switch c.(type) {
//	case *Cat:
//		cat := c.(*Cat)
//		cat.Quack()
//	}
//}
func main() {
	var c interface{} = &Cat{Name: "draven"}
	switch c.(type) {
	case *Cat:
		cat := c.(*Cat)
		cat.Quack()
	}
}
