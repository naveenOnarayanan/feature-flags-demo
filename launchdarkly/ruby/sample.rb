#
# Copyright (c) 2024 Gitar, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

module WelcomeHelper
    def check_feature_flags(user, ld_client)
        # If dark mode is enabled, check for high contrast mode
        if ld_client.variation("is-dark-mode-enabled", user, false)
            # Check and report the 'enable-high-contrast' feature flag for the specified user
            high_contrast = ld_client.variation("enable-high-contrast", user, false)
            puts "Dark mode - high contrast mode is #{high_contrast}"
        else
            puts "Light mode!"
        end
    end
    
    def enabled_status(flag_value)
        flag_value ? "enabled" : "disabled"
    end
end
