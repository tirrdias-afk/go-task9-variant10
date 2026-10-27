package student

import (
	"fmt"
	"strings"
	"time"
)

type Student struct {
	ID        int
	LastName  string
	FirstName string
	BirthYear int
	Faculty   string
	Course    int
	Group     string
}

// ID қолмен берілгендіктен, қайталанатын ID қатесін де тексеруге болады.
func NewStudent(id int, lastName, firstName string, birthYear int,
	faculty string, course int, group string) Student {
	return Student{
		ID: id, LastName: strings.TrimSpace(lastName),
		FirstName: strings.TrimSpace(firstName), BirthYear: birthYear,
		Faculty: strings.TrimSpace(faculty), Course: course,
		Group: strings.TrimSpace(group),
	}
}

func AddStudent(list *[]Student, s Student) error {
	if list == nil {
		return fmt.Errorf("студенттер тізімі берілмеді")
	}
	if s.ID <= 0 {
		return fmt.Errorf("ID оң сан болуы керек")
	}
	for _, existing := range *list {
		if existing.ID == s.ID {
			return fmt.Errorf("студент ID %d бұрыннан бар", s.ID)
		}
	}
	if strings.TrimSpace(s.LastName) == "" || strings.TrimSpace(s.FirstName) == "" ||
		strings.TrimSpace(s.Faculty) == "" || strings.TrimSpace(s.Group) == "" {
		return fmt.Errorf("аты, тегі, факультеті және тобы бос болмауы керек")
	}
	if s.BirthYear < 1900 || s.BirthYear > time.Now().Year() {
		return fmt.Errorf("туған жыл 1900 мен %d аралығында болуы керек", time.Now().Year())
	}
	if s.Course < 1 || s.Course > 6 {
		return fmt.Errorf("курс 1 мен 6 аралығында болуы керек")
	}
	*list = append(*list, s)
	return nil
}

func PrintAll(list []Student) {
	if len(list) == 0 {
		fmt.Println("Студент табылмады.")
		return
	}
	for _, s := range list {
		fmt.Printf("ID %d | %s %s | %d ж. | %s | %d курс | %s\n",
			s.ID, s.LastName, s.FirstName, s.BirthYear, s.Faculty, s.Course, s.Group)
	}
	fmt.Printf("Барлығы: %d студент\n", len(list))
}

func FilterByFaculty(list []Student, faculty string) []Student {
	result := []Student{}
	for _, s := range list {
		if strings.EqualFold(strings.TrimSpace(s.Faculty), strings.TrimSpace(faculty)) {
			result = append(result, s)
		}
	}
	return result
}

func FilterByCourse(list []Student, course int) []Student {
	result := []Student{}
	for _, s := range list {
		if s.Course == course {
			result = append(result, s)
		}
	}
	return result
}
