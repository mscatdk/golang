package comment

//DAO simple in memory data store
type DAO struct {
	currentID int
	data      []Comment
}

func (d *DAO) UpdateComment(comment Comment) Comment {
	if comment.ID >= 0 && comment.ID < len(d.data) {
		d.data[comment.ID] = comment
	} else {
		panic("Update failed")
	}
	return comment
}

func (d *DAO) AddComment(comment Comment) Comment {
	comment.ID = d.currentID
	d.currentID = d.currentID + 1
	d.data = append(d.data, comment)
	return comment
}

func (d *DAO) GetComment(id int) Comment {
	return d.data[id]
}

func (d *DAO) List() []Comment {
	return d.data
}

func (d *DAO) Init() {
	d.AddComment(Comment{Author: "Demo Author1"})
	d.AddComment(Comment{Author: "Demo Author2"})
}
