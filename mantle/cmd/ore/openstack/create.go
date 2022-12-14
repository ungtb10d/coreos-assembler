// Copyright 2018 Red Hat
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openstack

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	coreosarch "github.com/coreos/stream-metadata-go/arch"
)

var (
	cmdCreate = &cobra.Command{
		Use:   "create-image",
		Short: "Create image on OpenStack",
		Long: `Upload an image to OpenStack.

After a successful run, the final line of output will be the ID of the image.
`,
		RunE: runCreate,

		SilenceUsage: true,
	}

	path       string
	name       string
	arch       string
	visibility string
	protected  bool
)

func init() {
	OpenStack.AddCommand(cmdCreate)
	cmdCreate.Flags().StringVar(&arch, "arch", coreosarch.CurrentRpmArch(), "The architecture of the image")
	cmdCreate.Flags().StringVar(&path, "file", "", "path to OpenStack image")
	cmdCreate.Flags().StringVar(&name, "name", "", "image name")
	cmdCreate.Flags().StringVar(&visibility, "visibility", "private", "Image visibility within OpenStack")
	cmdCreate.Flags().BoolVar(&protected, "protected", false, "Image deletion protection")
}

func runCreate(cmd *cobra.Command, args []string) error {
	if path == "" {
		fmt.Fprintf(os.Stderr, "--file is required\n")
		os.Exit(1)
	}
	id, err := API.UploadImage(name, path, arch, visibility, protected)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create image: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(id)
	return nil
}
