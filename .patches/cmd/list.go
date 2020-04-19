---
+++
@@ -1,57 +1,33 @@
-/*
-Copyright © 2020 William Floyd <william.png2000@gmail.com>
-
-Permission is hereby granted, free of charge, to any person obtaining a copy
-of this software and associated documentation files (the "Software"), to deal
-in the Software without restriction, including without limitation the rights
-to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
-copies of the Software, and to permit persons to whom the Software is
-furnished to do so, subject to the following conditions:
-
-The above copyright notice and this permission notice shall be included in
-all copies or substantial portions of the Software.
-
-THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
-IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
-FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
-AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
-LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
-OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
-THE SOFTWARE.
-*/
 package cmd
 
 import (
 	"fmt"
+	"strconv"
 
+	"github.com/W-Floyd/solus-package-tools/solus-package-tools/cmd/packages"
 	"github.com/spf13/cobra"
 )
 
 // listCmd represents the list command
 var listCmd = &cobra.Command{
 	Use:   "list",
-	Short: "A brief description of your command",
-	Long: `A longer description that spans multiple lines and likely contains examples
-and usage of using your command. For example:
-
-Cobra is a CLI library for Go that empowers applications.
-This application is a tool to generate the needed files
-to quickly create a Cobra application.`,
+	Short: "Lists all packages in the current folder.",
+	Long: `Lists all packages in the current folder.
+
+Specifically, folders that are not filtered, and have a 'package.yml' file.`,
+	Args: cobra.NoArgs,
 	Run: func(cmd *cobra.Command, args []string) {
-		fmt.Println("list called")
+
+		for _, p := range packages.List() {
+			fmt.Println("Name: " + p.Name)
+			fmt.Println("Version: " + p.Version)
+			fmt.Println("Release: " + strconv.Itoa(p.Release))
+			fmt.Println("")
+		}
+
 	},
 }
 
 func init() {
 	rootCmd.AddCommand(listCmd)
-
-	// Here you will define your flags and configuration settings.
-
-	// Cobra supports Persistent Flags which will work for this command
-	// and all subcommands, e.g.:
-	// listCmd.PersistentFlags().String("foo", "", "A help for foo")
-
-	// Cobra supports local flags which will only run when this command
-	// is called directly, e.g.:
-	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
 }
