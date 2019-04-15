package main

import (
//	"io"
	"log"
	"net/http"
	"encoding/json"
//	"fmt"
//	"strconv"
)

//Структура для входных параметров 
//Имена переменныех должны начинаться с большой буквы
type Data struct{
  X int `json: "x"`
  Y int `json: "y"`
}

//Структура данных результата
type RData struct{
  R int `json: "r"`
}

//обработчик для функции Ajax запроса расчет потенциала поля по координатам 
func myfuncAjax(w http.ResponseWriter, r *http.Request){
    var d Data    
    var rp RData
    //fmt.Println("%v",r.Body)
    //Читаем тело запроса r.Body и преобразуем данные в структуру Data
    err := json.NewDecoder(r.Body).Decode(&d)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    // Здесь пиши свой расчет значения потенциала поля
    // пока возвращаем сумму значений
    // fmt.Println("x= %d y= %d",d.X,d.Y) //это я для отладки ставил вывод в консоль
    //fmt.Println("%+v",d)
    rp.R = d.X+d.Y


    // Создаем json response from struct RData
    a, err := json.Marshal(rp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    // Задаем заголовок данным для возврата ответа
    w.Header().Set("Content-Type", "application/json")
    //Возвращаем результат в виде json структуры
    w.Write(a)
}


func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))
        // Определяем маршрут и обработчик для функции Ajax запроса
	//myfuncAjax - это функция описанная выше, 
        http.HandleFunc("/myfunc", myfuncAjax)

	log.Fatal(http.ListenAndServe(":8081", nil))
}