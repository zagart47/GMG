package language

import (
	"fyne.io/fyne/v2/widget"
)

var (
	ExitButtonLabel = "Выход"
	StartWindowText = "Будет показано игровое поле\n" +
		"размером 5х5. Необходимо\n" +
		"находить и кликать на числа\n" +
		"от 1 до 25 в порядке их\n" +
		"возрастания. Игра начнется\n" +
		"после нажатия кнопки \"Старт\"."
	StartButtonLabel   = "Старт"
	RestartButtonLabel = "Перезапуск"
	ResultText         = "Ваш результат:\n %.3f секунд(ы)."
	ErrorWindowText    = "Вы ошиблись.\n" +
		"Вы можете начать игру\n" +
		"заново нажав кнопку\n" +
		"\"Перезапуск\"."
	EnglishButton          = "ENGLISH"
	RussianButton          = "RUSSIAN"
	LabelLanguage          = widget.NewLabel("Choose your language:")
	SendScoreToDb          = "Оправить очки в базу"
	WhatMyPlaceButtonLabel = "Какое место у меня?"
	ShowUserTopButtonLabel = "Показать Топ-10 игроков"
)

func UpdateLanguage() {
	ExitButtonLabel = "Exit"
	StartWindowText =
		"A 5x5 playing field will now\n" +
			"be shown. You need to\n" +
			"find and click on numbers\n" +
			"from 1 to 25 in ascending\n" +
			"order. The game will start\n" +
			"after pressing the \"Start\"."
	StartButtonLabel = "Start"
	RestartButtonLabel = "Restart"
	ResultText = "Your result:\n %.3f second(s)."
	ErrorWindowText = "You made a mistake.\n" +
		"You can start the game again\n" +
		"by pressing the \"Restart\"\n" +
		"button."

}
