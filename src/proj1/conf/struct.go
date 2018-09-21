package conf

import (
	"fmt"
)

type Mytype struct {
	id   int16
	Flag bool
}

func (c Mytype) Show() {

	fmt.Println(c.id, c.Flag)
}

func (c *Mytype) add(a Mytype) {
	c.id = c.id + a.id
	c.Flag = c.Flag || a.Flag

}

// we can use Bigtype.Show()
type Bigtype struct {
	Mytype
	login string
}

func Test() {
	fmt.Println("testing stuct")
	c := Mytype{
		id:   22,
		Flag: true,
	}

	a := Mytype{
		id:   999,
		Flag: false,
	}

	c.Show()
	c.add(a)
	c.Show()
}
