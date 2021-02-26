
import "fmt"
import "time"
import "sync"

var next int = 0
var semaphore sync.Mutex

func even() int {
  next++
  time.Sleep(time.Millisecond * 10)
  next++
  return next
}

func print(n string) {

  for i := 1; i <= 10; i++ {
    fmt.Println(n, "(", i, ")", even())
  }
  semaphore.Unlock()
}

func main() {

  semaphore.Lock()
  go print("Thread 1:")
  semaphore.Lock()
  go print("Thread 2:")

  var input string
  fmt.Scanln(&input)
}