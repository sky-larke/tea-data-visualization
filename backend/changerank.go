package main
import (
    "context"
    "log"
    "os"
	"io"
	"strconv"
	"encoding/csv"
    //"github.com/google/uuid"
    "github.com/jackc/pgx/v5"
)

type Tea struct {
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
	// open the csv
	file, err := os.Open("backend/tea.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	// SELECT the rows
	rowptr, err := conn.Query(context.Background(), "SELECT name, rank, year, vendor, cost, type, subtype, cultivar, amount FROM Teas ORDER BY rank ASC NULLS LAST")
	if err != nil {
        log.Fatal(err)
    }
	defer rowptr.Close()

	teaMap := make(map[string]Tea)
	for rowptr.Next() {
		var t Tea
		err := rowptr.Scan(&t.Name, &t.Rank, &t.Year, &t.Vendor, &t.Cost, &t.Type, &t.Subtype, &t.Cultivar, &t.Amount)
		if err != nil {
			log.Fatal("DB read error:", err)
		}
		key := t.Vendor + "|" + t.Name
		teaMap[key] = t
	}
	var localTea Tea

	csvptr := 1
	rank := 1
	update := false
	for {
		row, err := reader.Read() // get row

		// checks for the csv's validity
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Skipping line %d: %v", row, err)
			csvptr++
			continue
		} else if len(row) < 8 {
			log.Printf("Skipping line %d: not enough fields (%d)", row, len(row))
			csvptr++
			continue
		}
		
		// get tea of the csv row 
		year, _ := strconv.Atoi(row[0])     
		cost, _ := strconv.ParseFloat(row[6], 64)  
		amount, _ := strconv.ParseFloat(row[7], 64) 

		localTea = Tea{
			Rank: rank,
			Year:     year,
			Vendor:   row[1],
			Name:     row[2],
			Type:     row[3],
			Subtype:  row[4],
			Cultivar: row[5],
			Cost:     cost,
			Amount:   amount,
		}

		// compare local to db version
		key := localTea.Vendor + "|" + localTea.Name
		dbTea, exists := teaMap[key]

		
		if exists {
			update = false
			if localTea.Year != dbTea.Year {
				dbTea.Year = localTea.Year
				update = true
			}
			if localTea.Type != dbTea.Type {
				dbTea.Type = localTea.Type
				update = true
			}
			if localTea.Subtype != dbTea.Subtype {
				dbTea.Subtype = localTea.Subtype
				update = true
			}
			if localTea.Cultivar != dbTea.Cultivar {
				dbTea.Cultivar = localTea.Cultivar
				update = true
			}
			if localTea.Cost != dbTea.Cost {
				dbTea.Cost = localTea.Cost
				update = true
			}
			if localTea.Amount != dbTea.Amount {
				dbTea.Amount = localTea.Amount
				update = true
			}

			if rank != dbTea.Rank {
				dbTea.Rank = rank
				update = true
			}

			if update {
				_, err := conn.Exec(context.Background(), `
				UPDATE Teas
				SET year = $1,
					rank = $2,
					vendor = $3,
					name = $4,
					type = $5,
					subtype = $6,
					cultivar = $7,
					cost = $8,
					amount = $9
				WHERE name = $4 AND vendor = $3`, 
				dbTea.Year, dbTea.Rank, dbTea.Vendor, dbTea.Name, dbTea.Type, dbTea.Subtype, dbTea.Cultivar, dbTea.Cost, dbTea.Amount)
				if err != nil {
					log.Printf("Error UPDATEing %s/%s: %v", dbTea.Vendor, dbTea.Name, err)
				}
			}
			
			log.Println("Updating: ", dbTea.Vendor, dbTea.Name)
			
			// move rowptr
			if rowptr.Next() {
				err := rowptr.Scan(
					&dbTea.Name,
					&dbTea.Rank,
					&dbTea.Year,
					&dbTea.Vendor,
					&dbTea.Cost,
					&dbTea.Type,
					&dbTea.Subtype,
					&dbTea.Cultivar,
					&dbTea.Amount,
				)
				if err != nil {
					log.Fatal("166", err)
				}
			}
		} else {
			log.Println("Adding: ", localTea.Vendor, localTea.Name)
			_, err := conn.Exec(context.Background(), `
				INSERT INTO Teas (
					year,
					rank,
					vendor,
					name,
					type,
					subtype,
					cultivar,
					cost,
					amount
				) VALUES (
					$1, $2, $3, $4, $5, $6, $7, $8, $9
				)`,
				localTea.Year, localTea.Rank, localTea.Vendor, localTea.Name, localTea.Type, localTea.Subtype, localTea.Cultivar, localTea.Cost, localTea.Amount)
			if err != nil {
				log.Printf("Error INSERTing %s/%s: %v", dbTea.Vendor, dbTea.Name, err)
			}
		}

		rank++
		csvptr++
	}


    return nil
}

func printTeas(conn *pgx.Conn) error {
    rows, err := conn.Query(context.Background(), "SELECT name, rank, year, vendor, cost FROM Teas ORDER BY rank ASC NULLS LAST")
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
    log.Println("\n\nAll tea:")
    printTeas(conn)

}