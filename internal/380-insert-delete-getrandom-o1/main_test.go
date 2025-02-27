package _80_insert_delete_getrandom_o1

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	c := Constructor()
	c.Insert(0)
	c.Remove(0)
	c.Insert(-1)
	c.Remove(0)
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
	fmt.Println(c.GetRandom())
}
