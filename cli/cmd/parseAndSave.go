package cmd

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/taylankalkan01/onedio-case/cli/database"
	"github.com/taylankalkan01/onedio-case/cli/model"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "football-fixtures-cli",
	Short: "A CLI tool for football fixtures",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	var parseAndSaveCmd = &cobra.Command{
		Use:   "parseAndSave [filename]",
		Short: "Parse and save fixtures from CSV file to MongoDB",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]

			err := parseAndSave(filename)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Data successfully saved to MongoDB.")
		},
	}

	rootCmd.AddCommand(parseAndSaveCmd)
}

func parseAndSave(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("CSV reading error: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("CSV reading error: %s", err)
	}

	client := database.ConnectWithMongodb()
	defer client.Disconnect(context.Background())

	collection := client.Database("football").Collection("fixtures")

	var fixtures []interface{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("CSV reading error: %s", err)
		}

		fixture := model.Fixture{
			Date:     record[1],
			HomeTeam: record[2],
			AwayTeam: record[3],
			FTHG:     parseInt(record[4]),
			FTAG:     parseInt(record[5]),
			Referee:  record[10],
		}

		fixtures = append(fixtures, fixture)
	}

	// Insert into MongoDB
	_, err = collection.InsertMany(context.Background(), fixtures)
	if err != nil {
		return fmt.Errorf("inserting into MongoDB error: %s", err)
	}

	return nil
}

func parseInt(numStr string) int {
	var num int
	fmt.Sscanf(numStr, "%d", &num)
	return num
}
