package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"studentapp/config"
	"studentapp/logger"
	"studentapp/student"
	"studentapp/utils"
)

func main() {
	configPath := flag.String("config", "config.json", "JSON конфигурациясының жолы")
	demo := flag.Bool("demo", false, "дайын мысалды автоматты көрсету")
	flag.Parse()
	cfg, err := config.Load(*configPath)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info(cfg.AppName)
	students := []student.Student{}
	if cfg.SeedStudents {
		initial := []student.Student{
			student.NewStudent(1, "Сүлеймен", "Алия", 2005, "IT", 2, "CS-22"),
			student.NewStudent(2, "Талғат", "Ерлан", 2006, "Math", 1, "M-11"),
			student.NewStudent(3, "Серік", "Айгерім", 2004, "IT", 3, "CS-31"),
		}
		for _, s := range initial {
			addStudent(&students, s)
		}
	}
	if *demo {
		runDemo(students, cfg)
		return
	}
	runMenu(&students)
}

func addStudent(list *[]student.Student, s student.Student) {
	if err := student.AddStudent(list, s); err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info(fmt.Sprintf("Студент қосылды: %s %s (%s)", s.FirstName, s.LastName, s.Group))
	}
}

func runDemo(students []student.Student, cfg config.Config) {
	logger.Info("Барлық студенттер")
	student.PrintAll(students)
	logger.Info("Факультет бойынша сүзу: " + cfg.DefaultFaculty)
	student.PrintAll(student.FilterByFaculty(students, cfg.DefaultFaculty))
	logger.Info(fmt.Sprintf("Курс бойынша сүзу: %d", cfg.DefaultCourse))
	student.PrintAll(student.FilterByCourse(students, cfg.DefaultCourse))
	if len(students) > 0 {
		logger.Info("Қайталанатын ID қатесін тексеру")
		addStudent(&students, students[0])
		fmt.Printf("Тізімде әлі де %d студент бар.\n", len(students))
	}
}

func runMenu(students *[]student.Student) {
	for {
		fmt.Println("\n1 — Студент қосу\n2 — Барлығын көрсету\n3 — Факультет бойынша сүзу\n4 — Курс бойынша сүзу\n0 — Шығу")
		choice := utils.ReadInt("Таңдау: ")
		if inputEnded() {
			return
		}
		switch choice {
		case 0:
			logger.Info("Бағдарлама аяқталды.")
			return
		case 1:
			id := utils.ReadInt("ID: ")
			last := utils.ReadString("Тегі: ")
			first := utils.ReadString("Аты: ")
			birth := utils.ReadInt("Туған жылы: ")
			faculty := utils.ReadString("Факультеті: ")
			course := utils.ReadInt("Курсы (1–6): ")
			group := utils.ReadString("Тобы: ")
			if inputEnded() {
				return
			}
			addStudent(students, student.NewStudent(id, last, first, birth, faculty, course, group))
		case 2:
			logger.Info("Барлық студенттер")
			student.PrintAll(*students)
		case 3:
			faculty := utils.ReadString("Факультет: ")
			if inputEnded() {
				return
			}
			student.PrintAll(student.FilterByFaculty(*students, faculty))
		case 4:
			course := utils.ReadInt("Курс: ")
			if inputEnded() {
				return
			}
			if course < 1 || course > 6 {
				logger.Error("Курс 1 мен 6 аралығында болуы керек")
				continue
			}
			student.PrintAll(student.FilterByCourse(*students, course))
		default:
			logger.Error("Мәзірден 0 мен 4 аралығындағы санды таңдаңыз.")
		}
	}
}

func inputEnded() bool {
	err := utils.Err()
	if err == nil {
		return false
	}
	if errors.Is(err, io.EOF) {
		logger.Info("Енгізу аяқталды.")
	} else {
		logger.Error(err.Error())
	}
	return true
}
