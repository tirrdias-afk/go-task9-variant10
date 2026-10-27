package logger

import (
	"github.com/fatih/color"
	"time"
)

// INFO жасыл, ERROR қызыл түспен көрсетіледі.
// Файлға бағытталған шығыста кітапхана түстерді автоматты өшіреді.
func Info(msg string) {
	color.New(color.FgGreen).Printf("[INFO] %s — %s\n", timestamp(), msg)
}

func Error(msg string) {
	color.New(color.FgRed).Printf("[ERROR] %s — %s\n", timestamp(), msg)
}

func timestamp() string { return time.Now().Format("2006-01-02 15:04:05") }
