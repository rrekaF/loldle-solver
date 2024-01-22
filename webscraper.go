package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Champion struct{
	Name	string
	Position	[]string 
	Resource	[]string 
	RangeType []string
	ReleaseYear uint64
	Gender	string
	Species	string
	Region	[]string
	// ReleaseYear string
}

func prepPositions(s string) []string{
	s = strings.Trim(s, "\n \t")
	prepped := strings.Split(s, " ")
	prepped = prepped[1:] // gets rid of the "Position(s)" text

	for i := 0; i < len(prepped); i++ {
		prepped[i] = strings.Trim(prepped[i], " \n")
	}
	return prepped 
}

func prepRelease(s string) uint64{
	s = strings.TrimRight(s, " ")
	prepped2, _ := strconv.ParseUint(s[len(s)-11: len(s)-7], 10, 16)

	return prepped2
}

func prepResource(s string)[]string{
	s = strings.Trim(s, "\n \t")
	prepped := strings.Split(s, " ")
	prepped = prepped[1:] // gets rid of the "Position(s)" text

	for i := 0; i < len(prepped); i++ {
		prepped[i] = strings.Trim(prepped[i], " \n")
	}
	return prepped 
	
}

func prepName(s string) string{
	prepped := strings.Split(s, " ")

	return prepped[0]
}

func prepRange(s string)[]string{
	s = strings.Trim(s, "\n \t")
	prepped := strings.Split(s, " ")
	prepped = prepped[2:] // gets rid of the "Position(s)" text

	for i := 0; i < len(prepped); i++ {
		prepped[i] = strings.Trim(prepped[i], " \n")
	}
	return prepped 
}

func prepGender(s string)string{
	s = strings.Trim(s, "\n \t")
	prepped := strings.Split(s, "\t")
	prepped = prepped[1:]

	for i := 0; i < len(prepped); i++ {
		prepped[len(prepped)-1] = strings.Trim(prepped[len(prepped)-1], "\n \t")
	}
	switch prepped[len(prepped)-1]{
	case "She/Her":
		return "Female"
	case "He/Him":
		return "Male"
	}
	// if prepped[len(prepped)-1] == "She/Her"{
	// 	return "Female"
	// }
	return prepped[len(prepped)-1]
}
func prepRegion(s string)[]string{
	s = strings.Trim(s, "\n \t")
	prepped := strings.Split(s, ",")
	// prepped = prepped[2:] // gets rid of the "Position(s)" text

	fmt.Println(prepped)
	for i := 0; i < len(prepped); i++ {
		// prepped[i] = strings.Trim(prepped[i], " \n\t")
		fmt.Printf("a%va\n", prepped[i])
	}
	// fmt.Println(prepped)

	return prepped 
}

func main(){
	var character Champion

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request){
		fmt.Println("Visiting...")
	})
	
	c.OnError(func(_ *colly.Response, err error){
		log.Println("error: ", err)
	})

	c.OnHTML("div[data-source]", func(e *colly.HTMLElement){
		switch e.Attr("data-source"){
		case "release":
			character.ReleaseYear = prepRelease(e.Text)

		case "position":
			character.Position = prepPositions(e.Text)

		case "resource":
			if len(prepResource(e.Text)) != 0{
				character.Resource = prepResource(e.Text)
			}

		case "rangetype":
			character.RangeType = prepRange(e.Text)

		case "pronoun":
			character.Gender = prepGender(e.Text)

		case "originplace":
			character.Region = prepRegion(e.Text)

		}

	})
	c.OnHTML("title", func(e *colly.HTMLElement){ //gets the name from the title
		character.Name = prepName(e.Text)
	})
	
	c.Visit("https://leagueoflegends.fandom.com/wiki/Shen/LoL")
	c.Visit("https://leagueoflegends.fandom.com/wiki/Shen")
	fmt.Println(character)
}
