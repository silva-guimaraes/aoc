package main

import (
    "bufio"
    "os"
    "fmt"
)

func is_distinct(signal string) bool {
    fmt.Println(signal)
    for i, ic := range signal {
	for j, jc := range signal {
	    if i == j {
		continue
	    } else if ic == jc {
		return false
	    }
	}
    }
    return true
}

func main() {
    file, err := os.Open("./stream.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    scanner.Scan()
    stream := scanner.Text()

    trail := 0
    head := 14 // <--

    fmt.Println(len(stream))
    for head < len(stream) {
	if is_distinct(stream[trail:head]) {
	    fmt.Println(head)
	    return
	}
	trail += 1
	head += 1
    }
    fmt.Println("nada?")

}
