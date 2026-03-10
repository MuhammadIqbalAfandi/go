package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidator(t *testing.T) {
	var validate = validator.New()
	if validate == nil {
		t.Error("Failed to create validator")
	}
}

func TestValidatorVariable(t *testing.T) {
	validate := validator.New()

	user := "john"

	err := validate.Var(user, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidatorTwoVariables(t *testing.T) {
	validate := validator.New()

	password := "password123"
	confirmPassword := "password123"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidatorMultipleTagParameters(t *testing.T) {
	validate := validator.New()

	user := "john"

	err := validate.Var(user, "required,min=3,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidatorStruct(t *testing.T) {
	type User struct {
		Name     string `validate:"required,min=3,max=10"`
		Password string `validate:"required,min=6"`
	}

	user := User{
		Name:     "john",
		Password: "password123",
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidatorErrors(t *testing.T) {
	type User struct {
		Name     string `validate:"required,min=3,max=10"`
		Password string `validate:"required,min=6"`
	}

	user := User{
		Name:     "j",
		Password: "pass",
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			fmt.Printf("Field: %s, Tag: %s, Value: %v\n", validationError.Field(), validationError.Tag(), validationError.Value())
		}
	}
}

func TestValidatorCrossField(t *testing.T) {
	type User struct {
		Name            string `validate:"required,min=3,max=10"`
		Password        string `validate:"required,min=6"`
		ConfirmPassword string `validate:"required,eqfield=Password"`
	}

	user := User{
		Name:            "jhon",
		Password:        "pass123",
		ConfirmPassword: "pass123",
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidatorNestedStruct(t *testing.T) {
	type Address struct {
		City  string `validate:"required"`
		State string `validate:"required"`
	}

	type User struct {
		Name    string  `validate:"required,min=3,max=10"`
		Address Address `validate:"required"`
	}

	user := User{
		Name: "john",
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidatorCollections(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Name      string    `validate:"required,min=3,max=10"`
		Addresses []Address `validate:"required,dive"`
	}

	user := User{
		Name: "john",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidatorBasicCollections(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Name      string    `validate:"required,min=3,max=10"`
		Addresses []Address `validate:"required,dive"`
		Hobbies   []string  `validate:"required,dive,required,min=3"`
	}

	user := User{
		Name: "john",
		Addresses: []Address{
			{
				City:    "New York",
				Country: "USA",
			},
			{
				City:    "Los Angeles",
				Country: "USA",
			},
		},
		Hobbies: []string{
			"reading",
			"swimming",
		},
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidatorMaps(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Name        string            `validate:"required,min=3,max=10"`
		Addresses   []Address         `validate:"required,dive"`
		Hobbies     []string          `validate:"required,dive,required,min=3"`
		Preferences map[string]string `validate:"required,dive,keys,required,min=3,endkeys,required,min=3"`
	}

	user := User{
		Name: "john",
		Addresses: []Address{
			{
				City:    "New York",
				Country: "USA",
			},
			{
				City:    "Los Angeles",
				Country: "USA",
			},
		},
		Hobbies: []string{
			"reading",
			"swimming",
		},
		Preferences: map[string]string{
			"theme": "dark",
			"lang":  "eng",
		},
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidatorAlias(t *testing.T) {
	type User struct {
		Name string `validate:"username"`
	}

	validate := validator.New()
	validate.RegisterAlias("username", "required,min=3,max=10")

	user := User{
		Name: "john",
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}

	return true
}

func TestValidatorCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required,min=6"`
	}

	loginRequest := LoginRequest{
		Username: "JOHNNY",
		Password: "password123",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

func TestValidatorCustomValidationWithParameter(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	type User struct {
		Name string `validate:"required,min=3,max=10"`
		Pin  string `validate:"required,pin=4"`
	}

	user := User{
		Name: "john",
		Pin:  "1234",
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestOrRule(t *testing.T) {
	type loginRequest struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required,min=6"`
	}

	request := loginRequest{
		Username: "john@example.com",
		Password: "password123",
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}
