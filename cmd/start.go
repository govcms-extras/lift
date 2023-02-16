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

	"github.com/govcms-extras/lift/pkg/utils/docker"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the service",
	Long:  "Start the service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")

		docker.DockerContainerCreate()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
