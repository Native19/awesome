package main

import (
	"fmt"
	"os"
	"runtime"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {

	/*	con := true
		for i := 0; con; {
			i++
			fmt.Println(i)
			if i > 20 {
				con = !con
			}
		}

		var j interface{}
		j = "hi"

		switch j.(type) {
		case string:
			fmt.Println("hi")
		case int:
			fmt.Println("123")
		default:
			fmt.Println("hello")
		}*/

	/*	slice1 := []int{6, 1, 2}
		slice2 := []int{9, 3}

		var mySlice = make([]int, 2, 5)

		for i := 10; i > 7; i-- {
			mySlice = append(mySlice, i)
		}
		//mySlice[5] = 10
		fmt.Println(mySlice, len(mySlice), cap(mySlice))

		mySlice = mySlice[2:3]
		fmt.Println(mySlice, len(mySlice), cap(mySlice))

		mySlice = append(mySlice, 3)
		fmt.Println(mySlice, len(mySlice), cap(mySlice))

		for item := range mySlice {
			fmt.Println(item)
		}

		var k []int
		k = append(k, 3)
		fmt.Println(k, len(k), cap(k))*/

	/*	t := make(map[int]string)
		t[20] = "Hello"
		t[66] = " World"
		//t[1] = " !"*/
	/*for i := 0; i < 100; i++ {
		t[i] = strconv.Itoa(i)
	}
	fmt.Println(t, len(t))

	for key, value := range t {
		fmt.Println(key, value)
	}

	var y map[int]string
	y = map[int]string{1: "fgfggf", 2: "ggggggg"}
	fmt.Println(y, len(y))

	f := func(num int) int {
		return num * num
	}
	fmt.Println(f(10))

	var v interface{}
	v = 44
	b, ok := v.(string)
	if ok {
		fmt.Println(b)
	}*/

	// slices of any length can be assigned to other slice types
	/*	slice1 = slice2
		fmt.Println(slice1)*/
	/*
		slice2 = slice1
		fmt.Println(slice2)*/

	/*	st1 := "Hi"
		st2 := " dude"
		//sum := st1 + st2
		//fmt.Println(sum)
		var mass [3]int
		fmt.Println(mass)*/
	/*	var s string
		fmt.Println(s)
		s = "hi"*/
	/*	//gh := new(string)
		fr := make([]int, 0, 10)
		var g string = "rt"
		fmt.Println(g)
		fmt.Println(fr)*/
	//for {
	/*	switch s {
		case "hello":
			fmt.Println("hello")
		case "hi":
			fmt.Println("hi")
			break
			fmt.Println("poo-poo-poo")
			fallthrough
		default:
			fmt.Println("default")
		}

		var x interface{}
		x = 54
		x = "dddd"
		switch i := x.(type) {
		case string:
			fmt.Println("string")

		case int:
			fmt.Println("int")
			i++
		default:
			fmt.Println("other")
		}

		var mass [3][]int
		fmt.Println(mass)*/
	/*	var mass2 = make([]int, 2)
		g := f(mass2)
		fmt.Println(mass2, len(mass2), cap(mass2))
		fmt.Println(mass2[:2])
		fmt.Println(g, len(g), cap(g))
		fmt.Println(mass2[:], len(mass2), cap(mass2))
		//var b []int
		//b := []int{}
		b := make([]int, 10)
		if b == nil {
			fmt.Println("b is nil")
		} else {
			fmt.Println(b, len(b), cap(b))
		}
		b = b[3:10]
		fmt.Println(b, len(b), cap(b))
		t := append(b, 10)
		fmt.Println(t, len(t), cap(t))
		var q []int
		fmt.Println(append(q, 0))*/

	/*	var w map[int]string = make(map[int]string)
		w[20] = "hello"
		w[10] = "world"
		fmt.Println(len(w))
		for k, v := range w {
			fmt.Println(k, v)
		}

		var z map[int]string
		//z[2] = "eeee"
		el, ok := z[2]
		fmt.Println(el, ok)
		fmt.Println(z, len(z))

		//g := f([]int{})
		x := func(num int) int {
			return num * num
		}
		fmt.Println(x(3))

		func() int {
			fmt.Println("func anon")
			return 1
		}()

		b := 1
		func(num *int) {
			*num++
		}(&b)
		fmt.Println(b)*/

	/*	s := " hello п"
		fmt.Println(len(s), []byte(s), len([]byte(s)), []rune(s), len([]rune(s)))*/

	/*	var u []int
		//u = u[:3]
		t := append(u, 0)
		fmt.Println(u, t)*/
	/*	for i := 0; ; {
		fmt.Println(i)
		i++
		if i > 5 {
			break
		}
	}*/
	//}

	//var r int
	/*	var r = []int{5, 4, 3, 2, 1}
		var q = make([]int, 10)
		copy(q, r)
		fmt.Println(q)*/

	/*	bs := make([]byte, 1)
		bl := 0

		bl += copy(bs[bl:], []byte(st1))
		bl += copy(bs[bl:], []byte(st2))*/

	//fmt.Println(string(bs[:]))

	/*	buffer := bytes.Buffer{}
		buffer.WriteString(st1)
		buffer.WriteString(st2)
		fmt.Println(buffer.String())*/
	/*	s := "gopher"
		fmt.Println("Hello and welcome, %s!", s)

		for i := 1; i <= 5; i++ {
			//TIP You can try debugging your code. We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
			// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>. To start your debugging session,
			// right-click your code in the editor and select the <b>Debug</b> option.
			fmt.Println("i =", 100/i)
		}
	*/

	/*	t := z{10}
		t.y = 5
		r := t.io()
		fmt.Println(r, t.y)

		var tp interface{}
		tp = z{10}

		var baser base
		baser = &z{10}
		baser = base(&t)
		fmt.Println(reflect.TypeOf(baser))

		yi := baser.(*z)
		fmt.Println(reflect.TypeOf(yi))

		k := baser.(*h)
		fmt.Println(reflect.TypeOf(k))

		inter, ok := tp.(z)
		if ok {
			fmt.Println(inter)
		}
		fmt.Println(reflect.TypeOf(inter))*/
	//back, ok := inter.(z)

	/*	if ok {
			fmt.Println(back.io())
		}

		q := base(t)

		b := q.(z)*/

	defer printStack()
	f(3)

}

/*func f(mass []int) []int {
	for i := 0; i < 10; i++ {
		if len(mass) > i {
			mass[i] = i
			continue
		}
		mass = append(mass, i)
	}
	return mass
}*/

type base interface {
	io() int
}

type z struct {
	y int
}

func (i *z) io() int {
	i.y += 10
	return i.y
}

type h struct {
	y int
}

func (i *h) io() int {
	i.y += 50
	return i.y
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // Сбой при x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
