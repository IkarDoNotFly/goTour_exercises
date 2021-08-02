package main
//Exercise: Slices
import (
	"golang.org/x/tour/pic"
)
func Pic(dx, dy int) [][]uint8 {

	sliceDx:=make([]uint8,dx)
	sliceDy:=make([][]uint8,dy)
	for i:=range sliceDy{
		for j:=range sliceDx{
			sliceDx[j]=uint8(i+j)
		}
		sliceDy[i]=sliceDx
	}
	return sliceDy
}
func main(){
	pic.Show(Pic)
}
