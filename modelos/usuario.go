package modelos

type Persona struct {
	ID        int64  `json:"id" gorm:"primary_key;auto_increment"`
	Usuario    string `json:"usuario"`
	Password  string `json:"password"`
	Nombre string `json:"nombre"`
	Correo  string `json:"correo"`
}
