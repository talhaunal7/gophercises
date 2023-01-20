package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"log"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("my.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("TaskBucket"))
			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Printf("%s- %s\n", k, v)
			}
			return nil
		})

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
