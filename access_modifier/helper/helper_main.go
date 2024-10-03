package helper

var version = "1.0.0"      // <- cannot be accessed from other package
var Application = "Golang" // <- can be accessed from other package

// Cannot be accessed from other package
func printMessage(msg string) string {
	return "Your message is: " + msg
}

// Can be accessed from other package
func ShowMsg(msg string) string {
	return version + "-msg: " + msg
}
