package interfacesI
import (
	"fmt"
)
type PrintVowels interface{
	FindVowels() []rune
}
type MyStr string

func (st MyStr) FindVowels() []rune {
	var vowels []rune
	for _, str := range st {
		if str =='a' ||str =='e' ||str =='i' ||str =='o' ||str =='u'  {
			vowels = append(vowels, str)
		}
	}
	return vowels
}

func CallerValue ()  {
	str := MyStr("harish") 
	// implementing the interfaces
	var v PrintVowels
	v = str // implementing an interface
	fmt.Printf("%c \n",v.FindVowels()) // both works 
	fmt.Printf("%c \n",str.FindVowels())
}