package space
//Planet is used to determine on which planet we are calculating the age
type Planet string
const earthyears= float64(31557600)
//Age is a function used to calculate the age based upon the given planet
func Age(ageinseconds float64,planet Planet)float64{
	planets:=map[Planet]float64{
		"Earth":earthyears*1,
		"Mercury":earthyears*0.2408467,
		"Venus":earthyears*0.61519726,
		"Mars":earthyears*1.8808158,
		"Jupiter":earthyears*11.862615,
		"Saturn":earthyears*29.447498,
		"Uranus":earthyears*84.016846,
		"Neptune":earthyears*164.79132,
	}
	
 return ageinseconds/planets[planet]
}
  



	


 
