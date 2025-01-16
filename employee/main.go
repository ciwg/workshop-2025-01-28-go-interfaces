package main

import "fmt"

type GenericEmployee struct {
	FullName             string
	EmailAddress         string
	WeeksWorkedAtCompany int
	HoursWorkedPerWeek   int
}

func (g GenericEmployee) Name() string {
	return g.FullName
}

func (g GenericEmployee) Email() string {
	return g.EmailAddress
}

func (g GenericEmployee) WorkDetails() (int, int) {
	return g.WeeksWorkedAtCompany, g.HoursWorkedPerWeek
}

// Employee interface defines the contract for all professions
type Employee interface {
	Name() string            // Get the employee's name
	Email() string           // Get the employee's email
	WorkDetails() (int, int) // Get weeks worked and hours worked per week
	Skills() []string        // Get a list of skills
}

/*
// Maker struct implements the Employee interface
type Maker struct {
	FullName         	string
	EmailAddress     	string
	WeeksWorkedAtCompany int
	HoursWorkedPerWeek   int
	ProjectsWorkedOn 	[]string
}

// Implement the Employee interface for Maker
func (m Maker) Name() string {
	return m.FullName
}

func (m Maker) Email() string {
	return m.EmailAddress
}

func (m Maker) WorkDetails() (int, int) {
	return m.WeeksWorkedAtCompany, m.HoursWorkedPerWeek
}

func (m Maker) Skills() []string {
	return m.ProjectsWorkedOn
}
*/
// Writer struct implements the Employee interface
type Writer struct {
	GenericEmployee
	GenresWritten []string
}

// Implement the Employee interface for Writer
func (w Writer) Skills() []string {
	return w.GenresWritten
}

// Artist struct implements the Employee interface
type Artist struct {
	FullName             string
	EmailAddress         string
	WeeksWorkedAtCompany int
	HoursWorkedPerWeek   int
	ArtStyles            []string
}

// Implement the Employee interface for Artist
func (a Artist) Name() string {
	return a.FullName
}

func (a Artist) Email() string {
	return a.EmailAddress
}

func (a Artist) WorkDetails() (int, int) {
	return a.WeeksWorkedAtCompany, a.HoursWorkedPerWeek
}

func (a Artist) Skills() []string {
	return a.ArtStyles
}

// PrintEmployeeDetails is a reusable function to display employee details
func PrintEmployeeDetails(e Employee) {
	fmt.Printf("Name: %s\n", e.Name())
	fmt.Printf("Email: %s\n", e.Email())
	weeks, hours := e.WorkDetails()
	fmt.Printf("Worked for %d weeks, %d hours per week\n", weeks, hours)
	fmt.Printf("Skills: %v\n\n", e.Skills())
}

func main() {
	// Create a Writer instance
	writer := Writer{
		GenericEmployee: GenericEmployee{
			FullName:             "Charlie",
			EmailAddress:         "charlie@example.com",
			WeeksWorkedAtCompany: 104,
			HoursWorkedPerWeek:   6,
		},
		GenresWritten: []string{"Fiction", "Non-Fiction", "Poetry"},
	}

	/*	// Create a Maker instance
			maker := Maker{
		    	FullName:         	"Alice",
		    	EmailAddress:     	"alice@example.com",
		    	WeeksWorkedAtCompany: 52,
		    	HoursWorkedPerWeek:   40,
		    	ProjectsWorkedOn: 	[]string{"Photography", "Cutting Boards", "Candles"},
			}*/

	// Create an Artist instance
	artist := Artist{
		FullName:             "Bob",
		EmailAddress:         "bob@example.com",
		WeeksWorkedAtCompany: 26,
		HoursWorkedPerWeek:   9,
		ArtStyles:            []string{"Painting", "Sculpture"},
	}

	// Print details of both employees
	PrintEmployeeDetails(writer)
	PrintEmployeeDetails(artist)
	/*	PrintEmployeeDetails(maker)*/

}
