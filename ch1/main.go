package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	syncutil "github.com/kaepa3/goup/ch1/util"
)

func main() {
	pointer()
	field()
	token()
	sync()
	uniq()
	mapping()
	devide()
	fillter()
	reverse()
	slicing()
	typing()
	hex()
	inter()
	cast()
	switcher()
	funcswitcher()
}

type MyError1 struct {
	error
}
type MyError2 struct {
	error
}

func funcswitcher() {
	do := func() error {
		return &MyError1{}
	}
	switch do().(type) {
	case nil:
	case *MyError1:
		fmt.Println("my error1")
	default:
		fmt.Println("derf")
	}

}
func switcher() {
	i := interface{}("Gp expert")
	switch i.(type) {
	case int:
		fmt.Println("integer:", i)
	case string:
		fmt.Println("str:", i)
	default:
		fmt.Println("no name")
	}
}
func cast() {
	var i int32 = 100
	var j int64

	j = int64(i)
	fmt.Println(j)

	msg := "Go expert"
	bs := []byte(msg)
	fmt.Println(bs)
	s := string(bs)
	fmt.Println(s)

	hoge := interface{}("Go expert")
	moji := hoge.(string)
	fmt.Println(moji)

	n, ok := hoge.([]byte)
	fmt.Println(n, ok)
}

type Footstepper interface {
	Footsteps() string
}
type CryFootstepper interface {
	Crier
	Footstepper
}

func inter() {
	var cf CryFootstepper
	cf = &Person{}
	fmt.Println(cf.Cry(), cf.Footsteps())
	cf = &PartyPeople{}
	fmt.Println(cf.Cry(), cf.Footsteps())
}

type Person struct{}

func (p *Person) Cry() string {
	return "Hi"
}
func (p *Person) Footsteps() string {
	return "PitaPat"
}

type PartyPeople struct {
	Person
}

func (p *PartyPeople) Cry() string {
	return "Sup?"
}

type Crier interface {
	Cry() string
}

func hex() {
	fv := Hex(1024).String
	fmt.Println(fv())
	fe := Hex.String
	fmt.Println(fe(1024))
}

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))

}
func pointer() {
	u := &User{
		Name: "Rech",
		Age:  33,
	}
	u.Aging()
	fmt.Println(u.Age)
	u.AgingBut()
	fmt.Println(u.Age)

}

type User struct {
	Name string
	Age  int
}

func (u *User) Aging() {
	u.Age++
}
func (u User) AgingBut() {
	u.Age++
}

type Chip struct {
	Number int
}

type Card struct {
	string
	Chip
	Number int
}

func (c *Chip) Scan() {
	fmt.Println(c.Number)
}

func field() {
	c := Card{
		string: "Creadit",
		Chip: Chip{
			Number: 2,
		},
		Number: 1,
	}
	c.Scan()

}
func token() {
	s := syncutil.NewSecret()
	err := json.NewEncoder(os.Stdout).Encode(s)
	if err != nil {
		fmt.Println("fail:", err)
	}
}
func sync() {
	c := &syncutil.Counter{
		Name: "Access",
	}
	fmt.Println(c.Increment())
	fmt.Println(c.View())
}
func uniq() {
	followers := []string{"Jo", "hRE", "Jo"}
	uniq := make([]string, 0, len(followers))

	m := make(map[string]struct{})
	for _, v := range followers {
		if _, ok := m[v]; ok {
			continue
		}
		uniq = append(uniq, v)
		m[v] = struct{}{}
	}
}
func mapping() {
	empty := map[string]int{}
	maker := make(map[string]int, 10)
	fmt.Println(empty, maker)
}
func devide() {
	src := []int{1, 2, 3, 4, 5}
	size := 2
	dst := make([][]int, 0, (len(src)+size-1)/size)

	for size < len(src) {
		src, dst = src[size:], append(dst, src[0:size:size])
	}
	dst = append(dst, src)

	fmt.Println(dst)
}
func fillter() {
	src := []int{1, 2, 3, 4, 5}

	dst := src[:0]
	for _, v := range src {
		if even(v) {
			dst = append(dst, v)
		}
	}
	fmt.Println(dst)

	for i := len(dst); i < len(src); i++ {
		src[i] = 0
	}
}
func even(n int) bool {
	return n%2 == 0
}
func reverse() {
	src := []int{1, 2, 3, 4, 5}
	for i := len(src)/2 - 1; i >= 0; i-- {
		opp := len(src) - 1 - i
		src[i], src[opp] = src[opp], src[i]
	}
	fmt.Println(src)
	for left, right := 0, len(src)-1; left < right; left, right = left+1, right-1 {
		src[left], src[right] = src[right], src[left]
	}
	fmt.Println(src)

}
func slicing() {
	src := []int{1, 2, 3, 4, 5}
	dst := append(src[:2], src[3:]...)
	fmt.Println(src)
	fmt.Println(dst)

	src = []int{1, 2, 3, 4, 5}
	dst = src[:2+copy(src[2:], src[3:])]
	fmt.Println(dst)
	fmt.Println(src)
}
func typing() {
	type MyDuration time.Duration
	d := MyDuration(100)
	fmt.Printf("%T", d)

	td := time.Duration(d)
	md := 100 * d

	fmt.Printf("td: %T. md: %T", td, md)

}
