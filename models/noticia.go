package models

type Noticia struct {
	ID        int64  `json:"id" gorm:"primary_key;auto_increment"`
	Titulo    string `json:"titulo"`
	Descripcion  string `json:"descripcion"`
	Autor string `json:"autor"`
	Publicado  string `json:"publicado"`
	Fuente  string `json:"fuente"`
	Categoria  string `json:"categoria"`
}