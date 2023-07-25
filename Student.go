package main

import (
	"encoding/csv"

	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	// "github.com/gocarina/gocsv"

	"github.com/labstack/echo"
)

var headercounter int = 0
var count int = 0

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Number int    `json:"number"`
	Score  int    `json:"score"`
	Id     int    `json:"id"`
}

type PersonReq struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Number int    `json:"number"`
}

func main() {
	e := echo.New()
	e.GET("/", GetAllStudent)
	e.GET("/:id", Getstudent)

	e.POST("/", CreateStudent)
	e.Logger.Fatal(e.Start(":3000"))

}

func GetAllStudent(c echo.Context) error {

	allstudents, _ := ReadFileCsv("people.csv")
	return c.JSON(http.StatusOK, allstudents)
}

func Getstudent(c echo.Context) error {

	// id := c.Param("id")
	// if ReadFileCsv(){
	// 	id := strconv.ParseInt(int64(newStudent.Id), 10)
	// 	id := newStudent.Id +
	// }
	// for id := 0, csv.Reader, id++{
	// 	return id

	// }

	students, _ := ReadFileCsv("people.csv")
	return c.JSON(http.StatusOK, students)
}

func CreateStudent(c echo.Context) error {
	var req PersonReq
	err := c.Bind(&req)
	if err != nil {
		print(err)
		return err
	}
	newStudent := Person{
		Name:   req.Name,
		Age:    req.Age,
		Number: req.Number,
	}
	//Add Id

	data := []string{strconv.FormatInt(int64(newStudent.Id), 10), newStudent.Name, strconv.FormatInt(int64(newStudent.Age), 10), strconv.FormatInt(int64(newStudent.Number), 10)}

	newStudent.Id = count
	count++
	data = append(data, strconv.FormatInt(int64(newStudent.Id), 10))

	err = WriteInCsv(newStudent)

	if err != nil {
		return err
	}
	fmt.Println(newStudent)
	return c.JSON(http.StatusOK, newStudent)

}

func WriteInCsv(newStudent Person) error {

	file, err := os.OpenFile("people.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()
	if headercounter == 0 {
		header := []string{"Id", "Name", "Age", "Number"}
		w.Write(header)
		headercounter++ 
	}

	// fmt.Printf("Id", "Name", "Age", "Number")

	// if err != nil {
	// 	panic(err)
	// }
	defer file.Close()

	newdata, _ := ReadFileCsv("people.csv")
	// CreateStudent(newStudent)
	data := []string{strconv.FormatInt(int64(newStudent.Id), 10), newStudent.Name, strconv.FormatInt(int64(newStudent.Age), 10), strconv.FormatInt(int64(newStudent.Number), 10)}
	fmt.Println("Print Data", newdata)
	//Id
	// newStudent.Id++
	// data = append(data, strconv.FormatInt(int64(newStudent.Id), 10))
	// fmt.Println(newStudent.Id)

	// totalQuestions := len(data)
	// fmt.Println("Total no: of rows:", totalQuestions)
	// for e, value := range data {
	// 	fmt.Println(e, value)
	// }

	// }
	// 	count := 0
	// 	dataLength := len(data)
	// 	for dataLength > 0 {
	// 		newStudent.Id = count
	// 		newStudent.Id++
	// 		fmt.Println(newStudent.Id)
	// 		if dataLength == len(data) {
	// 			break
	// 		}
	// 	}
	data = append(data)
	err = w.Write(data)
	if err != nil {
		fmt.Print("error in right in csv file", err)
	}

	w.Flush()
	return nil
}

func ReadFileCsv(fileName string) ([]Person, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []Person{}, err
	}

	defer file.Close()

	r := csv.NewReader(file)

	// skip first line
	if _, err := r.Read(); err != nil {
		return []Person{}, err
	}
	studens := []Person{}
	records, err := r.ReadAll()

	for _, rec := range records {

		age, _ := strconv.Atoi(rec[2])
		id, _ := strconv.Atoi(rec[0])
		number, _ := strconv.Atoi(rec[3])
		name := rec[1]

		student := Person{
			Name:   name,
			Age:    age,
			Id:     id,
			Number: number,
		}
		studens = append(studens, student)
	}

	if err != nil {
		return []Person{}, err
	}
	fmt.Println(records)
	return studens, nil
}
