package main

import (
	"reflect"
	"testing"
)

func TestTransformation(t *testing.T) {
	tests := []struct {
		name      string
		inputData []string
		want      []string
		column    int
		numeric   bool
		revers    bool
		unique    bool
	}{
		{
			name: "Обратная сортировка text1 (по имени)",
			inputData: []string{
				"Иван\t35",
				"Пётр\t28",
				"Олег\t42",
			},
			want: []string{
				"Пётр\t28",
				"Олег\t42",
				"Иван\t35",
			},
			column: 0,
			revers: true,
		},
		{
			name: "Обратная сортировка text2 (по имени)",
			inputData: []string{
				"Иван\t35",
				"Пётр\t28",
				"Олег\t42",
				"Павел\t42",
			},
			want: []string{
				"Пётр\t28",
				"Павел\t42",
				"Олег\t42",
				"Иван\t35",
			},
			column: 0,
			revers: true,
		},
		{
			name: "Обратная сортировка text3 с уникальными строками (по имени)",
			inputData: []string{
				"Иван\t35",
				"Иван\t35",
				"Иван\t35",
				"Пётр\t28",
				"Олег\t42",
				"Олег\t42",
				"Олег\t42",
				"Павел\t42",
				"Павел\t42",
			},
			want: []string{
				"Пётр\t28",
				"Павел\t42",
				"Олег\t42",
				"Иван\t35",
			},
			column: 0,
			revers: true,
			unique: true,
		},
		{
			name: "Нормальная сортировка text1 по возрасту (числовая)",
			inputData: []string{
				"Иван\t35",
				"Пётр\t28",
				"Олег\t42",
			},
			want: []string{
				"Пётр\t28",
				"Иван\t35",
				"Олег\t42",
			},
			column:  1,
			numeric: true,
			revers:  false,
		},
		{
			name: "Обратная числовая сортировка text1 по возрасту",
			inputData: []string{
				"Иван\t35",
				"Пётр\t28",
				"Олег\t42",
			},
			want: []string{
				"Олег\t42",
				"Иван\t35",
				"Пётр\t28",
			},
			column:  1,
			numeric: true,
			revers:  true,
		},
		// --- Новые тесты ---
		{
			name: "Уникальные строки с обычной сортировкой (по имени)",
			inputData: []string{
				"Иван\t35",
				"Иван\t35",
				"Пётр\t28",
				"Пётр\t28",
				"Олег\t42",
			},
			want: []string{
				"Иван\t35",
				"Олег\t42",
				"Пётр\t28",
			},
			column: 0,
			unique: true,
		},
		{
			name: "Обратная сортировка с уникальными строками",
			inputData: []string{
				"Иван\t35",
				"Иван\t35",
				"Пётр\t28",
				"Пётр\t28",
				"Олег\t42",
				"Олег\t42",
			},
			want: []string{
				"Пётр\t28",
				"Олег\t42",
				"Иван\t35",
			},
			column: 0,
			revers: true,
			unique: true,
		},
		{
			name: "Уникальные строки с числовой сортировкой по возрасту",
			inputData: []string{
				"Иван\t35",
				"Иван\t35",
				"Пётр\t28",
				"Пётр\t28",
				"Олег\t42",
				"Олег\t42",
			},
			want: []string{
				"Пётр\t28",
				"Иван\t35",
				"Олег\t42",
			},
			column:  1,
			numeric: true,
			unique:  true,
		},
	}

	for _, testCase := range tests {
		got := transformation(testCase.inputData, testCase.column, testCase.numeric, testCase.revers, testCase.unique)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("FAIL: %s\nОжидалось: %v\nПолучено: %v", testCase.name, testCase.want, got)
		}
	}
}
