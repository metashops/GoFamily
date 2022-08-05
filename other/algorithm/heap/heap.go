package main

func replaceSpace(s string) string {
	b := []byte(s)
	result := make([]byte,0)
	for i := 0; i < len(b);i++{
		if b[i] == ' ' {
			result = append(result,[]byte("%20")...)
		} else {
			result = append(result,b[i])
		}
	}
	return string(result)
}
func main() {
	s := "a b c  e"
	print(replaceSpace(s))

}