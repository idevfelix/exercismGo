package hamming
import ("errors")

func Distance(a, b string) (int, error) {

    countErrors := 0

    if(len(a) != len(b)) {
        return 0, errors.New("The strings should be have same length")
    }
    
    for i, char:= range a {
        if(string(b[i]) != string(char)){
            countErrors +=1
        }
    }
    return countErrors, nil
}
