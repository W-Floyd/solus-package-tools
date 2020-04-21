---
+++
@@ -22,23 +22,19 @@
 package cmd
 
 import (
-	"fmt"
-
+	"github.com/W-Floyd/solus-package-tools/solus-package-tools/cmd/eopkg"
 	"github.com/spf13/cobra"
 )
 
 // infoCmd represents the info command
 var infoCmd = &cobra.Command{
 	Use:   "info",
-	Short: "A brief description of your command",
-	Long: `A longer description that spans multiple lines and likely contains examples
-and usage of using your command. For example:
-
-Cobra is a CLI library for Go that empowers applications.
-This application is a tool to generate the needed files
-to quickly create a Cobra application.`,
+	Short: "Extracts and displays JSON metadata for a .eopkg file",
+	Long:  `Extracts and displays JSON metadata for a .eopkg file`,
 	Run: func(cmd *cobra.Command, args []string) {
-		fmt.Println("info called")
+		for _, input := range args {
+			eopkg.PrintMetaDataJSON(input)
+		}
 	},
 }
 
