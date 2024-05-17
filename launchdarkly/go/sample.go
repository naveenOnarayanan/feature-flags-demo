/*
 * Copyright (c) 2024 Gitar, Inc.
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
 */
package main

import (
	"fmt"
	"gopkg.in/launchdarkly/go-server-sdk.v5"
	"gopkg.in/launchdarkly/go-server-sdk.v5/lduser"
)

func checkFeatureFlags(user lduser.User, ldClient *ld.LDClient) {
	// Check the 'is-dark-mode-enabled' feature flag for the specified user
	darkModeEnabled, err := ldClient.BoolVariation("is-dark-mode-enabled", user, false)
	if err != nil {
		fmt.Printf("Error evaluating feature flag 'is-dark-mode-enabled': %s\n", err)
		return
	}

	// Report the status of the dark mode
	fmt.Println("Dark mode is", enabledStatus(darkModeEnabled))

	// If dark mode is enabled, check for high contrast mode
	if darkModeEnabled {
		// Check the 'enable-high-contrast' feature flag for the specified user
		highContrastEnabled, err1 := ldClient.BoolVariation("enable-high-contrast", user, false)
		if err1 != nil {
			fmt.Printf("Error evaluating feature flag 'enable-high-contrast': %s\n", err1)
			return
		}
		// Report the status of the high contrast mode
		fmt.Println("High contrast mode is", enabledStatus(highContrastEnabled))
	}
}
