package pg

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type ConnectionData struct {
	user, password, host, port string
}

func NewConnData(user, password, host, port string) *ConnectionData {
	connData := ConnectionData{user, password, host, port}
	return &connData
}

func (con *ConnectionData) Pg() {
	conn, err := con.pgConnection()
	if err != nil {
		panic(err)
	}
	//defer func() { _ = conn.Close() }()

	/*	if err := insertProduct(conn, "IPhoneX", "Apple", 100000); err != nil {
		panic(err)
	}*/

	//DelAll(conn)
	/*	go func(i int) {
		if err := insertProduct(conn, "IPhone"+strconv.Itoa(i), "Apple", i*1000); err != nil {
			fmt.Println(err)
		}

	}(5)*/
	var id int
	var model string
	var price int
	result, err := SelectAll(conn)
	//result, err := conn.Query("select id, model, price  from Products")
	if err != nil {
		return
	}
	for result.Next() {
		err := result.Scan(&id, &model, &price)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, model, price)
	}

	/*	for results.Next() {

		}*/
	//DelAll(conn)
	/*	time.Sleep(1 * time.Second)
		for i := range 1000 {
			go func(i int) {
				if err := insertProduct(conn, "IPhone"+strconv.Itoa(i), "Apple", i*1000); err != nil {
					fmt.Println("ffdfdff", err)
				}
			}(i)
			//fmt.Println("Итерация ", i)
		}*/
	//time.Sleep(30 * time.Second)

	//DelAll(conn)
}

func insertProduct(conn *sql.DB, model string, company string, price int) error {
	var id int
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err := conn.QueryRowContext(ctx, "insert into Products (model, company, price) values ($1, $2, $3) returning id",
		model, company, price).Scan(&id)
	cancel()
	if err != nil {
		return err
	}

	fmt.Println("Добавлен id: ", id)
	return nil
}

func DelAll(conn *sql.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := conn.ExecContext(ctx, "delete from Products *")
	if err != nil {
		fmt.Println(err)
	}
	cancel()
	fmt.Println(res)
}

func SelectAll(conn *sql.DB) (*sql.Rows, error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	/*	ch := make(chan *sql.Rows, 1)
		go func(ch chan *sql.Rows) {*/
	//data, err := conn.QueryContext(ctx, "select id, model, price  from Products")
	data, err := conn.Query("select id, model, price  from Products")
	if err != nil {
		return nil, err
	}
	return data, nil
	/*		ch <- data
	}(ch)*/
	/*	if err != nil {
		fmt.Println(err)
	}*/
	//_ = cancel
	//return res
	/*	select {
		case row := <-ch:
			return row, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}*/
}

func (con *ConnectionData) pgConnection() (*sql.DB, error) {
	connStr :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
			con.user, con.password, con.host, con.port)

	db, err := sql.Open("postgres", connStr)
	db.SetMaxOpenConns(10)
	if err != nil {
		return nil, fmt.Errorf("подключение к БД: %w", err)
	}
	fmt.Println(db)
	return db, nil
}
