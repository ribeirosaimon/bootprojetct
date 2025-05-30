package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func (m *mysqlConnection) RunMigrations(ctx context.Context) error {
	log.Printf("Running migrations")
	out, err := exec.Command("go", "env", "GOMOD").Output()
	if err != nil {
		panic(err)
	}
	gomodDir := filepath.Dir(strings.TrimSpace(string(out)))
	fileRoot := fmt.Sprintf("%s/static/migrations", gomodDir)
	entries, err := os.ReadDir(fileRoot)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		log.Printf("Running migration %s", e.Name())
		file, err := os.ReadFile(fmt.Sprintf("%s/%s", fileRoot, e.Name()))
		if err != nil {
			log.Fatalf("Erro read SQL: %v", err)
		}
		_, err = m.Connection.Exec(string(file))
		if err != nil {
			log.Fatalf("Erro exec SQL: %v", err)
			return err
		}
		log.Printf("SQL: %s", e.Name())
	}
	return nil
}
