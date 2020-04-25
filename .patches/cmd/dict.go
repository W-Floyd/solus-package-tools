---
+++
@@ -22,23 +22,18 @@
 package cmd
 
 import (
-	"fmt"
-
+	"github.com/W-Floyd/solus-package-tools/solus-package-tools/cmd/pkgconfig"
 	"github.com/spf13/cobra"
 )
 
 // dictCmd represents the dict command
 var dictCmd = &cobra.Command{
 	Use:   "dict",
-	Short: "A brief description of your command",
-	Long: `A longer description that spans multiple lines and likely contains examples
-and usage of using your command. For example:
-
-Cobra is a CLI library for Go that empowers applications.
-This application is a tool to generate the needed files
-to quickly create a Cobra application.`,
+	Short: "Update pkgconfig dictionary from file",
+	Long:  `Update pkgconfig dictionary from file`,
 	Run: func(cmd *cobra.Command, args []string) {
-		fmt.Println("dict called")
+		pkgconfig.UpdateDictionary()
+		pkgconfig.WriteDictionary()
 	},
 }
 
@@ -53,5 +48,5 @@
 
 	// Cobra supports local flags which will only run when this command
 	// is called directly, e.g.:
-	// dictCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
+	dictCmd.Flags().BoolP("delete", "d", false, "Deletes the old dictionary")
 }
