package greeter

// Greet returns a greeting message
func Greet(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return "Hello, " + name + "!"
}
