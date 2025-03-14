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


func insertRows(ctx context.Context, tx pgx.Tx, accts [4]uuid.UUID) error {
    // Insert four rows into the "accounts" table.
    log.Println("Creating new rows...")
    if _, err := tx.Exec(ctx,
        "INSERT INTO teas (id, balance) VALUES ($1, $2), ($3, $4), ($5, $6), ($7, $8)", accts[0], 250, accts[1], 100, accts[2], 500, accts[3], 300); err != nil {
        return err
    }
    return nil
}

func printTeas(conn *pgx.Conn) error {
    rows, err := conn.Query(context.Background(), "SELECT name, rank FROM teas ORDER BY rank ASC NULLS LAST")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        var name string
        var rank int
        if err := rows.Scan(&name, &rank); err != nil {
            log.Fatal(err)
        }
        log.Printf("%s: %d\n", name, rank)
    }
    return nil
}

// func add(conn *pgx.Conn, after int, entry string) error {
// 	_, err := tx.Exec(ctx,
//         "INSERT INTO teas (id, rank, name, year, type, subtype, cost, amount, vendor) VALUES $1", entry); 
//     if err != nil {
// 		log.Fatal(err)
// 	}

// 	rows, err := conn.Query(context.Background(), "UPDATE teas SET rank = rank + 1 WHERE id >= $1",after) 
// 	if err != nil {
// 		log.Fatal(err)
// 	}
    
	
// 	return nil
// }


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
    printTeas(conn)

	var rank int = 121
    var name string = ""    
    var year int = 0
    var tea_type string = ""
    var subtype string = ""
    var cost int = 0
    var amount int = 0
    var vendor string = ""
	
	var command string = fmt.Sprintf("(%d, %s, %d, %s, %s, %d, %d, %s)", rank, name, year, tea_type, subtype, cost, amount, vendor)
	log.Printf(command)
	//add(conn, command)



}