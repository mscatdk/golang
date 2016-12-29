package comment

import "time"

//Comment represent one comment
type Comment struct {
	ID      int       `json:"id"`
	Author  string    `json:"author"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}

//Comments is short hand for an array of comments
type Comments []Comment
