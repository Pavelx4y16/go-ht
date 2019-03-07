package main

import (
	"encoding/json"
	E "errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type NotFoundedError struct {
	message string
}

func (err *NotFoundedError) Error() string {
	return err.message
}

type Arguments map[string]string

type Person struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func (person *Person) Set(value string) error {
	return json.Unmarshal([]byte(value), person)
}

func (person *Person) String() string {
	result, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(result)
}

var operationFlag *string
var fileNameFlag *string
var idFlag *string
var itemFlag Person
var Exists bool

//var itemFlag string

func init() {
	operationFlag = flag.String("operation", "findById", "Getting type of Operation.")
	fileNameFlag = flag.String("fileName", "", "For file path.")
	idFlag = flag.String("id", "3", "for Id")
	flag.Var(&itemFlag, "item", "Getting item from console.")
	Exists = true
}

func GetListOfPersons(fileName string) ([]Person, error) {
	var result []Person
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(bytes, &result)
	file.Close()
	return result, err
}

func LoadPerson(fileName string) (Person, error) {
	var result Person
	file, err := os.Open(fileName)
	if err != nil {
		return result, err
	}
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&result)
	file.Close()
	return result, err
}

func GetListOfPersonsCommand(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte(""), err
	}
	file.Close()
	return bytes, err
}

func AddItem(item Person, fileName string) error {
	persons, err := GetListOfPersons(fileName)
	if err != nil {
		return err
	}
	if CheckItemExists(item.Id, persons) {
		return E.New("Item with id " + item.Id + " already exists")
	}
	persons = append(persons, item)

	return WriteToFile(persons, fileName)
}

func AddItemCommand(item string, fileName string) error {
	var personItem Person
	err := json.Unmarshal([]byte(item), &personItem)
	if err != nil {
		return err
	}
	return AddItem(personItem, fileName)
}

func findById(id string, fileName string) (Person, error) {
	Exists = true
	var result Person
	list, err := GetListOfPersons(fileName)
	if err != nil {
		return result, err
	}
	for _, value := range list {
		if value.Id == id {
			return value, err
		}
	}
	//err = E.New("Item with such id (" + id + ") was not founded!!!")
	Exists = false
	return result, err
}

func FindByIdCommand(id string, fileName string) (Person, error) {
	item, err := findById(id, fileName)
	if err != nil {
		return item, err
	}
	return item, err
}

func RemoveItem(id string, fileName string) error {
	persons, err := GetListOfPersons(fileName)
	if err != nil {
		return err
	}
	find := -1
	for index, value := range persons {
		if value.Id == id {
			find = index
		}
	}
	if find == -1 {
		return E.New("Item with such id (" + id + ") was not founded!!!")
	}
	persons = append(persons[:find], persons[find+1:]...)

	return WriteToFile(persons, fileName)
}

func RemoveItemCommand(id string, fileName string) error {
	err := RemoveItem(id, fileName)
	return err
}

func parseArgs() (args Arguments) {
	args = make(Arguments)
	args["operation"] = *operationFlag
	args["item"] = itemFlag.String()
	args["fileName"] = *fileNameFlag
	args["id"] = *idFlag
	return
}

func Perform(args Arguments, writer io.Writer) error {
	var err error
	if args["fileName"] == "" {
		return E.New("-fileName flag has to be specified")
	}
	switch args["operation"] {
	case "add":
		if args["item"] == "" {
			return E.New("-item flag has to be specified")
		}
		err = AddItemCommand(args["item"], args["fileName"])
	case "list":
		var listBytes []byte
		listBytes, err = GetListOfPersonsCommand(args["fileName"])
		writer.Write(listBytes)
	case "findById":
		if args["id"] == "" {
			return E.New("-id flag has to be specified")
		}
		var found Person
		found, err = FindByIdCommand(args["id"], args["fileName"])
		if Exists {
			writer.Write([]byte(found.String()))
		}
	case "remove":
		if args["id"] == "" {
			return E.New("-id flag has to be specified")
		}
		err = RemoveItemCommand(args["id"], args["fileName"])
	case "abcd":
		err = E.New("Operation abcd not allowed!")
	default:
		err = E.New("-operation flag has to be specified")
	}
	return err
}

func main() {
	flag.Parse()
	InfoFlags()
	//myFunction()
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	Success(*operationFlag)
}

func myFunction() {
	persons, err := GetListOfPersons("Db.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print("Before: ")
	fmt.Println(persons)
	index := 0
	persons = append(persons[:index], persons[index+1:]...)
	fmt.Print("After: ")
	fmt.Println(persons)
}

func InfoFlags() {
	fmt.Println(*operationFlag)
	fmt.Println(*idFlag)
	fmt.Println(itemFlag)
	fmt.Println(*fileNameFlag)
}

func Success(operation string) {
	fmt.Println("Operation \"" + operation + "\" was succesfylly complited!!!")
}

func WriteToFile(persons []Person, fileName string) error {
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}
	bytePersons, err := json.Marshal(persons)
	if err != nil {
		return err
	}

	_, err = file.Write(bytePersons)
	file.Close()
	return err
}

func Print(list []Person) (result string) {
	result = "["
	if len(list) > 0 {
		result += list[0].String()
	}
	for i := 1; i < len(list); i++ {
		result += "," + list[i].String()
	}
	result += "]"
	return
}

func CheckItemExists(id string, persons []Person) bool {
	for _, value := range persons {
		if value.Id == id {
			return true
		}
	}
	return false
}
