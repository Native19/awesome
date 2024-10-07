package go3

import (
	"fmt"
	"time"
)

func Timme() {
	////fmt.Printf("time Now: %v\n", time.Now())
	////t := time.Now().Location().String()
	////loc := time.Now().Location()
	//ti := time.Now()
	////time.Minute
	//ti.Add(time.Minute * 10)
	//fmt.Printf("Сейчас + 10m: %v\n", ti)
	////fmt.Printf("time zone: %v\n", t)
	////fmt.Println(time.Now().Unix())
	//t := time.Now()
	//fmt.Println(t)
	////fmt.Printf(t.Format(time.RFC3339))
	////fmt.Printf(t.Format(time.RFC3339))
	//fmt.Println(t.Format(time.TimeOnly))
	//fmt.Println(t.Format(time.Kitchen))
	//fmt.Println(t.Format(time.DateOnly))
	//fmt.Println(t.Format(time.RFC3339))
	//
	//fmt.Println(t.After(ti))
	//fmt.Println(t.Before(ti))
	//fmt.Println(t.Format(time.DateTime))
	//fmt.Println(t.Zone())
	//fmt.Println(t.ZoneBounds())
	//
	//phx, err := time.LoadLocation("America/Phoenix")
	//if err != nil {
	//	panic(err)
	//}
	//
	//t2 := t.In(phx)
	//fmt.Printf("Конвертированное время: %v\n", t2)

	/*	dt1 := time.Date(2024, time.April, 29, 12, 52, 4, 0, time.Local)
		dt2, err := time.Parse("2006-01-02 15:04:05", "2024-04-29 12:52:04")
		if err != nil {
			panic(err)
		}

		fmt.Println("Date and time:", dt1)
		fmt.Println("Date and time:", dt2)

		dt := time.Now()

		fmt.Println("RFC822:", dt.Format(time.RFC822))
		fmt.Println("Unix:  ", dt.Format(time.UnixDate))
		fmt.Println("Custom:", dt.Format("2006/01/02 3:04 PM"))*/

	/*dateTimeStr1 := "2024-05-02T15:28:00+07:00"
	dateTimeStr2 := "2024-05-02 15:28:00"

	dt1, err := time.Parse(time.RFC3339, dateTimeStr1)
	if err != nil {
		panic(err)
	}

	locationNsk, err := time.LoadLocation("Asia/Novosibirsk")
	if err != nil {
		panic(err)
	}
	dt2, err := time.ParseInLocation(time.DateTime, dateTimeStr2, locationNsk)
	if err != nil {
		panic(err)
	}

	fmt.Println("date-time var.1: ", dt1)
	fmt.Println("date-time var.2: ", dt2)

	locationMsk, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		panic(err)
	}

	fmt.Println("date-time in MSK:", dt1.In(locationMsk))

	v := time.Duration(13777770000000000)
	fmt.Println(v)

	fmt.Println(time.Now().Format(time.RFC3339))
	fmt.Println(time.Now().After(dt1))

	apochNow := time.Now().Unix()
	fmt.Printf("Время эпохи в секундах: %d\n", apochNow)
	*/

	/*	t := time.Now().Unix()
		fmt.Println(t)*/

	//timeCh := make(chan time.Time)
	/*	t := time.Now()
		fmt.Println(t.Clock())*/
	t := <-time.After(2 * time.Second)
	fmt.Println(t)
	fmt.Println("ready")
	tn := time.Now()
	location := time.FixedZone("UTC", 0)
	newTime := tn.In(location)
	fmt.Println(newTime)
	//tn.After()
	//time.After()
	//time.Second

	fmt.Println(time.Now().UTC().Format(time.RFC3339))
}
