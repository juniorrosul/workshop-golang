package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Menu struct {
	coffee     string
	juice      string
	mainCourse string
	candy      string
}

func breakfastGenerator(timestamp time.Time) Menu {
	coffees := []string{"paranaense", "mineiro", "mogiano"}
	juices := []string{"laranja", "limão com açucar"}
	mainCourses := []string{"Sucrilhos", "Pão de trigo", "Pão de queijo", "Tapioca", "Croissant"}
	candies := []string{"Pão com chimia"}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	m := Menu{}

	m.mainCourse = mainCourses[r.Intn(len(mainCourses))]
	m.coffee = coffees[r.Intn(len(coffees))]
	m.juice = juices[r.Intn(len(juices))]
	m.candy = ""

	if timestamp.Weekday() == time.Thursday {
		m.mainCourse = mainCourses[5]
	}

	if timestamp.Weekday() == time.Monday || timestamp.Weekday() == time.Wednesday {
		m.candy = candies[r.Intn(len(candies))]
	}

	return m
}

func main() {
	fmt.Println(breakfastGenerator(time.Now()))
}
