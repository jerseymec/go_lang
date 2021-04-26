package main

import (
	"context"
	"fmt"
	"github.com/apache/beam/sdks/go/pkg/beam"
	_ "github.com/apache/beam/sdks/go/pkg/beam/io/filesystem/gcs"
	_ "github.com/apache/beam/sdks/go/pkg/beam/io/filesystem/local"
	_ "github.com/apache/beam/sdks/go/pkg/beam/io/filesystem/memfs"
	"github.com/apache/beam/sdks/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/go/pkg/beam/runners/direct"
	"github.com/apache/beam/sdks/go/pkg/beam/transforms/stats"
	"log"
	"regexp"
)

var wordRE = regexp.MustCompile(`[a-zA-Z]+('[a-z])?`)

func main() {
	beam.Init()
	p := beam.NewPipeline()
	s := p.Root()

	lines := textio.Read(s, "data.txt")

	words := beam.ParDo(s, func(line string, emit func(string)) {
		for _, word := range wordRE.FindAllString(line, -1) {
			emit(word)
		}
	}, lines)

	counted := stats.Count(s, words)

	formatted := beam.ParDo(s, func(w string, c int) string {
		return fmt.Sprintf("%s: %v", w, c)
	}, counted)

	textio.Write(s, "wordcount.txt", formatted)
	_, err := direct.Execute(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}

}
