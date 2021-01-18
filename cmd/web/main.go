package main
import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_"github.com/jackc/pgx"
)
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}
func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "Psql data source name")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil

	//connStr := "user=postgres password=mypass dbname=productdb sslmode=disable"
	//db, err := sql.Open("postgres", connStr)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//
	//result, err := db.Exec("insert into Products (model, company, price) values ('iPhone X', $1, $2)",
	//	"Apple", 72000)
	//if err != nil{
	//	panic(err)
	//}

}