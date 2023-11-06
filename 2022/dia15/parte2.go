package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "regexp"
    "strconv"
)


type Pos struct {
    x, y int
}


type sensor struct {
    x, y int
    bx, by int
    beacon_distance int
}
const FILE = "sensors.txt"
// const MAX = 20 
const MAX = 4000000 // importante !!!!!!

func (p Pos) inside() bool {
    return p.x >= 0 && p.x <= MAX && p.y >= 0 && p.y <= MAX
}


func main() {
    file, err := os.Open(FILE)
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    r, _ := regexp.Compile("-?\\d+")

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
            math.Abs(float64(sensor_x) - float64(beacon_x)),
        )

        sensor := sensor{sensor_x, sensor_y, beacon_x, beacon_y, beacon_distance}
        sensors = append(sensors, sensor)
    }
    file.Close()

    // gera o perimetro exterior de todos os sensores e procura qual
    // deles não se encontra dentro da area de nenhum outro sensor

    for _, sensor := range sensors {
        edge := sensor.beacon_distance+1 
        var pos []Pos
        // gerar perimetro
        for a := 0; a < sensor.beacon_distance; a++ {

            A := Pos{x: sensor.x + a, y: sensor.y - edge + a}
            B := Pos{x: sensor.x + edge - a, y: sensor.y + a}
            C := Pos{x: sensor.x - a, y: sensor.y + edge - a}
            D := Pos{x: sensor.x - edge + a, y: sensor.y - a}

            if A.inside() { pos = append(pos, A) }
            if B.inside() { pos = append(pos, B) }
            if C.inside() { pos = append(pos, C) }
            if D.inside() { pos = append(pos, D) }

        }
        for _, p := range pos {
            outside := true
            for _, s := range sensors {
                d := int(math.Abs(
                    float64(p.x - s.x)) + math.Abs(float64(p.y - s.y),
                ))
                if d <= s.beacon_distance {
                    outside = false
                    break
                }
            }
            if outside {
                fmt.Println(p.x * 4000000 + p.y)
                return;
            }

        }
    }
}
