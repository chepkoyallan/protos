package main

import "fmt"

type Job struct{
	title string
	workplace string
} 

//iota
const(
	dip = iota
	undergrad;
	grad;
	postgrad;
)

func main(){
	const personalWebsite string = "www.allano.com";
	// Value types
	var githubLink *string = new(string);
	//dereference operator
	*githubLink = "https://github.com/chepkoyallan";
	var age int32 = 29;
	headerFinished := true;
	firstName := "Allano";
	secondName := "Kiplang'at";
	lastName := "Chepkoy";
	fullName :=[3]string{firstName, secondName, lastName};
	addressLocation := []string{"Kaminoge", "Takatsu"};
	schoolDetails := map[string]string{
		"university": "university of Nairobi",
	}
	

	//Address of
	addressOfFirstName := &firstName; 
	placeOfWork := Job{title: "software engineer", workplace: "Rakuten mobile"}

	linkedinProfile := "https://www.linkedin.com/in/allankiplangat/";
	fmt.Println(
		firstName, 
		secondName, 
		linkedinProfile, 
		age, 
		headerFinished, 
		*githubLink, 
		*addressOfFirstName, 
		personalWebsite, 
		dip,
		fullName,
		addressLocation,
		schoolDetails,
		placeOfWork,
	)
}