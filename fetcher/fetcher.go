package fetcher

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/go-co-op/gocron"
	"github.com/vishnusomank/GoXploitDB/models"
	"github.com/vishnusomank/GoXploitDB/services"
	"github.com/vishnusomank/GoXploitDB/utils"
)

func StartGit() {
	var err error
	utils.CURRENT_DIR, err = os.Getwd()
	if err != nil {
		fmt.Printf("[%s][%s] Failed to get current directory: %v\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.RedString("ERR"), err)

	}
	utils.GIT_DIR = utils.CURRENT_DIR + "/exploitdb"

	s := gocron.NewScheduler(time.UTC)

	s.Every("2m").Do(func() {

		services.Git_Operation(utils.GIT_DIR)

		iterate(utils.GIT_DIR + "/exploits")

	})
	s.StartAsync()

}

func iterate(path string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}

		if !info.IsDir() {
			split := strings.Split(info.Name(), ".")
			if len(split) > 0 {
				XploitDBCreate(split[0], path)
			}
		}
		return nil
	})
}

func XploitDBCreate(edb_id, path string) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var cveValue = ""
	var authorValue = ""
	var titleValue = ""
	for scanner.Scan() {
		if strings.Contains(string(scanner.Text()), "Exploit Title") {
			slice1 := strings.Split(scanner.Text(), ":")

			if len(slice1) > 1 {
				titleValue = slice1[1]
			}

			for scanner.Scan() {
				if strings.Contains(string(scanner.Text()), "Exploit Author") {
					slice2 := strings.Split(scanner.Text(), ":")

					if len(slice2) > 1 {
						authorValue = slice2[1]
					}

					for scanner.Scan() {
						if strings.Contains(string(scanner.Text()), "CVE :") || strings.Contains(string(scanner.Text()), "CVE:") {
							slice3 := strings.Split(scanner.Text(), ":")
							if len(slice3) > 1 {
								cveValue = slice3[1]
							}

							break
						}
					}
				}
			}
		}
	}
	if cveValue != "" {

		xploitdb := models.XploitDB{Title: titleValue, URL: "https://www.exploit-db.com/exploits/" + edb_id, CVE: strings.ToUpper(cveValue), Author: authorValue}
		models.DB.Create(&xploitdb)
	}

}
