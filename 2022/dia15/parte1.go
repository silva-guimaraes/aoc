package main

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"os"
	"regexp"
	"strconv"
)


type sensor struct {
    x, y int
    bx, by int
    beacon_distance int
}

// 3245601 too low
// 2887169 too low

func uniqueInts(lists [][]int) []int {
  uniqueIntMap := make(map[int]bool)
  for _, list := range lists {
    for _, n := range list {
      uniqueIntMap[n] = true
    }
  }
  uniqueInts := make([]int, 0, len(uniqueIntMap))
  for n := range uniqueIntMap {
    uniqueInts = append(uniqueInts, n)
  }
  sort.Ints(uniqueInts)
  return uniqueInts
}

func contains(list []sensor, n int) bool {

    for i := range list {
	if list[i].bx == n {
	    return true
	}
    }
    return false
}

func difference(a []int, b []sensor) []int {
    var diff []int

    for _, n := range a {
	if !contains(b, n) {
	    diff = append(diff, n)
	}
    }
    return diff
}


func main() {
    file, err := os.Open("./sensors.txt")
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    r, _ := regexp.Compile("\\d+")

    var sensors []sensor

    for scanner.Scan() {
	extract := r.FindAllString(scanner.Text(), -1)

	sensor_x, _ := strconv.Atoi(extract[0])
	sensor_y, _ := strconv.Atoi(extract[1])
	beacon_x, _ := strconv.Atoi(extract[2])
	beacon_y, _ := strconv.Atoi(extract[3])

	// sensores usam taxicap distance
	beacon_distance := int(
	    math.Abs(float64(sensor_y) - float64(beacon_y)) +
	    math.Abs(float64(sensor_x) - float64(beacon_x)))

	sensor := sensor{sensor_x, sensor_y, beacon_x, beacon_y, beacon_distance}
	sensors = append(sensors, sensor)
    }
    file.Close()

    // hashmap pra evitar que duplicações ocorram
    var intersections [][]int

    const y_slice int = 2000000
    // const y_slice int = 10

    fmt.Println(len(sensors), y_slice)

    // iterar por todos os sensores e encontrar as subseções das areas que façam interseção com o Y alvo
    for i := range sensors {

	sensor := sensors[i]
	// fmt.Println("sensor", i + 1, sensor.y, y_slice, sensor.y + sensor.beacon_distance, "\t", sensor.y - sensor.beacon_distance, y_slice, sensor.y)
	// se camada faz interseção com area do sensor 
	if sensor.y <= y_slice && y_slice <= sensor.y + sensor.beacon_distance ||
	sensor.y - sensor.beacon_distance <= y_slice && y_slice <= sensor.y {

	    // distancia entre camada e posição do sensor seria o Y de um beacon imaginario
	    // a diferença de Y e a distancia do beacon pro sensor seria o X do beacon imaginario
	    beacon_x := sensor.beacon_distance - int(math.Abs(float64(sensor.y) - float64(y_slice)))
	    end := sensor.x + beacon_x
	    start := sensor.x - beacon_x
	    // fmt.Println(i + 1, "dentro")
	    // fmt.Println(sensor.x, sensor.y, sensor.beacon_distance, "\t", start, end)

	    // range entre as possiveis posições de X do beacon imaginario que não estejam fora da area
	    // do sensor
	    var intersection []int
	    for x := start; x <= end; x++ {
		intersection = append(intersection, x)
	    }

	    intersections = append(intersections, intersection)
	     // fmt.Println(start,end, beacon_x)
	}
    }

    // remover posições que ja possuam um beacon ou sensor
    // for i := range sensors {
    //     if sensors[i].y == y_slice{
    //         delete(row, sensors[i].x)
    //     }
    //     if sensors[i].by == y_slice{
    //         delete(row, sensors[i].bx)
    //     }
    // }

    // fmt.Println(len(intersections))
    // fmt.Println(intersections)

    // for i := range intersections {
    //     fmt.Println(len(intersections[i]))
    // }

    var filtered_beacons []sensor
    for i := range sensors {
	if sensors[i].by == y_slice {
	    filtered_beacons = append(filtered_beacons, sensors[i])
	}
    }

    diff := difference(uniqueInts(intersections), filtered_beacons)
    fmt.Println(len(diff))
    // fmt.Println(diff)


    // var visualize []int

    // for i := range row {
    //     visualize = append(visualize, i)
    // }

    // sort.Slice(visualize, func(i, j int) bool {
    //     return visualize[i] < visualize[j]
    // })

    // for i := range visualize {
    //     fmt.Println(visualize[i])
    // }
}
