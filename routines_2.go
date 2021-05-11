
func f(n int) {
	
	for i := 0; i < 10; i++ {
		fmt.Println(n, ": ", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	
	for i := 0; i < 10; i++ { // create 10 threads...
		go f(i)
	}
	
	var input string
	fmt.Scanln(&input)
}

 
//	Output:
//	order cannot be forseen
