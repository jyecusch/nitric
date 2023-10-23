# Copyright Nitric Pty Ltd.

# SPDX-License-Identifier: Apache-2.0

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at:

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

ARG BASE_IMAGE

# Wrap any base image in this runtime wrapper
FROM ${BASE_IMAGE}

# ARG RUNTIME_URI
ARG RUNTIME_FILE

COPY ${RUNTIME_FILE} /bin/runtime
RUN chmod +x-rw /bin/runtime

# Inject original wrapped command here
CMD [%s]
ENTRYPOINT ["/bin/runtime"]
