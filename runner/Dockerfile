# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.20.4
LABEL maintainer="Krzysztof Kwaśniewski <krzyk@google.com>"

ENV GOPATH /gopath
ENV PATH $GOPATH/bin:$PATH

# Add Google Cloud repository
RUN curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | gpg --dearmor | tee /usr/share/keyrings/cloud.google.gpg > /dev/null 2>&1
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list

RUN apt-get update && apt-get --yes install google-cloud-sdk \
	google-cloud-sdk-gke-gcloud-auth-plugin kubectl libseccomp-dev \
	&& go version && go install github.com/jstemmer/go-junit-report@v0.9.1

# Add link for the Google Cloud SDK as installed in Prow. Without the link,
# Go client library fails to authenticate.
RUN ln -s /usr/lib/google-cloud-sdk /google-cloud-sdk

CMD ["/bin/bash"]
