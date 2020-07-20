package scrabble
import  "strings"
//Score is a function used to find the score of given string
//score returns the count for occurence of characters
func Score(s string) (c int) {
	s= strings.ToUpper(s)
    c=0
	maptype :=map[string]int{
		 "AEIOULNRST":1,
	      "DG":2,
	     "BCMP":3,
	     "FHVWY":4,
	       "K":5,
	     "JX":8,
		 "QZ":10,
	}
	for _, str:=range s{
	  for i, v:=range maptype{
			 if strings.ContainsAny(i,string(str)){
				 c+=v
			 }
		 }
	}

	return c
}
