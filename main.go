package main

import (
	"github.com/DnullP/GraphLang-GO/preprocess"
	"github.com/DnullP/GraphLang-GO/tasks"
)

func main() {
	textList := preprocess.TextRead("./data/xt1.txt", 500)

	for i, textSection := range textList {
		tasks.ReadNoval(textSection)
		if i%5 == 0 {
			tasks.TryMergePerson()
		}
	}
}
