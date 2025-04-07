package main

import (
	"fmt"
	"net/http"
	"reflect"
)

type Login struct {
	Email    string `json:"email" required:"true" min:"10" max:"100"`
	Password string `json:"password" required:"true" min:"3" max:"10"`
}

type User struct {
	Name     string `json:"name" required:"true" min:"3" max:"10"`
	Age      int    `json:"age" required:"true" min:"1" max:"100"`
	Email    string `json:"email" required:"true" min:"10" max:"100"`
	Password string `json:"password" required:"true" min:"3" max:"10"`
}

func ValidateStruct(s interface{}) error {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct, got %T", s)
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i)

		if field.Tag.Get("required") == "true" && value.String() == "" {
			return fmt.Errorf("%s is required", field.Name)
		}

		min := field.Tag.Get("min")
		max := field.Tag.Get("max")

		if min != "" && value.Len() < len(min) {
			return fmt.Errorf("%s must be at least %s characters long", field.Name, min)
		}
		if max != "" && value.Len() > len(max) {
			return fmt.Errorf("%s must be at most %s characters long", field.Name, max)
		}
	}

	return nil
}

func calculateHandler() {
	fmt.Println("Calculating endpoint...")
}

func calculateHandler2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Calculating endpoint 2...")
	}
}

func main() {
	newLogin := Login{
		Email:    "asd@mail.com",
		Password: "123456",
	}

	err := ValidateStruct(newLogin)
	if err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Validation successful")
	}

	newUser := User{
		Name:     "John",
		Age:      25,
		Email:    "John@mail.com",
		Password: "123456",
	}
	err = ValidateStruct(newUser)
	if err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Validation successful")
	}

	// userType := reflect.TypeOf(newUser)
	// userValue := reflect.ValueOf(newUser)

	// if newUser.Email < userType.Field(1).Tag.Get("min") {
	// 	fmt.Println("Email is too short")
	// }
	// if newUser.Email > userType.Field(1).Tag.Get("max") {
	// 	fmt.Println("Email is too long")
	// }
	// if newUser.Password < userType.Field(1).Tag.Get("min") {
	// 	fmt.Println("Password is too short")
	// }
	// if newUser.Password > userType.Field(1).Tag.Get("max") {
	// 	fmt.Println("Password is too long")
	// }

	// fmt.Println("Type:", userType)
	// fmt.Println("Value:", userValue)
	// fmt.Println("Email:", userValue.FieldByName("Email"))
	// fmt.Println("Password:", userValue.FieldByName("Password"))
	// fmt.Println("Password > Tag > Json:", userType.Field(1).Tag.Get("json"))
	// fmt.Println("Password > Tag > Required:", userType.Field(1).Tag.Get("required"))
	// fmt.Println("Password > Tag > Min:", userType.Field(1).Tag.Get("min"))
	// fmt.Println("Password > Tag > Max:", userType.Field(1).Tag.Get("max"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Guys!")
	})

	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		calculateHandler()
	})

	http.HandleFunc("/calculate2", calculateHandler2())

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
