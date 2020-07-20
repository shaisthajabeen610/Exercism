package isogram
import "strings"
//IsIsogram is a method to check whether the given string is an isogram or not
func IsIsogram(s string) bool{
	//if the string is empty then we don"t need the rest of the code
	 if s==""{
		return true
	}
	//in case the string is not empty we do the following comparisions
	s=strings.ToLower(s)
	re:=strings.NewReplacer("-",""," ","")
	//instead of using Stirngs.RplaceAll multiple times it better to use strings.NewReplacer
	s=re.Replace(s)
	for i, v:=range s{
		for j, t:=range s{
			if i<j{
				if v==t {
					return false
				}
				
			}
		}
	}
 return true
}
