// 参考: https://text.baldanders.info/golang/detecting-character-encoding/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version string = "0.0.1"

// Main logic
func Version(cmd *cobra.Command, args []string) {
	fmt.Println(version)
}
