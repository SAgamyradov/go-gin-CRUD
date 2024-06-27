package model

type Album struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Actor string  `json:"actor"`
	Price float64 `json:"price"`
}

var Albums = []Album{
	{ID: "1", Title: "Film", Actor: "Johny Dep", Price: 34.23},
	{ID: "2", Title: "Jeru", Actor: "Gerry Mulligan", Price: 17.99},
}
