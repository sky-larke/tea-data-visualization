package main
import (
    "context"
    "log"
    "os"
	"fmt"

    // "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5"
)

func main() {
    // Read in connection string
    config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal(err)
    }
    config.RuntimeParams["application_name"] = "$ docs_simplecrud_gopgx"
    conn, err := pgx.ConnectConfig(context.Background(), config)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close(context.Background())


    // Print out the balances
    log.Println("All tea:")
	
	var command string = fmt.Sprintf("(%d, %s, %d, %s, %s, %d, %d, %s)", rank, name, year, tea_type, subtype, cost, amount, vendor)
	log.Printf(command)
	//add(conn, command)



}