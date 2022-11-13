package cmd

/*
This package provides content that is responsible for displaying information in the app windows
*/

import (
	"GMG/db"
	"GMG/grpc"
	"GMG/language"
	"GMG/number"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"time"
)

var (
	Game        = app.New()
	Count       = 1
	StartWindow = Game.NewWindow("Welcome")
	GameGrid    = Game.NewWindow("Game")
	Start       = time.Now()
	TextLabel   = widget.NewLabel("")
	EndText     = widget.NewLabel("")
	Array       [25]int
	EndTime     float32
	InputName   = widget.NewEntry()
	InputEmail  = widget.NewEntry()
)

// TimeSet is used to calculate the time spent to find numbers.
func TimeSet(Start time.Time) time.Duration {
	Duration := time.Since(Start)
	return Duration
}

// StartApp determines the dimensions of the main window and the dimension of the array for the playing field.
func StartApp() {
	StartWindow.Resize(fyne.NewSize(233, 280))
	for i := 0; i <= 24; i++ {
		Array[i] = i + 1
	}
	LanguageChange()
	StartWindow.ShowAndRun()
}

// LanguageChange allows you to change the language of the program interface
func LanguageChange() {
	StartWindow.CenterOnScreen()
	StartWindow.SetContent(container.NewVBox(
		language.LabelLanguage,
		widget.NewButton(language.EnglishButton, func() {
			language.UpdateLanguage()
			StartWindowContent()
		}),
		widget.NewButton(language.RussianButton, func() {
			StartWindowContent()
		})))
}

func StartWindowContent() {
	GameContentUpdater()
	StartText := widget.NewLabel(language.StartWindowText)
	StartWindow.CenterOnScreen()
	StartWindow.SetContent(container.NewVBox(
		StartText,
		widget.NewButton(language.StartButtonLabel, func() {
			SetUpdatedContent()
		})))
}

func NumberChecker(number int) {
	if number == 25 && Count == 25 {
		EndTime = float32(TimeSet(Start).Seconds())
		EndText.SetText(fmt.Sprintf(language.ResultText, EndTime))
		TextLabel.SetText("")
		EndContentUpdater()
	}
	if number != Count {
		TextLabel.SetText("")
		ErrorContentUpdater()
	} else {
		Count++
	}
}

func EndContentUpdater() {
	StartWindow.SetContent(container.NewVBox(
		EndText,
		widget.NewButton(language.SendScoreToDb, func() {
			ScoreToDbContent()
		}),
		widget.NewButton(language.RestartButtonLabel, func() {
			Count = 1
			StartWindowContent()
		}),
		widget.NewButton(language.ExitButtonLabel, func() {
			Game.Quit()
		})))
}

func ScoreToDbContent() {
	var a string
	InputEmail.SetPlaceHolder("Введите эл.почту...")
	InputName.SetPlaceHolder("Введите имя...")
	StartWindow.SetContent(container.NewVBox(
		EndText,
		InputName,
		InputEmail,
		widget.NewButton(language.SendScoreToDb, func() {
			if db.EmailCheck(InputEmail.Text) {

			}
			grpc.ConnectGRPC(InputName.Text, EndTime)
			a = InputName.Text
			fmt.Println(a, EndTime)
			time.Sleep(2 * time.Second)
			ShowTopContent()
		}),
		widget.NewButton(language.RestartButtonLabel, func() {
			Count = 1
			StartWindowContent()
		}),
		widget.NewButton(language.ExitButtonLabel, func() {
			Game.Quit()
		})))
}

func ShowTopContent() {
	StartWindow.SetContent(container.NewVBox(
		db.GetAllFromDb(),
		widget.NewButton(language.RestartButtonLabel, func() {
			Count = 1
			StartWindowContent()
		}),
		widget.NewButton(language.ExitButtonLabel, func() {
			Game.Quit()
		})))
}

func ErrorContentUpdater() {
	ErrorText := widget.NewLabel(language.ErrorWindowText)
	StartWindow.SetContent(container.NewVBox(
		ErrorText,
		widget.NewButton(language.RestartButtonLabel, func() {
			Count = 1
			StartWindowContent()
		}),
		widget.NewButton(language.ExitButtonLabel, func() {
			Game.Quit()
		})))
}

func SetUpdatedContent() {
	StartWindow.SetContent(container.NewBorder(TextLabel, nil, nil, nil, GameGrid.Content()))
	Start = time.Now()
}

