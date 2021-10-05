package models

type ReaderDB interface {
	Seleziona(cols ...string) error
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
