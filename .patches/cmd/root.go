---
+++
@@ -23,8 +23,10 @@
 
 import (
 	"fmt"
+	"log"
 	"os"
 
+	"github.com/W-Floyd/solus-package-tools/solus-package-tools/cmd/pkgconfig"
 	"github.com/spf13/cobra"
 
 	homedir "github.com/mitchellh/go-homedir"
@@ -35,22 +37,21 @@
 
 // rootCmd represents the base command when called without any subcommands
 var rootCmd = &cobra.Command{
-	Use:   "solus-package-tools",
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
+	Use:   "spt",
+	Short: "Solus Package Tools - A utility to help build unofficial package sets for Solus.",
+	Long: `Solus Package Tools
+A utility to manage the building of large unofficial package sets for Solus.
+Specifically created for the development of the Deepin DE on Solus.`,
 }
 
 // Execute adds all child commands to the root command and sets flags appropriately.
 // This is called by main.main(). It only needs to happen once to the rootCmd.
 func Execute() {
+	err := pkgconfig.LoadPkgConfigDictionary()
+	if err != nil {
+		log.Fatal("LoadPkgConfigDictionary failed")
+	}
+
 	if err := rootCmd.Execute(); err != nil {
 		fmt.Println(err)
 		os.Exit(1)
@@ -64,7 +65,7 @@
 	// Cobra supports persistent flags, which, if defined here,
 	// will be global for your application.
 
-	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.solus-package-tools.yaml)")
+	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.solus-package-tools.yaml and ./.solus-package-tools.yaml)")
 
 	// Cobra also supports local flags, which will only run
 	// when this action is called directly.
@@ -86,13 +87,12 @@
 
 		// Search config in home directory with name ".solus-package-tools" (without extension).
 		viper.AddConfigPath(home)
+		viper.AddConfigPath(".")
 		viper.SetConfigName(".solus-package-tools")
 	}
 
 	viper.AutomaticEnv() // read in environment variables that match
 
 	// If a config file is found, read it in.
-	if err := viper.ReadInConfig(); err == nil {
-		fmt.Println("Using config file:", viper.ConfigFileUsed())
-	}
+	viper.ReadInConfig()
 }
