package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

//go:embed assets/adjectives.txt
var adjectivesRaw string

//go:embed assets/nouns.txt
var nounsRaw string

//go:embed assets/verbs.txt
var verbsRaw string

//go:embed assets/jobs.txt
var jobsRaw string

var nouns []string
var adjectives []string
var verbs []string
var jobs []string

func main() {

	webserver := false

	flag.BoolVar(&webserver, "w", false, "Run as a webserver")
	flag.Parse()

	/*
		f, err := os.Open("assets/adjectives.txt")
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			adjectives = append(adjectives, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		f.Close()
	*/

	nouns = strings.Split(nounsRaw, "\n")
	adjectives = strings.Split(adjectivesRaw, "\n")
	verbs = strings.Split(verbsRaw, "\n")
	jobs = strings.Split(jobsRaw, "\n")

	nounsRaw = ""
	adjectivesRaw = ""
	verbsRaw = ""
	jobsRaw = ""

	// fmt.Println("Number of adjectives: ", len(adjectives))
	// fmt.Println("Number of nouns: ", len(nouns))
	// fmt.Println("Number of verbs: ", len(verbs))

	alen := len(adjectives)
	nlen := len(nouns)
	vlen := len(verbs)
	jlen := len(jobs)

	if !webserver {
		r := rand.Intn(jlen)
		j := strings.TrimSpace(jobs[r])
		s := "\"" + j + "\""

		r = rand.Intn(alen)
		s += " " + adjectives[r]

		r = rand.Intn(nlen)
		s += " " + nouns[r]

		r = rand.Intn(vlen)
		s += " " + verbs[r]

		fmt.Println(s)

		r = rand.Intn(jlen)
		j = strings.TrimSpace(jobs[r])
		s = "\"" + j + "\""

		r = rand.Intn(alen)
		s += " " + adjectives[r]

		r = rand.Intn(nlen)
		s += " " + nouns[r]

		r = rand.Intn(vlen)
		s += " " + verbs[r]

		fmt.Println(s)
		os.Exit(0)
	}

	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func getRoot(w http.ResponseWriter, r *http.Request) {

	top := "<html><head><title>Word List Ideas</title></head>"
	top += "<body><div class=\"container\"><div class=\"header\"><h1>Word List for writing prompts.</h1>"
	top += "<p><h2>Use these words to help spark your creative side.</h2><p></div><div class=\"content\">"
	fmt.Fprintln(w, top)

	alen := len(adjectives)
	nlen := len(nouns)
	vlen := len(verbs)

	for i := 0; i < 2; i++ {
		fmt.Fprintln(w, "<p style=\"font-size:26px;\" >")

		rn := rand.Intn(alen)
		s := adjectives[rn]

		rn = rand.Intn(nlen)
		s += " " + nouns[rn]

		rn = rand.Intn(vlen)
		s += " " + verbs[rn] + "</p>"

		fmt.Fprintln(w, s)
	}

	bottom := "</div><div class=\"footer\"><p><a href=\"index.html\">Reload</a></p></div></div></body></html>"
	fmt.Fprintln(w, bottom)

	// fmt.Fprintln(w, "<p style=\"font-size:26px;\" >")

	// rn = rand.Intn(alen)
	// s = adjectives[rn]

	// rn = rand.Intn(nlen)
	// s += " " + nouns[rn]

	// rn = rand.Intn(vlen)
	// s += " " + verbs[rn]

	// fmt.Fprintln(w, s, "</p>")
	// //fmt.Fprintln(w, s, "<br>")

	// fmt.Fprintln(w, "</body></html>")

}
