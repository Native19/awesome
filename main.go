package main

import (
	"flag"
	"time"

	"awesome/awesomeProject/go1"
	"awesome/awesomeProject/go2"
	"awesome/awesomeProject/go3"
	"awesome/awesomeProject/pg"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func init() {
	go1.Go1()
	go2.Go2()
	go3.Go3()
}

func main() {

	pg.Pg2()

}

func go3Main() {
	/*conData := pg.NewConnData(
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DB_NAME"),
		os.Getenv("PG_SSL_MODE"),
	)
	conData.Pg()*/

	//go2.Wg()
	//go2.Pract()
	//go3.Spr()
	//go3.Er()
	//go3.Json()
	//go3.Timme()
	//go3.ReaderIo()
	//go3.Env()
	//go3.Ht()
	//go3.Cont()
	//go3.Chmod()

	/*	var ch = make(chan int)
			close(ch)
		Loop:
			for {
				select {
				case _, ok := <-ch:
					if !ok {
						ch = nil
						//break Loop
					}
					fmt.Println("1")
				}
			}*/

	//go2.Pract()

	//go2.Timer()
	//defer fmt.Println("hello world")
	//go2.Uint()
	/*	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()*/
	/*	fmt.Println("Hello main")
		defer fmt.Println("defer main")
		os.Exit(0)*/
	//go2.Chans()
	/*	go2.Wg()
		os.Exit(0)*/
	//go2.Sum()
	//time.Sleep(5 * time.Second)
	//go2.Spinner1()
	/*	go2.Sum()
		var r int64
		j := int8(r)
		_ = j*/
	//
	//go2.Clock1()

	/*	var ch2 chan int
		ch := make(<-chan int)
		close(ch)
		close(ch2)*/
	//close(ch)

	/*	fmt.Errorf("hi")
		errors.New("new")
		//go2.Main2()
		var inst *go1.Z
		var inter go1.Base
		inter = inst

		if inter == nil {
			fmt.Println("inter is nil")
		} else {
			fmt.Println("inter is not nil")
			fmt.Println(inter.Io())
		}*/

	/*	flag.Parse()
		fmt.Printf("Ожидание %v...", *period)
		time.Sleep(*period)
		fmt.Println()*/
}
