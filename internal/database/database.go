package database

type Idatabase interface {
}

type database struct {
	// db connection
}

func New( /*db conn*/ ) *database {
	return &database{
		//db:db
	}
}

// func Connect...
