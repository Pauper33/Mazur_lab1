package main
import (
  "fmt"
  "time"
  "log"
  "encoding/json"
  "net/http"
)

func main() {
  //структура для создание документа с временем сайта
  type Time struct{
  //строка вывода в документе
    Time string `json:"Server time"`
  }
  //создание документа
  http.HandleFunc("/time", func(doc http.ResponseWriter, r *http.Request) {
    //задаем переменной Time_server строку, которая выводится в документе (строка st+время)
    Time_server := &Time{Time: time.Now().Format(time.RFC3339)} 
    //задаём переменной doc_json документ с форматом JSON с заданной строкой
    doc_json, _ := json.Marshal(Time_server)
    //создание документа
    doc.Header().Set("content-type", "aplication/json")
    doc.WriteHeader(200)
    doc.Write(doc_json)
  })
  //вывод в консоль, если все хорошо
  fmt.Println("Starting HTTP server")
  //вывод ошибки, если была
  log.Fatal(http.ListenAndServe(":8795", nil))
}
