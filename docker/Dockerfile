# Copyright 2017 Google Inc. All Rights Reserved.
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

# The repository does not use version tags, so the only way to pin a version is to use the full hash.
# The hash below is 'latest' as of 2024-11-12
FROM  gcr.io/cloud-marketplace/google/debian12@sha256:ce96510e0ac954fbc65db75901c71766f0137fcd16a8915f717e1fb6fbe0f2a3

RUN apt-get update && apt-get install -y iptables mount e2fsprogs util-linux
COPY gce-containers-startup /bin/gce-containers-startup

CMD ["/bin/gce-containers-startup"]

