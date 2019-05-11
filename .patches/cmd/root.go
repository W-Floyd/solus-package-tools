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
@@ -31,19 +11,12 @@
 
 var cfgFile string
 
-// rootCmd represents the base command when called without any subcommands
 var rootCmd = &cobra.Command{
 	Use:   "solus-package-util",
-	Short: "A brief description of your application",
-	Long: `A longer description that spans multiple lines and likely contains
-examples and usage of using your application. For example:
-
-Cobra is a CLI library for Go that empowers applications.
-This application is a tool to generate the needed files
-to quickly create a Cobra application.`,
-	// Uncomment the following line if your bare application
-	// has an action associated with it:
-	//	Run: func(cmd *cobra.Command, args []string) { },
+	Short: "A utility to manage the building of large unofficial package sets for Solus.",
+	Long: `A utility to manage the building of large unofficial package sets for Solus.
+
+Specifically created for the development of the Deepin DE on Solus.`,
 }
 
 // Execute adds all child commands to the root command and sets flags appropriately.
