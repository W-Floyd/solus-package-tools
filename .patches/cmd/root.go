---
+++
@@ -36,13 +36,10 @@
 // rootCmd represents the base command when called without any subcommands
 var rootCmd = &cobra.Command{
 	Use:   "solus-package-tools",
-	Short: "A brief description of your application",
-	Long: `A longer description that spans multiple lines and likely contains
-examples and usage of using your application. For example:
+	Short: "A utility to manage the building of large unofficial package sets for Solus.",
+	Long: `A utility to manage the building of large unofficial package sets for Solus.
 
-Cobra is a CLI library for Go that empowers applications.
-This application is a tool to generate the needed files
-to quickly create a Cobra application.`,
+Specifically created for the development of the Deepin DE on Solus.`,
 	// Uncomment the following line if your bare application
 	// has an action associated with it:
 	//	Run: func(cmd *cobra.Command, args []string) { },
