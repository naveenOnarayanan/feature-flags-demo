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

/**
 * Example of using the LaunchDarkly JavaScript SDK to evaluate feature flags.
 *
 * https://docs.launchdarkly.com/sdk/client-side/javascript
 */

import * as LaunchDarkly from "@launchdarkly/node-server-sdk";

const SDK_KEY = "<SDK_KEY>";
const context = {
	kind: "user",
	name: "Sandy",
	key: "example-context-key",
};
const client = LaunchDarkly.init(SDK_KEY);

client.once("ready", () => {
	client.variation(
		"is-dark-mode-enabled",
		context,
		false,
		(err, useDarkMode) => {
			if (useDarkMode) {
				console.log("Dark mode is enabled");
			} else {
				console.log("Dark mode is disabled");
			}
		},
	);
});
