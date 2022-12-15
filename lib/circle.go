package circle

type Circle struct {
	str      string
	integger int
}

func New(s string, i int) *Circle {
	return &Circle{s, i}
}
func (c *Circle) Get_str() string {
	return c.str
}
func (c *Circle) Get_integger() int {
	return c.integger
}
