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

// Name returns the employee's name and implements the Employee
// interface Name() method
func (g GenericEmployee) Name() string {
	return g.FullName
}

// Email returns the employee's email and implements the Employee
// interface Email() method
func (g GenericEmployee) Email() string {
	return g.EmailAddress
}

// WorkDetails returns the number of weeks worked and hours worked per week
// and implements the Employee interface WorkDetails() method
func (g GenericEmployee) WorkDetails() (int, int) {
	return g.WeeksWorkedAtCompany, g.HoursWorkedPerWeek
}

// Skills returns the skills of the employee and implements the
// Employee interface Skills() method.  This method is overridden by
// the Writer, Artist, and Maker structs so we just return an empty
// slice here.  We normally would never call this method unless we
// actually create an instance of GenericEmployee, but we include it
// here for completeness and so ":GoImplements" can test if
// GenericEmployee implements the rest of the Employee interface.
func (g GenericEmployee) Skills() []string {
	panic("Skills() method not implemented on GenericEmployee and should not be called.  It should be overridden by e.g. Writer, Artist, and Maker.")
	return []string{}
}

// Employee interface defines the methods that must be implemented by
// all professions.
type Employee interface {
	Name() string
	Email() string
	WorkDetails() (int, int)
	Skills() []string
}

// Maker is a struct that represents someone who makes things in
// projects.  Because they are also an employee, this struct
// implements the Employee interface.
type Maker struct {
	GenericEmployee
	ProjectsWorkedOn []string
}

// Skills returns the projects worked on by the Maker
func (m Maker) Skills() []string {
	return m.ProjectsWorkedOn
}

// Writer is a struct that represents someone who writes things that
// fit into genres.  Because they are also an employee, this struct
// implements the Employee interface.
type Writer struct {
	GenericEmployee
	GenresWritten []string
}

// Skills returns the genres written by the Writer
func (w Writer) Skills() []string {
	return w.GenresWritten
}

// Artist is a struct that represents someone who creates art in
// various styles.  Because they are also an employee, this struct
// implements the Employee interface.
type Artist struct {
	GenericEmployee
	ArtStyles []string
}

// Skills returns the art styles of the Artist
func (a Artist) Skills() []string {
	return a.ArtStyles
}

// PrintEmployeeDetails is a reusable function to display employee
// details.  It takes an Employee interface as an argument so it can
// be used with any profession that implements the Employee interface.
// We have to write this as a function rather than a method on the
// Employee interface because Go does not support methods on interfaces.
// (In more formal terms, Go does not support interfaces as receivers;
// receivers need to be concrete types.)
func PrintEmployeeDetails(e Employee) {
	fmt.Printf("Name: %s\n", e.Name())
	fmt.Printf("Email: %s\n", e.Email())
	weeks, hours := e.WorkDetails()
	fmt.Printf("Worked for %d weeks, %d hours per week\n", weeks, hours)
	fmt.Printf("Skills: %v\n\n", e.Skills())
}

// MostSenior returns the most senior employee based on weeks worked.
// If two employees have worked the same number of weeks, it will
// return the first one it encounters.  It takes a slice of Employee
// interfaces as an argument so it can be used with any profession
// that implements the Employee interface. It will return something
// that also implements the Employee interface. As with
// PrintEmployeeDetails, we need to write this as a function rather
// than a method.
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
		// Because the GenericEmployee struct is embedded in the Writer struct,
		// we initialize it here as a nested struct literal.  Note
		// that the left-hand side of the colon is the field name in
		// the Writer struct, and the right-hand side is the field
		// value.
		GenericEmployee: GenericEmployee{
			FullName:             "Charlie",
			EmailAddress:         "charlie@example.com",
			WeeksWorkedAtCompany: 104,
			HoursWorkedPerWeek:   6,
		},
		// The []string{...} syntax is a slice literal.  It creates a
		// slice of strings.  We use it here to initialize the GenresWritten
		// field of the Writer struct.
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

	// Create a list of employees from a slice literal.
	employees := []Employee{writer, artist, maker}

	fmt.Println("Most Senior Employee:")
	PrintEmployeeDetails(MostSenior(employees))

}
