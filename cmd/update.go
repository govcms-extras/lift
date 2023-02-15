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
	"github.com/govcms-extras/lift/pkg/utils/docker"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Pulls Docker images and recreates the containers",
	Long:  "Pulls Docker images and recreates the containers",
	Run: func(cmd *cobra.Command, args []string) {
		docker.DockerImagePull()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
