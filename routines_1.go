
import
{
	"fmt"
	"time"
	"math/rand"
}

func f(n int) {
	
	for i := 0; i < 10; i++ { 
		fmt.Println(n, ": ", i)
	}
}

func main() {

	go f(0) // go-routine (Thread) runs in background
	
	var input string
	fmt.Scanln(&input)
}


/*
 *	Output:
 *	0: 1
 * 	0: 2
 *	0: 3
 *	0: 4
 *	0: 5
 *	
 *	etc.
 */ 
