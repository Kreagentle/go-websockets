package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

var setJet = jet.NewSet(jet.NewOSFileSystemLoader("./html"), jet.InDevelopmentMode())

func MainHandler(w http.ResponseWriter, r *http.Request) {
	err := pageRender(w, "home.jet", nil)
	if err != nil {
		log.Println("cant start the ui", err)
	}
}

func pageRender(w http.ResponseWriter, template string, data jet.VarMap) error {
	setupjet, err := setJet.GetTemplate(template)
	if err != nil {
		log.Println("cant setup jet library", err)
		return err
	}

	err = setupjet.Execute(w, data, nil)
	if err != nil {
		log.Println("cant execute the rendering", err)
		return err
	}

	return nil
}
