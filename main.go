package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/confetty/confetti"
)

type Question struct {
	Text     string
	Options  []string
	Correct  int
	Answered int
}

func runConfetti() {
	var wg sync.WaitGroup

	wg.Add(1)
	model := confetti.InitialModel()
	t := tea.NewProgram(model)

	go t.Run()

	go func(t *tea.Program) {
		defer wg.Done()
		for range time.Tick(time.Second) {
			t.Send(tea.KeyMsg{})
		}
	}(t)

	wg.Wait()
}

func waitInput() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}

func runQuiz() {
	questions := []Question{
		{
			Text:    "Чему будет равно это условие? \n[] == ![];",
			Options: []string{"True", "False"},
			Correct: 0,
		},
		{
			Text:    "Какая строка получится в результате этой конкатенации строк? \"foo\" + +\"bar\";",
			Options: []string{"fooNaN", "foobar"},
			Correct: 0,
		},
		{
			Text:    "Чему будет равно это условие? NaN === NaN;",
			Options: []string{"True", "False"},
			Correct: 1,
		},
		{
			Text:    "Чему будет равно это выражение? 0.1 + 0.2;",
			Options: []string{"0.3", "0.3000000004"},
			Correct: 1,
		},
		{
			Text:    "Заблокирована ли кнопка или нет <button disabled=\"false\">",
			Options: []string{"Да", "Нет"},
			Correct: 0,
		},
		{
			Text:    "Чему будет равно это условие? Math.min() > Math.max();",
			Options: []string{"True", "False"},
			Correct: 0,
		},
		{
			Text:    "Сплит по строке \"\".split(\"\"); возвращает пустой массив []. Как изменить код, чтобы он вернул массив с пустой строкой [\"\"]?",
			Options: []string{"\"\".split(\" \")", "\" \".split(\"\")"},
			Correct: 0,
		},
		{
			Text:    "Кто твой любимый музыкальный исполнитель",
			Options: []string{"Неправильный ответ", "Егор Крид"},
			Correct: 1,
		},
	}

	score := 0

	for i, q := range questions {
		fmt.Println("Вопрос", i+1, ":", q.Text)
		for j, option := range q.Options {
			fmt.Println(j+1, ":", option)
		}

		var answer int
		fmt.Print("Напиши ответ (1-2): ")
		fmt.Scanln(&answer)

		if answer == q.Correct+1 {
			fmt.Println("Верно!")
			score++
		} else {
			fmt.Println("Неверно. Правильный ответ:", q.Options[q.Correct])
		}

		fmt.Println()

		time.Sleep(time.Second * 3)
	}
}

func main() {
	fmt.Println("Дима, привет!")
	fmt.Println()
	<-time.After(time.Second * 3)

	fmt.Println("У тебя сегодня день рождения?")
	fmt.Println()
	<-time.After(time.Second * 3)

	fmt.Println("Кхмм...")
	fmt.Println()
	<-time.After(time.Second * 3)

	fmt.Println("Просто так мы подарок не дадим ⊂(◉‿◉)つ")
	fmt.Println()
	<-time.After(time.Second * 4)

	fmt.Println("Придется решить несколько задачек по любимому JavaScript!")
	fmt.Println()
	<-time.After(time.Second * 5)

	fmt.Println("Ты готов? Напиши Да")

	waitInput()

	runQuiz()

	runConfetti()
}
