package cmd

import (
	"GMG/grpc"
	lang "GMG/language"
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
)

func TimeChecker(Start time.Time) time.Duration {
	Duration := time.Since(Start)
	return Duration
}

func StartApp() {
	StartWindow.Resize(fyne.NewSize(233, 280))
	for i := 0; i <= 24; i++ {
		Array[i] = i + 1
	}
	Language()
	StartWindow.ShowAndRun()
}

func Language() {
	StartWindow.CenterOnScreen()
	StartWindow.SetContent(container.NewVBox(
		lang.LabelLanguage,
		widget.NewButton(lang.EnglishButton, func() {
			lang.UpdateLanguage()
			StartWindowContent()
		}),
		widget.NewButton(lang.RussianButton, func() {
			StartWindowContent()
		})))
}

func StartWindowContent() {
	GameContentUpdater()
	StartText := widget.NewLabel(lang.StartWindowText)
	StartWindow.CenterOnScreen()
	StartWindow.SetContent(container.NewVBox(
		StartText,
		widget.NewButton(lang.StartButtonLabel, func() {
			SetUpdatedContent()
		})))
}

func NumberChecker(number int) {
	if number == 25 && Count == 25 {
		EndTime = float32(TimeChecker(Start).Seconds())
		EndText.SetText(fmt.Sprintf(lang.ResultText, EndTime))
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
		widget.NewButton(lang.SendScoreToDb, func() {
			ScoreToDbContent()
		}),
		widget.NewButton(lang.RestartButtonLabel, func() {
			Count = 1
			StartWindowContent()
		}),
		widget.NewButton(lang.ExitButtonLabel, func() {
			Game.Quit()
		})))
}

func ErrorContentUpdater() {
	ErrorText := widget.NewLabel(lang.ErrorWindowText)
	StartWindow.SetContent(container.NewVBox(
		ErrorText,
		widget.NewButton(lang.RestartButtonLabel, func() {
			Count = 1
			StartWindowContent()
		}),
		widget.NewButton(lang.ExitButtonLabel, func() {
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

func ScoreToDbContent() {
	var a string
	input := widget.NewEntry()
	input.SetPlaceHolder("Введите имя...")
	StartWindow.SetContent(container.NewVBox(
		EndText,
		input,
		widget.NewButton(lang.SendScoreToDb, func() {
			grpc.ConnectGRPC(input.Text, EndTime)
			a = input.Text
			fmt.Println(a, EndTime)
			StartWindowContent()
		}),
		widget.NewButton(lang.RestartButtonLabel, func() {
			Count = 1
			StartWindowContent()
		}),
		widget.NewButton(lang.ExitButtonLabel, func() {
			Game.Quit()
		})))
}
