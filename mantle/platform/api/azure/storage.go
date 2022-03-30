// Copyright 2016 CoreOS, Inc.
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

package azure

import (
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/arm/storage"
)

func (a *API) GetStorageServiceKeysARM(account, resourceGroup string) (storage.AccountListKeysResult, error) {
	return a.accClient.ListKeys(resourceGroup, account)
}

func (a *API) CreateStorageAccount(resourceGroup string) (string, error) {
	// Only lower-case letters & numbers allowed in storage account names
	name := strings.Replace(randomName("kolasa"), "-", "", -1)
	parameters := storage.AccountCreateParameters{
		Sku: &storage.Sku{
			Name: "Standard_LRS",
		},
		Kind:     "Storage",
		Location: &a.opts.Location,
	}
	_, err := a.accClient.Create(resourceGroup, name, parameters, nil)
	if err != nil {
		return "", fmt.Errorf("creating storage account: %v", err)
	}
	return name, nil
}
