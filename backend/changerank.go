package main
import (
    "context"
    "log"
    "os"
	"fmt"
	"strconv"
	"encoding/csv"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5"
)

type Tea struct {
	Id          int     `json:"id"`
	Year		int		`json:"year"`
	Rank		int		`json:"rank"`
	Vendor      string  `json:"vendor"`
	Name        string  `json:"name"`
	Type    	string  `json:"type"`
	Subtype 	string  `json:"subtype"`
	Cultivar 	string  `json:"cultivar"`
	Cost       	float64 `json:"cost"`
	Amount		float64	`json:"amount"`
	// Score		float32 `json:"score"`
}


func insertFromCSV(conn *pgx.Conn) error {
	var records []Tea // This would be your CSV records
	var dbRecords []Tea
	// open the csv

	file, err := os.Open("backend/tea.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rawRecords, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	// convert csv to correct format
	for _, row := range rawRecords {
		if len(row) < 9 {
			log.Println("Skipping invalid row:", row)
			continue
		}
		
		
		id, _  	:= strconv.Atoi(row[0])
		year, _ := strconv.Atoi(row[1])    
		rank, _ := strconv.Atoi(row[2])    
		cost, _ := strconv.ParseFloat(row[8], 64)  

		var teaType string 
		var subtype string
		if row[5] != "" {
			teaType = row[5]
		}

		if row[6] != "" {
			subtype = row[6]
		} 

		amount, _ := strconv.ParseFloat(row[9], 64) 

		tea := Tea{
			Id: id,
			Year:    year,
			Rank:    rank,
			Name:    row[3],
			Vendor:  row[4],
			Type:    teaType,
			Subtype: subtype,
			Cultivar: row[7],
			Cost:    cost,
			Amount:  amount,
		}

		records = append(records, tea)
		log.Printf(tea.Name)
	}

	// SELECT the rows
	rows, err := conn.Query(context.Background(), "SELECT name, rank, year, vendor, cost, type, subtype, amount FROM teas ORDER BY rank ASC NULLS LAST")
	if err != nil {
        log.Fatal(err)
    }
	defer rows.Close()

	// convert and add to records
	for rows.Next() {
		var r Tea
		if err := rows.Scan(&r.Name, &r.Rank, &r.Year, &r.Vendor, &r.Cost, &r.Type, &r.Subtype, &r.Amount); err != nil {
			log.Fatal(err)
		}
		dbRecords = append(dbRecords, r)
	}

	// sqlIndex := 0
	// sqlLen := len(dbRecords)


    return nil
}


func insertRows(ctx context.Context, tx pgx.Tx, accts [4]uuid.UUID) error {
    // Insert four rows into the "accounts" table.
    log.Println("Creating new entry...")
    if _, err := tx.Exec(ctx,
        "INSERT INTO teas (id, balance) VALUES ($1, $2)", accts[0], 250); err != nil {
        return err
    }
    return nil
}

func printTeas(conn *pgx.Conn) error {
    rows, err := conn.Query(context.Background(), "SELECT name, rank, year, vendor, cost FROM teas ORDER BY rank ASC NULLS LAST")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
		var year int
		var vendor string
        var name string
        var rank int
		var price float64
        if err := rows.Scan(&name, &rank, &year, &vendor, &price); err != nil {
            log.Fatal(err)
        }
		if year == 0 {
			log.Printf("%d. %s %s ($%.2f)\n",rank, vendor, name, price)
		} else {
			log.Printf("%d. %d %s %s ($%.2f)\n", rank, year, vendor, name, price)
		}
       
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
	log.Println("Adding from CSV:")
	insertFromCSV(conn)
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