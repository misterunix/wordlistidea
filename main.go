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

var nouns []string
var adjectives []string
var verbs []string

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

	nounsRaw = ""
	adjectivesRaw = ""
	verbsRaw = ""

	// fmt.Println("Number of adjectives: ", len(adjectives))
	// fmt.Println("Number of nouns: ", len(nouns))
	// fmt.Println("Number of verbs: ", len(verbs))

	alen := len(adjectives)
	nlen := len(nouns)
	vlen := len(verbs)

	if !webserver {
		r := rand.Intn(alen)
		s := adjectives[r]
		r = rand.Intn(nlen)
		s += " " + nouns[r]
		r = rand.Intn(vlen)
		s += " " + verbs[r]
		fmt.Println(s)
		r = rand.Intn(alen)
		s = adjectives[r]
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

	fmt.Fprint(w, "<html><title>Random Adjective Noun Verb</title><body>")

	fmt.Fprintln(w, "<p style=\"font-size:26px;\" >")

	alen := len(adjectives)
	nlen := len(nouns)
	vlen := len(verbs)

	rn := rand.Intn(alen)
	s := adjectives[rn]

	rn = rand.Intn(nlen)
	s += " " + nouns[rn]

	rn = rand.Intn(vlen)
	s += " " + verbs[rn] + "</p>"

	fmt.Fprintln(w, s)

	fmt.Fprintln(w, "<p style=\"font-size:26px;\" >")

	rn = rand.Intn(alen)
	s = adjectives[rn]

	rn = rand.Intn(nlen)
	s += " " + nouns[rn]

	rn = rand.Intn(vlen)
	s += " " + verbs[rn]

	fmt.Fprintln(w, s, "</p>")
	//fmt.Fprintln(w, s, "<br>")

	fmt.Fprintln(w, "</body></html>")

}