func GameContentUpdater() {
	number.ArrayShuffler(&Array)
	GameGrid.SetContent(container.NewGridWithColumns(5,
		widget.NewButton(strconv.Itoa(Array[0]), func() {
			TextLabel.SetText(strconv.Itoa(Array[0]))
			NumberChecker(Array[0])
		}),
		widget.NewButton(strconv.Itoa(Array[1]), func() {
			TextLabel.SetText(strconv.Itoa(Array[1]))
			NumberChecker(Array[1])
		}),
		widget.NewButton(strconv.Itoa(Array[2]), func() {
			TextLabel.SetText(strconv.Itoa(Array[2]))
			NumberChecker(Array[2])
		}),
		widget.NewButton(strconv.Itoa(Array[3]), func() {
			TextLabel.SetText(strconv.Itoa(Array[3]))
			NumberChecker(Array[3])
		}),
		widget.NewButton(strconv.Itoa(Array[4]), func() {
			TextLabel.SetText(strconv.Itoa(Array[4]))
			NumberChecker(Array[4])
		}),
		widget.NewButton(strconv.Itoa(Array[5]), func() {
			TextLabel.SetText(strconv.Itoa(Array[5]))
			NumberChecker(Array[5])
		}),
		widget.NewButton(strconv.Itoa(Array[6]), func() {
			TextLabel.SetText(strconv.Itoa(Array[6]))
			NumberChecker(Array[6])
		}),
		widget.NewButton(strconv.Itoa(Array[7]), func() {
			TextLabel.SetText(strconv.Itoa(Array[7]))
			NumberChecker(Array[7])
		}),
		widget.NewButton(strconv.Itoa(Array[8]), func() {
			TextLabel.SetText(strconv.Itoa(Array[8]))
			NumberChecker(Array[8])
		}),
		widget.NewButton(strconv.Itoa(Array[9]), func() {
			TextLabel.SetText(strconv.Itoa(Array[9]))
			NumberChecker(Array[9])
		}),
		widget.NewButton(strconv.Itoa(Array[10]), func() {
			TextLabel.SetText(strconv.Itoa(Array[10]))
			NumberChecker(Array[10])
		}),
		widget.NewButton(strconv.Itoa(Array[11]), func() {
			TextLabel.SetText(strconv.Itoa(Array[11]))
			NumberChecker(Array[11])
		}),
		widget.NewButton(strconv.Itoa(Array[12]), func() {
			TextLabel.SetText(strconv.Itoa(Array[12]))
			NumberChecker(Array[12])
		}),
		widget.NewButton(strconv.Itoa(Array[13]), func() {
			TextLabel.SetText(strconv.Itoa(Array[13]))
			NumberChecker(Array[13])
		}),
		widget.NewButton(strconv.Itoa(Array[14]), func() {
			TextLabel.SetText(strconv.Itoa(Array[14]))
			NumberChecker(Array[14])
		}),
		widget.NewButton(strconv.Itoa(Array[15]), func() {
			TextLabel.SetText(strconv.Itoa(Array[15]))
			NumberChecker(Array[15])
		}),
		widget.NewButton(strconv.Itoa(Array[16]), func() {
			TextLabel.SetText(strconv.Itoa(Array[16]))
			NumberChecker(Array[16])
		}),
		widget.NewButton(strconv.Itoa(Array[17]), func() {
			TextLabel.SetText(strconv.Itoa(Array[17]))
			NumberChecker(Array[17])
		}),
		widget.NewButton(strconv.Itoa(Array[18]), func() {
			TextLabel.SetText(strconv.Itoa(Array[18]))
			NumberChecker(Array[18])
		}),
		widget.NewButton(strconv.Itoa(Array[19]), func() {
			TextLabel.SetText(strconv.Itoa(Array[19]))
			NumberChecker(Array[19])
		}),
		widget.NewButton(strconv.Itoa(Array[20]), func() {
			TextLabel.SetText(strconv.Itoa(Array[20]))
			NumberChecker(Array[20])
		}),
		widget.NewButton(strconv.Itoa(Array[21]), func() {
			TextLabel.SetText(strconv.Itoa(Array[21]))
			NumberChecker(Array[21])
		}),
		widget.NewButton(strconv.Itoa(Array[22]), func() {
			TextLabel.SetText(strconv.Itoa(Array[22]))
			NumberChecker(Array[22])
		}),
		widget.NewButton(strconv.Itoa(Array[23]), func() {
			TextLabel.SetText(strconv.Itoa(Array[23]))
			NumberChecker(Array[23])
		}),
		widget.NewButton(strconv.Itoa(Array[24]), func() {
			TextLabel.SetText(strconv.Itoa(Array[24]))
			NumberChecker(Array[24])
		}),
	))
}
