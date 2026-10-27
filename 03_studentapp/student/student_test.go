package student

import (
	"testing"
	"time"
)

func example(id int) Student { return NewStudent(id, "Тегі", "Аты", 2005, "IT", 2, "CS-22") }

func TestAddAndDuplicate(t *testing.T) {
	list := []Student{}
	if err := AddStudent(&list, example(1)); err != nil {
		t.Fatal(err)
	}
	if err := AddStudent(&list, example(1)); err == nil {
		t.Fatal("duplicate accepted")
	}
	if len(list) != 1 {
		t.Fatalf("unexpected count: %d", len(list))
	}
}

func TestInvalidStudentDoesNotChangeList(t *testing.T) {
	cases := map[string]func(*Student){
		"id":          func(s *Student) { s.ID = 0 },
		"first_name":  func(s *Student) { s.FirstName = " " },
		"last_name":   func(s *Student) { s.LastName = "" },
		"faculty":     func(s *Student) { s.Faculty = " " },
		"group":       func(s *Student) { s.Group = "" },
		"old_year":    func(s *Student) { s.BirthYear = 1899 },
		"future_year": func(s *Student) { s.BirthYear = time.Now().Year() + 1 },
		"low_course":  func(s *Student) { s.Course = 0 },
		"high_course": func(s *Student) { s.Course = 7 },
	}
	for name, change := range cases {
		t.Run(name, func(t *testing.T) {
			s := example(2)
			change(&s)
			list := []Student{example(1)}
			if AddStudent(&list, s) == nil || len(list) != 1 {
				t.Fatal("invalid student accepted")
			}
		})
	}
	if AddStudent(nil, example(1)) == nil {
		t.Fatal("nil pointer accepted")
	}
}

func TestFilters(t *testing.T) {
	a, b, c := example(1), example(2), example(3)
	b.Faculty, b.Course = "Math", 1
	c.Course = 3
	list := []Student{a, b, c}
	if got := FilterByFaculty(list, " it "); len(got) != 2 || got[0].ID != 1 || got[1].ID != 3 {
		t.Fatalf("wrong faculty result: %v", got)
	}
	if got := FilterByCourse(list, 2); len(got) != 1 || got[0].ID != 1 {
		t.Fatalf("wrong course result: %v", got)
	}
	if len(FilterByFaculty(list, "Unknown")) != 0 || len(FilterByCourse(nil, 1)) != 0 {
		t.Fatal("empty filter failed")
	}
	if len(list) != 3 {
		t.Fatal("filter changed source")
	}
}
