package cmd

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
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
			redisClient := database.ConnectWithRedis()

			err := parseAndSave(filename, redisClient)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Data successfully saved to MongoDB.")
		},
	}

	rootCmd.AddCommand(parseAndSaveCmd)
}

func parseAndSave(filename string, redisClient *redis.Client) error {
	var ctx = context.Background()

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

	for i := 1; ; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("CSV reading error: %s", err)
		}

		key := fmt.Sprintf("fixture_%d", i)
		val, err := redisClient.Get(ctx, key).Bytes()
		if err == nil {
			var fixture model.Fixture
			err := json.Unmarshal(val, &fixture)
			if err != nil {
				return fmt.Errorf("error unmarshaling fixture from JSON: %s", err)
			}
			fmt.Println("Data found in cache:", fixture)
			continue
		} else if err != redis.Nil {
			return fmt.Errorf("redis error: %s", err)
		}

		fixture := model.Fixture{
			Date:     record[1],
			HomeTeam: record[2],
			AwayTeam: record[3],
			FTHG:     parseInt(record[4]),
			FTAG:     parseInt(record[5]),
			Referee:  record[10],
		}

		fixtureBytes, err := json.Marshal(fixture)
		if err != nil {
			return fmt.Errorf("error marshaling fixture to JSON: %s", err)
		}

		err = redisClient.Set(ctx, key, fixtureBytes, time.Hour).Err()
		if err != nil {
			return fmt.Errorf("redis error: %s", err)
		}
		fmt.Println("Data successfully saved to redis.")

		// Insert into MongoDB
		_, err = collection.InsertOne(context.Background(), fixture)
		if err != nil {
			return fmt.Errorf("inserting into MongoDB error: %s", err)
		}
	}

	return nil
}

func parseInt(numStr string) int {
	var num int
	fmt.Sscanf(numStr, "%d", &num)
	return num
}
