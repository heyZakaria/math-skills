package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err.Error())
	}

	var hold string
	var slc []float64
	for _, v := range f {
		if v == 10 {
			i, err := strconv.ParseFloat(hold, 64)
			if err != nil {
				log.Println(err)
			} else {

				slc = append(slc, i)
				hold = ""
			}
			continue
		}

		hold += string(v)
	}

	if hold != "" {
		i, err := strconv.ParseFloat(hold, 64)
		if err != nil {
			log.Println(err)
		}
		slc = append(slc, i)
		hold = ""
	}

	fmt.Println(slc)
	a := Average(slc)
	m := Median(slc)
	v := Variance(slc, a)
	sd := StandardDeviation(slc, a)
	fmt.Println("Average:",a)
	fmt.Println("Median:",m)
	fmt.Println("Vanriance:",v)
	fmt.Println("Standard Deviation:",sd)
}

func Average(a []float64) float64 {
	var res float64
	for _, v := range a {
		res += v
	}

	return math.Round(res / float64(len(a)))
}

func Median(s []float64) float64 {
	sort.Float64s(s)
	var res float64
	if len(s)%2 == 0 {
		res = (s[len(s)/2-1] + s[len(s)/2]) / 2
	} else {
		res = s[len(s)/2]
	}

	return math.Round(res)
}

func Variance(s []float64, a float64) float64 {
	var res float64
	for _, v := range s {
		res += (v - a) * (v - a)
	}
	res = res / float64(len(s))

	return math.Round(res)
}


func StandardDeviation(s []float64, a float64) float64 {
	var res float64
	for _, v := range s {
		res += (v - a) * (v - a)
	}
	res = res / float64(len(s))

	return math.Round(math.Sqrt(res))
}