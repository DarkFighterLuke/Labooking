package models

type ReaderDB interface {
	Seleziona() (interface{}, error)
}

type WriterDB interface {
	Aggiungi() error
	Modifica() error
	Elimina() error
}

type ReadWriteDB interface {
	ReaderDB
	WriterDB
}
