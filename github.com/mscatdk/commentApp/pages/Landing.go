package commentPages

import (
	"html/template"
	"net/http"
	"os"

	"github.com/mscatdk/commentApp/comment"
)

//Landing - Landing page model
type Landing struct {
	Title    string
	Comments comment.Comments
}

//LandingPage - Return the landing page
func LandingPage(w http.ResponseWriter, r *http.Request) {
	pwd, _ := os.Getwd()
	//TODO: Template path is hardcoded
	t, err := template.ParseFiles(pwd + "/github.com/mscatdk/commentApp/tmpl/landing.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, Landing{Title: "Welcome to the Comment App", Comments: comment.CommentDAO.List()})
}
