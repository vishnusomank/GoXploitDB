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

	services.Git_Operation(utils.GIT_DIR)

	iterate(utils.GIT_DIR + "/exploits")

	s := gocron.NewScheduler(time.UTC)

	s.Every("120m").Do(func() {

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

	pathSlice := strings.Split(path, "/")

	typeVal := strings.ToUpper(pathSlice[len(pathSlice)-2])
	platformVAl := strings.ToUpper(pathSlice[len(pathSlice)-3])

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
				titleValue = strings.Trim(titleValue, "#[]")
				titleValue = strings.TrimSpace(titleValue)
			}

			for scanner.Scan() {
				if strings.Contains(string(scanner.Text()), "Exploit Author") {
					slice2 := strings.Split(scanner.Text(), ":")

					if len(slice2) > 1 {
						authorValue = slice2[1]
						authorValue = strings.Trim(authorValue, "#[]")
						authorValue = strings.TrimSpace(authorValue)
					}

					for scanner.Scan() {
						if strings.Contains(string(scanner.Text()), "CVE :") || strings.Contains(string(scanner.Text()), "CVE:") {
							slice3 := strings.Split(scanner.Text(), ":")
							if len(slice3) > 1 {
								cveValue = slice3[1]
								cveValue = strings.Trim(cveValue, "#[]")
								cveValue = strings.TrimSpace(cveValue)

								if !strings.Contains(cveValue, "CVE") {
									cveValue = "CVE-" + cveValue

								}
								if strings.Contains(cveValue, "CVE-N/A") || strings.Contains(cveValue, "CVE-NA") || strings.Contains(cveValue, "CVE-n/a") || strings.Contains(cveValue, "CVE-na") {
									cveValue = "N/A"
								}
							}

							break
						}
					}
				}
			}
		}
	}

	if cveValue != "" {
		fmt.Println(cveValue)

		xploitdb := models.XploitDB{Title: titleValue, URL: "https://www.exploit-db.com/exploits/" + edb_id, CVE: strings.ToUpper(cveValue), Author: authorValue, Type: typeVal, Platform: platformVAl, EDBID: edb_id}
		models.DB.Create(&xploitdb)
	}

}
