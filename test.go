package main

import (
	"fmt"
	"math"
)

const imtpower = 2

func main() {
	fmt.Println("___КАЛЬКУЛЯТОР ИНДЕКСА МАССЫ ТЕЛА___")
	var kg, height = getUserInput()
	imt := calculateimt(kg, height)
	fmt.Println("Ваш ИМТ:", imt)

	if imt < 18.5 && imt > 10 {
		fmt.Println("Вам надо набрать массу")
	} else if imt < 25 && imt > 18.5 {
		fmt.Println("У вас нормальный вес")
	} else if imt > 25 && imt < 100 {
		fmt.Println("Вам надо сбросить вес")
	} else if imt < 10 || imt > 100 {
		fmt.Println("Вы ошиблись в записях")

	}

	outputResult(imt)

} //Вызываем функцию, возвращаем ее результат
func outputResult(imt float64) {
	fmt.Println("ваш имт:", imt)

} // делаем функцию для помещения в нее переменной отвечающей за расчет имт
func calculateimt(kg float64, height float64) float64 {
	var imt = float64(kg) / math.Pow(height/100, imtpower)
	return imt
} // создаем еще одну функции вводной части (ввод данных, перенос данных в часть с расчетами)
func getUserInput() (float64, float64) {
	height := 1.8 // рост
	kg := 63.0    // вес
	fmt.Printf("Введите свой рост (в сантиметрах)")
	fmt.Scan(&height) // записывает в перемнную данные которые он введет в этот блоке
	fmt.Printf("Введите свой вес (в киллограмах)")
	fmt.Scan(&kg) //записывает в перемнную данные которые он введет в этот блоке
	return kg, height

}
