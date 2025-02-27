package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// "suggest" command.
var suggestCmd = &cobra.Command{
	Use:   "suggest [query]",
	Short: "Suggest a Bash command from a natural language query",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args, " ")
		if !strings.HasPrefix(strings.ToLower(query), "translate english to bash:") {
			query = "translate English to Bash: " + query
		}

		command, err := queryNLP(query)
		if err != nil {
			fmt.Printf("Error querying NLP backend: %v\n", err)
			return
		}
		fmt.Println("Suggested Command:", command)
	},
}

func init() {
	//  suggest command is a subcommand of the root command.
	rootCmd.AddCommand(suggestCmd)
}

// queryNLP sends a POST request to the Python API and returns the generated command.
func queryNLP(query string) (string, error) {
	apiURL := "http://localhost:5000/query"
	payload := map[string]string{
		"query": query,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response map[string]string
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}
	return response["command"], nil
}
