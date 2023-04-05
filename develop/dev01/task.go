package dev01

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

type INtp interface {
	GetTime() (time.Time, error)
}

type Ntp struct {
	timeHost string
}

func (n *Ntp) GetTime() (time.Time, error) {
	return ntp.Time(n.timeHost)
}

func NewNtp(timeHost string) INtp {
	return &Ntp{
		timeHost: timeHost,
	}
}

type Timer struct {
	NtpClient INtp
}

func NewTimer(ntpClient INtp) *Timer {
	return &Timer{
		NtpClient: ntpClient,
	}
}

func (t *Timer) GetTime() (time.Time, error) {
	return t.NtpClient.GetTime()
}

func (t *Timer) FormatTime(currentTime time.Time) string {
	return currentTime.Format(time.RFC3339Nano)
}

func Task01() {
	myNtp := NewNtp("pool.ntp.org")
	timer := NewTimer(myNtp)

	currentTime, err := timer.GetTime()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ошибка получения времени:", err) // обрабатываем ошибку
		os.Exit(1)
	}

	formattedTime := timer.FormatTime(currentTime)
	fmt.Printf("Точное время: %s\n", formattedTime)
}
