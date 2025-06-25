package main

import (
	"fmt"
	"math"
)

const imtpower = 2

func main() {
	for {

		fmt.Println("___КАЛЬКУЛЯТОР ИНДЕКСА МАССЫ ТЕЛА___")
		var kg, height = getUserInput()
		imt := calculateimt(kg, height)

		if imt < 18.5 && imt > 10 {
			for {

				fmt.Println("Ваш IMT =", imt, "Вам надо набрать массу")
				fmt.Println("Подсказать упражнения для набора массы? (yes or not)")
				yesornot := ""
				fmt.Scan(&yesornot)
				if yesornot == "да" || yesornot == "if" || yesornot == "yes" || yesornot == "yeah" {
					/////////////////////////////////////////////////////////////////////////

				} else if yesornot == "not" || yesornot == "no" || yesornot == "нет" || yesornot == "не" {
					break

				} else {
					fmt.Println("Вы ввели что-то не то ")

				}
			}
		} else if imt < 25 && imt > 18.5 {
			for {
				fmt.Println("Ваш IMT =", imt, "У вас нормальный вес")
				fmt.Println("Подсказать упражнения для сохранения вашей формы?")
				yesornot := ""
				fmt.Scan(&yesornot)
				if yesornot == "да" || yesornot == "if" || yesornot == "yes" || yesornot == "yeah" {
					/////////////////////////////////////////////////////////\

				} else if yesornot == "not" || yesornot == "no" || yesornot == "нет" || yesornot == "не" {
					break

				} else {
					fmt.Println("Вы ввели что-то не то ")

				}
			}
		} else if imt > 25 && imt < 100 {
			for {
				fmt.Println("Ваш IMT =", imt, "Вам надо сбросить вес")
				fmt.Println("Подсказать упражнения для сброса лишнего веса?")
				yesornot := ""
				fmt.Scan(&yesornot)
				if yesornot == "да" || yesornot == "if" || yesornot == "yes" || yesornot == "yeah" {

				} else if yesornot == "not" || yesornot == "no" || yesornot == "нет" || yesornot == "не" {
					break

				} else {
					fmt.Println("Вы ввели что-то не то ")

				}
			}
		} else if imt < 10 || imt > 100 {
			fmt.Println("Ваш IMT =", imt, "Вы ошиблись в записях")

		}
	}
}

// делаем функцию для помещения в нее переменной отвечающей за расчет имт
func calculateimt(kg float64, height float64) float64 {
	var imt = float64(kg) / math.Pow(height/100, imtpower)
	return imt
} // создаем еще одну функции вводной части (ввод данных, перенос данных в часть с расчетами)
func getUserInput() (float64, float64) {
	height := 1.8 // рост
	kg := 63.0    // вес
	fmt.Printf("Введите свой рост (в метрах)")
	fmt.Scan(&height) // записывает в перемнную данные которые он введет в этот блоке
	fmt.Printf("Введите свой вес (в сантиметрах)")
	fmt.Scan(&kg) //записывает в перемнную данные которые он введет в этот блоке
	return kg, height

}
