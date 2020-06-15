package tammyNal

import (
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

func render(w http.ResponseWriter, tmpl string,r *http.Request) {
	t, err := template.ParseFiles(tmpl);
	if err != nil {
		logrus.Errorln(err)
	}
	if err := t.Execute(w, nil); err != nil {
		logrus.Fatalln(err)
	}
}