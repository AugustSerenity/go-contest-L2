package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindAnagram(words *[]string) map[string][]string {
	m := make(map[string][]string)

	for _, val := range *words {
		lowerWord := strings.ToLower(val)
		sortWord := sortWords(lowerWord)

		m[sortWord] = append(m[sortWord], lowerWord)
	}

	for key, val := range m {
		if len(val) <= 1 {
			delete(m, key)
		} else {
			sort.Strings(val)
			m[key] = val
		}
	}

	return m
}

func sortWords(word string) string {
	runeSlice := []rune(word)

	sort.Slice(runeSlice, func(i, j int) bool {
		return runeSlice[i] < runeSlice[j]
	})

	return string(runeSlice)
}

func main() {
	str := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}

	s := FindAnagram(&str)

	for _, val := range s {
		fmt.Printf("%s: %v\n", val[0], val)
	}
}

//stop the world происходит На стадии подготовки перед маркировкой и на стадии завершения маркировки
//Во время самой маркировки, когла GC проходит по всему дереву объектов, исполнение кода (работы программы) не останавливается
//это значит, что нет зависимости от размера КУЧИ.
//ВЫЗОВ GC - ПРОИСХОДИТ, когда новая куча достигает 100% от живой кучи (эти 100% можно менять программно, тем самым увеличивая или
// уменьшая частоту вызова GC и размеры памяти кучи)
//GOGC - процент новой необработанной памяти кучи от живой памяти, при достижении которой будет запущена сборка мусора.
// GOGC = 100 (% по умолчанию)
// можно отключить вызов GC указава GOGC = -1
