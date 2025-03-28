package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

var idc int = 1

type Student struct {
	Id   int    "json:id"
	Name string "json:name"
	Age  int    "json:age"
}

func AddToFile(student Student) bool {
	file, err := os.OpenFile("StudentData.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error while opening the file", err)

	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(student)
	if err != nil {
		fmt.Println("Error while encoding the file", err)
		return false
	}
	return true
}

func RetriveFromFile() []Student {
	file, err := os.OpenFile("StudentData.json", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("Error while opening file for reading", err)
	}
	defer func() {
		file.Close()
	}()
	students := []Student{}
	decoder := json.NewDecoder(file)
	for {
		var student Student
		err := decoder.Decode(&student)
		if err == io.EOF {
			break // End of file
		}
		if err != nil {
			log.Printf("Error decoding student: %v", err)
			return nil
		}
		students = append(students, student)
	}
	return students
}
func AddStudent(name string, age int) {

	Student := Student{
		Id:   idc,
		Name: name,
		Age:  age,
	}
	idc++
	check := AddToFile(Student)
	if check == false {
		fmt.Println("Error add the student")
		return
	} else {
		fmt.Println("Student Added successfully")
	}
}
func DeleteStudent(id int) (string, error) {

	return "", nil
}

func SearchStudent(id int) Student {

	StudentsArray := RetriveFromFile()
	if len(StudentsArray) == 0 {
		fmt.Println("No students found.")
	} else {

		for _, student := range StudentsArray {
			if student.Id == id {
				return student
			}

		}
	}
	return Student{}
}

func main() {
	choice := ""
	for choice != "exit" {
		fmt.Println("please enter your choice \n")
		fmt.Println("Press 1 to add student ")
		fmt.Println("Press 2 to delete student ")
		fmt.Println("Press 3 to search student ")
		fmt.Println("Press exit to exit the program ")

		fmt.Scanln(&choice)

		switch choice {
		case "exit":
			fmt.Println("Goodbye")
			return

		case "1":
			var name string
			var age int

			fmt.Println("please enter the name of the student:")
			fmt.Scanln(&name)

			fmt.Println("please enter the age of the student:")
			fmt.Scanln(&age)
			AddStudent(name, age)

		case "2":
			fmt.Println("2")

		case "3":
			fmt.Println("Please return the id of the student you want to find")
			var id int
			fmt.Scanln(&id)
			student := SearchStudent(id)
			if student.Id == 0 {
				fmt.Println("No student found")
			} else {
				fmt.Println("Id:", student.Id)
				fmt.Println("Name:", student.Name)
				fmt.Println("Age:", student.Age)
			}
		default:
			fmt.Println("invalid choice")
		}
	}

}
