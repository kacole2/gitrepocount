# OpenFaaS gitrepocount

[![](https://images.microbadger.com/badges/version/kacole2/gitrepocount.svg)](https://microbadger.com/images/kacole2/gitrepocount "Get your own version badge on microbadger.com") [![](https://images.microbadger.com/badges/image/kacole2/gitrepocount.svg)](https://microbadger.com/images/kacole2/gitrepocount "Get your own image badge on microbadger.com")

This repo contains an [OpenFaaS](https://github.com/openfaas/faas) function which uses the [google/go-github](https://github.com/google/go-github) Go library for the GitHub API to count how many repos an individual or group has. This uses pagination to continue counting.

## Deploying the Function

Make sure you have deployed an OpenFaaS stack to your cluster using the instructions on the [OpenFaaS repo](https://github.com/openfaas/faas).

### Use the CLI (`faas-cli`)

**Get the CLI**

The [faas-cli](https://github.com/openfaas/faas-cli) can be installed via `brew install faas-cli` or `curl -sSL https://get.openfaas.com | sudo sh`.

Now deploy the function as follows:

```
# git clone https://github.com/kacole2/gitrepocount
# cd gitrepocount
// make any changes to the gitrepocount.yml file for OpenFaaS API gateway
# faas-cli deploy -f gitrepocount.yml
200 OK
URL: http://localhost:31112/function/gitrepocount
```

### Testing the Function
Now that the function is running in your OpenFaaS environment you can test it from the command line by running:

```
$ curl localhost:31112/function/gitrepocount -d "kacole2"
$ {"user":"kacole2","repocount":58}
```

### License
Licensed under the Apache License, Version 2.0 (the “License”); you may not use this file except in compliance with the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
