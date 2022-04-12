package ls

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)
var echoTimes int

var cmdPrint = &cobra.Command{
	Use:   "Print Something",
	Short: "Print what you input to screen",
	Long:  "Long description of this command",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, ""))
	},
}

var cmdEcho = &cobra.Command{
	Use:   "Echo something",
	Short: "Echo something on screen",
	Long:  "Long description of echo",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}
var cmdTimes = &cobra.Command{
    Use:   "times [# times] [string to echo]",
    Short: "Echo anything to the screen more times",
    Long: `echo things multiple times back to the user by providing
a count and a string.`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      for i := 0; i < echoTimes; i++ {
        fmt.Println("Echo: " + strings.Join(args, " "))
      }
    },
  }

func init(){
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
}

var rootCmd = &cobra.Command{Use: "app"}
rootCmd.AddCommand(cmdPrint,cmdEcho)
cmdEcho.AddCommand(cmdTimes)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
