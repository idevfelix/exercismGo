package leap

func isDivisibleByFour(year int) bool {
    return (year % 4) == 0
}

func isDivisibleByOneHundred(year int) bool {
    return (year % 100) == 0
}

func isDivisibleByFourHundred(year int) bool {
    return (year % 400) == 0
}

func IsLeapYear(year int) bool {
	if(isDivisibleByFour(year)) {        
        if(!isDivisibleByOneHundred(year)){
        	return true            
        }        
        if(isDivisibleByOneHundred(year) && isDivisibleByFourHundred(year)) {            
            return true            
        }
    }
    return false
}
