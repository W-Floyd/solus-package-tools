---
+++
@@ -1,33 +1,13 @@
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
-	"github.com/spf13/cobra"
 	"os"
 
-	homedir "github.com/mitchellh/go-homedir"
+	"github.com/spf13/cobra"
 	"github.com/spf13/viper"
+
+	homedir "github.com/mitchellh/go-homedir"
 )
 
 var cfgFile string
@@ -35,16 +15,10 @@
 // rootCmd represents the base command when called without any subcommands
 var rootCmd = &cobra.Command{
 	Use:   "solus-package-tools",
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
@@ -59,14 +33,8 @@
 func init() {
 	cobra.OnInitialize(initConfig)
 
-	// Here you will define your flags and configuration settings.
-	// Cobra supports persistent flags, which, if defined here,
-	// will be global for your application.
-
 	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.solus-package-tools.yaml)")
 
-	// Cobra also supports local flags, which will only run
-	// when this action is called directly.
 	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
 }
 
