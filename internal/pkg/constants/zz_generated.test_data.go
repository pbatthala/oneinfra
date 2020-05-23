/**
 * Copyright 2020 Rafael Fernández López <ereslibre@ereslibre.es>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 **/

package constants

// Code generated by release-gen script. DO NOT EDIT.

const (
	// RawTestData represents the supported versions for this release
	RawTestData = `# This file contains information on how to build container images with
# the correct components and their versions for ease and faster
# testing.
#
# This information is not used for production environments.

containerdVersions:
- version: "1.3.4"
  criToolsVersion: "1.18.0"
  cniPluginsVersion: "0.8.6"
kubernetesVersions:
  "1.15.12":
    containerdVersion: "1.3.4"
    pauseVersion: "3.1"
  "1.16.9":
    containerdVersion: "1.3.4"
    pauseVersion: "3.1"
  "1.17.5":
    containerdVersion: "1.3.4"
    pauseVersion: "3.1"
  "1.18.2":
    containerdVersion: "1.3.4"
    pauseVersion: "3.1"
  "1.19.0-alpha.3":
    containerdVersion: "1.3.4"
    pauseVersion: "3.1"`
)
