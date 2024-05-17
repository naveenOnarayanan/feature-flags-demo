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
package co.gitar;

import com.launchdarkly.sdk.*;
import com.launchdarkly.sdk.server.*;

class Sample {

  public void sampleMethodCheckingFlag() {
    LDClient client = new LDClient("sdk-key-123abc");
    LDContext context = LDContext
      .builder("context-key-123abc")
      .name("Sandy")
      .build();
    // Check the 'is-dark-mode-enabled' feature flag for the specified user
    boolean flagValue = client.boolVariation("is-dark-mode-enabled", context, false);
    if (flagValue) {
      // The code to run if the feature is on
      System.out.println("Dark mode!");
    } else {
      // The code to run if the feature is off
      System.out.println("My eyes, they burn!");
    }
  }

}
