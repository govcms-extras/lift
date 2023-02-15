/*
Copyright © 2023 Joseph Zhao <pandaski@outlook.com.au>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a GovCMS local instance",
	Long:  "Create a GovCMS local instance",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")

		repo := "https://github.com/govCMS/scaffold.git"
		err := GitClone(repo, "./govcms")
		if err != nil {
			fmt.Println(err)
		}
	},
}

func GitClone(repoURL, destPath string) error {
	cmd := exec.Command("git", "clone", repoURL, destPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("git clone failed: %v", err)
	}
	return nil
}

func init() {
	govcmsCmd.AddCommand(createCmd)
}
