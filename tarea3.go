package main
import (
	"fmt"


)
type KEsimo struct {
    data []int // datos desordenados
    value int // valor del nodo (sólo si data == nil)
    lessThan int // cuantos valores hay a la izquierda
    left *KEsimo // hijo izquierdo
    right *KEsimo // hijo derecho
}
func New(data []int) *KEsimo {

    return &KEsimo{data: data}
}
func (ke *KEsimo) Get(k int) int{
  if ke.data==nil && ke.lessThan==k {
		fmt.Printf("hi ")
		fmt.Printf("%d ", ke.value)
		return ke.value
  }else if ke.data==nil && ke.lessThan<k {
    return ke.right.Get(k-ke.lessThan-1)
  }else if ke.data==nil && ke.lessThan>k {
    return ke.left.Get(k)
  }else{
    pivote:=ke.data[((len(ke.data)-1)/2)]
    ke.value=pivote
    i:=0
		fmt.Printf("%d ", pivote)
    left:=[]int{}
    right:=[]int{}
    for ((len(ke.data)-1)/2)>i{
      if ke.data[i]>ke.value{
        right=append(right,ke.data[i])
      }else{
        left=append(left,ke.data[i])
      }
      i+=1
    }
		i+=1
		for (len(ke.data)-1)>=i{
			if ke.data[i]>ke.value{
				right=append(right,ke.data[i])
			}else{
				left=append(left,ke.data[i])
			}
			i+=1
		}
    ke.left = &KEsimo{data: left}
    ke.right = &KEsimo{data: right}
    ke.lessThan = len(left)
		ke.data=nil
    return ke.Get(k)
  }
}
func main(){
  
}
