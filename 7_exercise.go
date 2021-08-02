package main
import "golang.org/x/tour/reader"
type MyReader struct{}
func (r MyReader) Read(b []byte) (int, error) {
	a := 'A'
	for i:=0; i < len(b); i++ {
		b[i] = byte(a)
	}
	return len(b), nil
}
// TODO: Add a Read([]byte) (int, error) method to MyReader.
func main() {
	reader.Validate(MyReader{})
}