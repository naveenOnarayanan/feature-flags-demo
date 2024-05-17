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
	"flagfn"
	"fmt"
)

func check_if_variant() {
	if flag_key_123abc {
		fmt.Println("Enabled pattern flag_key_123abc")
	} else {
		fmt.Println("Disabled pattern flag_key_123abc")
	}
}
