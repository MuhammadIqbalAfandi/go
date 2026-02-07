package helper

// Access Modifier
var version = "1.0.0" // unexported variable
var AppName = "GoApp" // exported variable


// SayHello is an exported function
func SayHello(name string) string {
	return "Hello " + name
}