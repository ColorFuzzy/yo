package zo

// Panic just panic on error, otherwise do nothing
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}