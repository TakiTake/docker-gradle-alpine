package main

import (
	"html/template"
	"os"
)

const (
	JAVA8   = "8"
	JAVA9   = "9"
	JAVA10  = "10"
	GRADLE2 = "2.14.1"
	GRADLE3 = "3.5"
	GRADLE4 = "4.10.2"
)

var tmpl = template.Must(template.ParseFiles("Dockerfile.tmpl"))

type Version struct {
	JAVA_VERSION   string
	GRADLE_VERSION string
}

func generate(k string, v Version) {
	w, err := os.Create(k)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	err = tmpl.Execute(w, v)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
}

func main() {
	versions := map[string]Version{
		"openjdk8/gradle2/Dockerfile":  Version{JAVA8, GRADLE2},
		"openjdk8/gradle3/Dockerfile":  Version{JAVA8, GRADLE3},
		"openjdk8/gradle4/Dockerfile":  Version{JAVA8, GRADLE4},
		"openjdk9/gradle2/Dockerfile":  Version{JAVA9, GRADLE2},
		"openjdk9/gradle3/Dockerfile":  Version{JAVA9, GRADLE3},
		"openjdk9/gradle4/Dockerfile":  Version{JAVA9, GRADLE4},
		"openjdk10/gradle2/Dockerfile": Version{JAVA10, GRADLE2},
		"openjdk10/gradle3/Dockerfile": Version{JAVA10, GRADLE3},
		"openjdk10/gradle4/Dockerfile": Version{JAVA10, GRADLE4},
	}

	for k, v := range versions {
		generate(k, v)
	}
}
