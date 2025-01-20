// Description: This program demonstrates the use of interfaces in Go.
// It defines an Employee interface with methods that must be implemented by all professions.
// It then defines three professions: Writer, Artist, and Maker, each implementing the Employee interface.

package main

import "fmt"

// GenericEmployee struct needed to implement the Employee interface
type GenericEmployee struct {
	FullName             string // Get the employee's name
	EmailAddress         string // Get the employee's email
	WeeksWorkedAtCompany int    // Get weeks worked at the company
	HoursWorkedPerWeek   int    // Get hours worked per week

}

// Implement the Employee interface for GenericEmployee
func (g GenericEmployee) Name() string {
	return g.FullName
}

// Implement the Employee interface for GenericEmployee
func (g GenericEmployee) Email() string {
	return g.EmailAddress
}

// Implement the Employee interface for GenericEmployee
func (g GenericEmployee) WorkDetails() (int, int) {
	return g.WeeksWorkedAtCompany, g.HoursWorkedPerWeek
}

// Implement the Employee interface for GenericEmployee
type Employee interface {
	Name() string
	Email() string
	WorkDetails() (int, int)
	Skills() []string
}

// Maker struct implements the Employee interface
type Maker struct {
	GenericEmployee
	ProjectsWorkedOn []string
}

// Implement the Employee interface for Maker
func (m Maker) Skills() []string {
	return m.ProjectsWorkedOn
}

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
	GenericEmployee
	ArtStyles []string
}

// Implement the Employee interface for Artist
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

// MostSenior returns the most senior employee based on weeks worked
func MostSenior(eList []Employee) Employee {
	mostSenior := eList[0]
	mostSeniorWeeks, _ := mostSenior.WorkDetails() // Extract the number of weeks worked
	for _, e := range eList {
		weeks, _ := e.WorkDetails() // Compare based on weeks worked
		if weeks > mostSeniorWeeks {
			mostSenior = e
			mostSeniorWeeks = weeks
		}
	}
	return mostSenior
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

	// Create a Maker instance
	maker := Maker{
		GenericEmployee: GenericEmployee{
			FullName:             "Alice",
			EmailAddress:         "alice@example.com",
			WeeksWorkedAtCompany: 52,
			HoursWorkedPerWeek:   10,
		},
		ProjectsWorkedOn: []string{"Photography", "Cutting Boards", "Candles"},
	}

	// Create an Artist instance
	artist := Artist{
		GenericEmployee: GenericEmployee{
			FullName:             "Bob",
			EmailAddress:         "bob@example.com",
			WeeksWorkedAtCompany: 26,
			HoursWorkedPerWeek:   9,
		},
		ArtStyles: []string{"Painting", "Sculpture"},
	}

	// Print details of both employees
	PrintEmployeeDetails(writer)
	PrintEmployeeDetails(artist)
	PrintEmployeeDetails(maker)

	// Create a list of employees
	employees := []Employee{writer, artist, maker}
	fmt.Println("Most Senior Employee:")
	PrintEmployeeDetails(MostSenior(employees))

}
