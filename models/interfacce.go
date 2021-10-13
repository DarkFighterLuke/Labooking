package models

type ReaderDB interface {
	Seleziona(cols ...string) error
}

type WriterDB interface {
	Aggiungi() (int64, error)
	Modifica() error
	Elimina(cols ...string) error
}

type ReadWriteDB interface {
	ReaderDB
	WriterDB
}
