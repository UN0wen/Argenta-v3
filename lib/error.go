package lib

// CheckErr is a super simple error handler.
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
