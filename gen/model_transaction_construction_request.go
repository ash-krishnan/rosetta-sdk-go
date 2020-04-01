// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package gen

// TransactionConstructionRequest struct for TransactionConstructionRequest
type TransactionConstructionRequest struct {
	NetworkIdentifier *NetworkIdentifier `json:"network_identifier"`
	AccountIdentifier *AccountIdentifier `json:"account_identifier"`
	// Some blockchains require different metadata for different types of transaction construction (ex: delegation versus a transfer).  Instead of requiring a blockchain node to return all possible types of metadata for construction (which may require multiple node fetches), the client can specify a `method` to limit the metadata returned to only the subset required.
	Method *string `json:"method,omitempty"`
}
