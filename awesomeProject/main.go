package main

import (
	"fmt"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {

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

	s := " hello п"
	fmt.Println(len(s), []byte(s), len([]byte(s)), []rune(s), len([]rune(s)))

	/*	for i := 0; ; {
		fmt.Println(i)
		i++
		if i > 5 {
			break
		}
	}*/
	//}

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
		}*/
}

func f(mass []int) []int {
	for i := 0; i < 10; i++ {
		if len(mass) > i {
			mass[i] = i
			continue
		}
		mass = append(mass, i)
	}
	return mass
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
