package pg

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type weather struct {
	city string
	mint int
	maxt int
	pcrp float32
	date time.Time
}

func Pg2() {

	//pgx.ConnConfig{}
	/*	conn, err := pgx.Connect(context.Background(), "postgres://postgres:123456@localhost:5432/postgres?sslmode=disable")
		pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))*/
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()
	fmt.Println("Successfully connected to postgres")

	// insert
	/*	{
		fmt.Println(createWeather(dbpool))
		//timeNow := time.Now()
		newDate, _ := time.Parse(time.DateOnly, "1994-05-05")
		err = insertWeather(dbpool, "San Francisco", 32, 40, 0.5, newDate)
	}*/

	// select
	{
		selected, err := selectAllFromWether(dbpool)
		if err != nil {
			fmt.Println("Pg2 Err: ", err)
		}

		for i, selectedWeather := range selected {
			fmt.Println(i, selectedWeather, selectedWeather.date.Format(time.DateOnly))
		}
	}

	/*	{
		selected, err := selectAvgTempFromWeather(dbpool)
		if err != nil {
			fmt.Println("Pg2 Err: ", err)
		}

		for i, selectedWeather := range selected {
			fmt.Println(i, selectedWeather)
		}
	}*/

	// delete
	/*	{
		if err := deleteWeather(dbpool); err != nil {
			fmt.Println("Pg2 delete: ", err)
		}
	}*/

	//fmt.Println(dropWeather(dbpool))

}

func insertWeather(pool *pgxpool.Pool,
	city string, mint int, maxt int, pcrp float32, date time.Time,
) error {
	row, err := pool.Query(context.Background(),
		"insert into weather (city, temp_lo, temp_hi, prcp, date) values ($1, $2, $3, $4, $5)", city, mint, maxt, pcrp, date)
	//"insert into weather (city, temp_lo, temp_hi, prcp, date) values ($1, $2, $3, $4, $5) returning temp_lo, temp_hi as new_temp",
	//city, mint, maxt, pcrp, date)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query err: %v\n", err)
		return err
	}
	for row.Next() {
		var temp int
		var temp2 int
		row.Scan(&temp, &temp2)
		fmt.Println("Row: ", temp, temp2)
	}
	row.Close()
	fmt.Println("Row err: ", row.Err())
	return row.Err()
}

func createWeather(pool *pgxpool.Pool) error {
	row, err := pool.Query(context.Background(),
		"create table weather (city varchar, temp_lo int, temp_hi int, prcp real, date date);")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query err: %v\n", err)
		return err
	}
	row.Close()
	fmt.Println(row.Err())
	return row.Err()
}

func dropWeather(pool *pgxpool.Pool) error {
	row, err := pool.Query(context.Background(),
		"drop table weather;")
	if err != nil {
		fmt.Printf("Query err: %v\n", err)
		return err
	}
	row.Close()
	fmt.Println(row.Err())
	return row.Err()
}

func selectAllFromWether(pool *pgxpool.Pool) ([]weather, error) {
	rows, err := pool.Query(context.Background(),
		"select * from weather /*order by date*/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query err: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	var weathers = make([]weather, 0)
	for rows.Next() {
		newWeather := weather{}
		err := rows.Scan(
			&newWeather.city,
			&newWeather.mint,
			&newWeather.maxt,
			&newWeather.pcrp,
			&newWeather.date)
		if err != nil {
			fmt.Printf("Scan err: %v\n", err)
			break
		}
		weathers = append(weathers, newWeather)
		fmt.Println("Row: ", len(weathers), newWeather)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Row err:", err)
		return weathers, err
	}
	return weathers, nil
}

func selectAvgTempFromWeather(pool *pgxpool.Pool) ([]int, error) {
	rows, err := pool.Query(context.Background(),
		"select (temp_hi+temp_lo)/2 as temp_avg from weather")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query err: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	temps := []int{}
	for rows.Next() {
		var temp int
		err := rows.Scan(&temp)
		if err != nil {
			fmt.Printf("Scan err: %v\n", err)
		}
		temps = append(temps, temp)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Row err:", err)
		return temps, err
	}
	return temps, nil
}

// select count(*)
// select *  using  QueryRow   will be error?

func deleteWeather(pool *pgxpool.Pool) error {
	commandTag, err := pool.Exec(context.Background(),
		"delete from weather where city is null")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query err: %v\n", err)
		return err
	}
	fmt.Println("commandTag.String: ", commandTag.String())
	fmt.Println("RowsAffected: ", commandTag.RowsAffected())
	fmt.Println("Delete: ", commandTag.Delete())
	return nil
}
