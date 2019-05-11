---
+++
@@ -1,23 +1,3 @@
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
@@ -36,6 +16,7 @@
 Cobra is a CLI library for Go that empowers applications.
 This application is a tool to generate the needed files
 to quickly create a Cobra application.`,
+	Args: cobra.MinimumNArgs(1),
 	Run: func(cmd *cobra.Command, args []string) {
 		fmt.Println("bump called")
 	},
@@ -43,14 +24,4 @@
 
 func init() {
 	gitCmd.AddCommand(bumpCmd)
-
-	// Here you will define your flags and configuration settings.
-
-	// Cobra supports Persistent Flags which will work for this command
-	// and all subcommands, e.g.:
-	// bumpCmd.PersistentFlags().String("foo", "", "A help for foo")
-
-	// Cobra supports local flags which will only run when this command
-	// is called directly, e.g.:
-	// bumpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
 }
