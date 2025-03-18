package main

import (
	"errors"
	"net/http"
	"strconv"

	"snippetbox.gteruithi.com/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) getSnippetView(w http.ResponseWriter, r *http.Request) {
	snippetId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || snippetId < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(snippetId)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.html", data)
}

func (app *application) getSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet!"))
}

func (app *application) postSnippetCreate(w http.ResponseWriter, r *http.Request) {

	title := "O snail"
	content := "O snail \n Climb Mount Fuji, \n But slowly, slowly! \n \n- Kobayashi Issa"

	err := app.snippets.Insert(title, content)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet!"))
}
