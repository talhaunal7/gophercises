package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func getNewKeyValue(b *bolt.Bucket) []byte {
	c := b.Cursor()
	k, _ := c.Last()
	keyIntValue, _ := strconv.Atoi(string(k))
	keyIntValue += 1
	k = []byte(strconv.Itoa(keyIntValue))
	return k
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Args:  cobra.MinimumNArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		db, _ := bolt.Open("my.db", 0600, nil)
		defer db.Close()

		err := db.Update(func(tx *bolt.Tx) error {
			b, _ := tx.CreateBucketIfNotExists([]byte("TaskBucket"))
			task := strings.Join(args, " ")
			err := b.Put(getNewKeyValue(b), []byte(task))
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
