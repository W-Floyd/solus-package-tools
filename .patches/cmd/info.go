---
+++
@@ -1,56 +1,51 @@
-// Copyright © 2019 William Floyd <william.png2000@gmail.com>
-//
-// Permission is hereby granted, free of charge, to any person obtaining a copy
-// of this software and associated documentation files (the "Software"), to deal
-// in the Software without restriction, including without limitation the rights
-// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
-// copies of the Software, and to permit persons to whom the Software is
-// furnished to do so, subject to the following conditions:
-//
-// The above copyright notice and this permission notice shall be included in
-// all copies or substantial portions of the Software.
-//
-// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
-// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
-// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
-// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
-// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
-// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
-// THE SOFTWARE.
-
 package cmd
 
 import (
 	"fmt"
+	"sort"
+	"strconv"
 
+	"github.com/W-Floyd/solus-package-tools/solus-package-util/cmd/packages"
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
+	Short: "Provides information on given packages.",
+	Long: `Provides information on given packages.
+
+Largely just used for testing.`,
+	Args:      packages.InputCheckPackage,
+	ValidArgs: packages.ListNames(),
 	Run: func(cmd *cobra.Command, args []string) {
-		fmt.Println("info called")
+		packages := packages.List()
+		for i, arg := range args {
+			fmt.Println("Name      : " + packages[arg].Name)
+			fmt.Println("Version   : " + packages[arg].Version)
+			fmt.Println("Release   : " + strconv.Itoa(packages[arg].Release))
+
+			fmt.Print("Builddeps : ")
+
+			if 0 < len(packages[arg].Builddeps) {
+				fmt.Println("")
+				d := packages[arg].Builddeps
+				sort.Strings(d)
+				for _, dep := range d {
+					fmt.Println(" - " + dep)
+				}
+			} else {
+				fmt.Println("None")
+			}
+			fmt.Println("Release   : " + strconv.Itoa(packages[arg].Release))
+
+			if i < len(args)-1 {
+				fmt.Println("")
+			}
+		}
 	},
 }
 
 func init() {
 	rootCmd.AddCommand(infoCmd)
-
-	// Here you will define your flags and configuration settings.
-
-	// Cobra supports Persistent Flags which will work for this command
-	// and all subcommands, e.g.:
-	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")
-
-	// Cobra supports local flags which will only run when this command
-	// is called directly, e.g.:
-	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
 }
