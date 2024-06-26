/*
Copyright © 2024 Elouan GOUINGUENET <elouangouinguenet@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize new password storage and use gpg-id for encryption. Selectively reencrypt existing passwords using new gpg-id.",
	Long: `Initialize  new  password storage and use gpg-id for encryption. Multiple gpg-ids may be specified, in order to encrypt each password with multiple ids.
This command must be run first before a password store can be used. If the specified gpg-id is different from the key used in any existing files,  these
files  will  be  reencrypted to use the new id.  Note that use of gpg-agent(1) is recommended so that the batch decryption does not require as much user
intervention. If --path or -p is specified, along with an argument, a specific gpg-id or set of gpg-ids is assigned for that specific sub folder of  the
password  store. If only one gpg-id is given, and it is an empty string, then the current .gpg-id file for the specified sub-folder (or root if unspeci‐
fied) is removed.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("path", "p", "", "Specific gpg-id, or set of gpg-ids, for specific sub-folder of the password store")
}
