package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/its-a-feature/Mythic/Mythic_CLI/cmd/internal"
	"github.com/spf13/cobra"
)

// configGetCmd represents the configGet command
var configGetCmd = &cobra.Command{
	Use:   "get <configuration> <configuration> ...",
	Short: "Get the specified configuration values",
	Long: `Get the specified configuration values. You can provide one value or
a list of values separated by spaces.
For example: mythic-cli config get ADMIN_PASSWORD POSTGRES_PASSWORD`,
	Run: configGet,
}

func init() {
	configCmd.AddCommand(configGetCmd)
}

func configGet(cmd *cobra.Command, args []string) {
	// initialize tabwriter
	writer := new(tabwriter.Writer)
	// Set minwidth, tabwidth, padding, padchar, and flags
	writer.Init(os.Stdout, 8, 8, 1, '\t', 0)

	defer writer.Flush()

	fmt.Println("[+] Getting configuration values:")
	fmt.Fprintf(writer, "\n %s\t%s", "Setting", "Value")
	fmt.Fprintf(writer, "\n %s\t%s", "–––––––", "–––––––")

	configuration := internal.GetConfigStrings(args)
	for key, val := range configuration {
		fmt.Fprintf(writer, "\n %s\t%s", strings.ToUpper(key), val)
	}
	fmt.Fprintln(writer, "")
}
