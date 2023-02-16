package setup

import (
	"fmt"
	"runtime"

	"github.com/spf13/viper"
)

func Setup() {
	viper.SetDefault("defaults", true)
	// Set the default domain.
	viper.SetDefault("domain", "govcms.local")

	if viper.GetBool("defaults") {
		switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("Running on macOS")
		case "linux":
			fmt.Println("Running on Linux")
		case "windows":
			fmt.Println("Running on Windows")
		default:
			fmt.Printf("Running on %s\n", os)
		}
	}
}
