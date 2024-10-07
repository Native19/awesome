package go3

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrInvalidWidth = errors.New("параметр ширины невалидный")

func Go3() {}

type myStruct struct {
	num int
	str string
}

func Spr() {
	/*	str := fmt.Sprintf("изображение принудительно ужато: %v", ErrInvalidWidth)
		fmt.Println(str)*/

	number := 100500
	str := "Hello, world"

	myStruct := myStruct{number, str}

	//https://gobyexample.com.ru/string-formatting
	fmt.Printf("Вывод строки: %v, вывод числа %v\n", str, number) //
	fmt.Printf("Структура v: %v\n", myStruct)                     // универсальный вывод
	fmt.Printf("Структура +v: %+v\n", myStruct)                   // вывод с названием полей
	fmt.Printf("Структура #v: %#v\n", myStruct)                   // синтаксическое представление
	fmt.Printf("Структура T: %T\n", myStruct)                     // тип структуры
	fmt.Printf("Структура t: %t\n", true)                         // булевы
	fmt.Printf("Структура v: %v\n", true)                         //
	fmt.Printf("Структура d: %d\n", 56565656)                     // десятеричный вывод
	fmt.Printf("Структура v: %v\n", 111111)                       //
	fmt.Printf("Структура b: %b\n", 111111)                       // бинарный вывод
	fmt.Printf("Структура c: %c\n", 1118)                         // символ по числу
	fmt.Printf("Структура x: %x\n", 2123)                         // шестнадцатеричный
	fmt.Printf("Структура f: %f\n", 111.52)                       // float
	fmt.Printf("Структура e: %e\n", 123400000.0)                  // e вид
	fmt.Printf("Структура E: %E\n", 123400000.0)                  // e вид
	fmt.Printf("Структура s: %s\n", "\"string\"")                 // строка
	fmt.Printf("Структура q: %q\n", "\"str\"ing\"")               // строка с кавычками
	fmt.Printf("Структура x: %x\n", "hex this")                   // hex
	fmt.Printf("Структура x: %x\n", "32tсвl")                     // hex
	fmt.Printf("Структура p: %p\n", &myStruct)                    // указатель
	fmt.Printf("|%6d|%2d|\n", 12, 345)                            // указывает ширину числа
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)                      // ширина числа float
	fmt.Printf("|%.3f|%.2f|\n", 1.25757747328942, 3.45)           // ширина числа float
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)                    // выравнивание по левому краю
	fmt.Printf("|%6s|%6s|\n", "foo", "b")                         // ширина строк
	s := fmt.Sprintf("a %s", "string")                            // форматированная строка, без печати
	fmt.Println(s)
	_, _ = fmt.Fprintf(os.Stderr, "an %s\n", "error") // форматирование + поток
	_, _ = fmt.Fprint(os.Stderr, "an", "error\n")
	logger := log.New(os.Stderr, "", 0)
	logger.Println("hello !!!!!!!!!!!!!!!!!!!!!!!!!!!")
	log.Println("Hi")
	//log.Fatal("fatal")
	err := fmt.Errorf("Структура w: %v\n", "error") //
	err2 := fmt.Errorf("Вывод err: %w\n", err)      //
	errors.Unwrap(err2)
	fmt.Print("-------------------------------------\n")
	fmt.Print(errors.Unwrap(err2))
	fmt.Printf("Ошибка: %v\n", err2)
	fmt.Printf("nil: %v\n", nil)

}
