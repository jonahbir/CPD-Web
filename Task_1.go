package main

import (
    "fmt"
    "strings"
)

func main() {
    var studentName string
    var subjectCount int

    fmt.Print("Enter your name: ")
    fmt.Scanln(&studentName)

    fmt.Print("Enter number of subjects: ")
    fmt.Scanln(&subjectCount)

    // Store subject names and grades
    grades := make(map[string]float64)

    for i := 0; i < subjectCount; i++ {
        var subject string
        var grade float64

        fmt.Printf("\nEnter name of subject #%d: ", i+1)
        fmt.Scanln(&subject)

        // Loop to ensure valid grade (0–100)
        for {
            fmt.Printf("Enter grade for %s (0–100): ", subject)
            fmt.Scanln(&grade)

            if grade >= 0 && grade <= 100 {
                break // valid grade
            } else {
                fmt.Println("❌ Invalid grade. Please enter a number between 0 and 100.")
            }
        }

        grades[subject] = grade
    }

    // Calculate average
    average := calculateAverage(grades)

    // Display result
    fmt.Println("\n📄 Report Card:")
    fmt.Printf("Student Name: %s\n", strings.Title(studentName))
    fmt.Println("Subjects and Grades:")

    for subject, grade := range grades {
        fmt.Printf("- %s: %.2f\n", strings.Title(subject), grade)
    }

    fmt.Printf("\n📊 Average Grade: %.2f\n", average)
}

// Function to calculate average from map
func calculateAverage(grades map[string]float64) float64 {
    var total float64
    for _, grade := range grades {
        total += grade
    }
    return total / float64(len(grades))
}
