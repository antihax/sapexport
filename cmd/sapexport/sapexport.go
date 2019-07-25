package sapexport

import (
	"fmt"
	"os"

	"github.com/sap/gorfc/gorfc"
	"github.com/spf13/cobra"
)

var abapSystem gorfc.ConnectionParameter
var rootCmd = &cobra.Command{Use: "sapexport"}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&abapSystem.User, "user", "u", getenv("SAPRFC_USER", ""), "RFC Username (or env SAPRFC_USER)")
	rootCmd.PersistentFlags().StringVarP(&abapSystem.Passwd, "pass", "p", "", "RFC Password (or env SAPRFC_PASS)")
	rootCmd.PersistentFlags().StringVarP(&abapSystem.Lang, "language", "l", getenv("SAPRFC_LANGUAGE", "EN"), "System Language (or env SAPRFC_LANGUAGE)")
	rootCmd.PersistentFlags().StringVarP(&abapSystem.Client, "client", "c", getenv("SAPRFC_CLIENT", "001"), "System Client (or env SAPRFC_CLIENT)")
	rootCmd.PersistentFlags().StringVarP(&abapSystem.Ashost, "address", "a", getenv("SAPRFC_ADDRESS", ""), "System Address (or env SAPRFC_ADDRESS)")
	rootCmd.PersistentFlags().StringVar(&abapSystem.Sysnr, "sysnr", "", "System Instance")
	rootCmd.PersistentFlags().StringVarP(&abapSystem.Mshost, "msgserver", "m", getenv("SAPRFC_MSGSVR", ""), "System Address (or env SAPRFC_MSGSVR)")
	rootCmd.PersistentFlags().StringVarP(&abapSystem.Saprouter, "router", "r", getenv("SAPRFC_ROUTER", ""), "Router (or env SAPRFC_ROUTER)")
	rootCmd.Execute()
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
