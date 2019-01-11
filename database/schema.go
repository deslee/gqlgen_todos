package todos_db

import "database/sql"

func CreateTablesIfNotExist(db *sql.DB) {
	var err error

	_,err = db.Exec(`
		CREATE TABLE IF NOT EXISTS "Users" (
			"ID" TEXT NOT NULL CONSTRAINT "PK_Users" PRIMARY KEY,
			"Name" TEXT NULL
		)
	`)

	if err != nil {
		panic(err)
	}

	_,err = db.Exec(`
		CREATE TABLE IF NOT EXISTS "Todos" (
			"ID" TEXT NOT NULL CONSTRAINT "PK_Todos" PRIMARY KEY,
			"Text" TEXT NULL,
		  	"UserID" TEXT NOT NULL,
		  	CONSTRAINT "FK_Todos_Users_UserID" FOREIGN KEY ("UserID") REFERENCES "Users" ("ID") ON DELETE CASCADE 
		)
	`)

	if err != nil {
		panic(err)
	}
}
