# cfna2018-app - A Test Application for Cloud Foundry Summit Labs

This repository contains the code for an example Go application
that will be used in the Cloud Foundry North America Summit 2018,
held in Boston.  Included in this repository is a Concourse
pipeline that handles running tests against all commits made to
the `master` branch.

The lab then discusses the theory of running tests on commits made
against arbitrary Github Pull Requests; in the hands-on segment of
the lab, you get to build out that functionality for real.

To get started, you'll need to fork this repo into your own Github
org, and generate a personal access token for Concourse to use.

For the lab, we have provided a jumpbox with all the necessary
tools on it, and a Concourse installation.  For self-study, all
you need is an instance of the jumpbox BOSH release, deployed
somewhere near your Concourse.

## Setup Instructions
### Step 1
Fork the cfna2018-app repository, then create a personal access token with repo scope granted.

### Step 2
SSH into the jumpbox: 
`ssh cfdev#@pr-lab.cfna2018.starkandwayne.com`

_(replace # with the number given to you)_

### Step 3
Clone the repository:
`git clone https://github.com/<yourgituser>/cfna2018-app.git`

_(replace <yourgituser> with your GitHub account name)_

### Step 4
Within the repository folder, edit the file `ci/settings.yml` and input your
information.

_(Nano, Vim, and Emacs are installed on the jumpbox)_

### Step 5
Run `ci/repipe` to push the pipeline to Concourse.

### Step 6 to End
Continue editing `ci/pipeline.yml` until you have GitHub PR Testing


## Links & Commands
### URLS
- GitHub Repository: https://github.com/starkandwayne/cfna2018-app
- GitHub Create Personal Token: https://github.com/settings/tokens/new
- Concourse: http://ci.system.cfna2018.starkandwayne.com
- GitHub Pull Request Resource: https://github.com/jtarchie/github-pullrequest-resource
### SSH Information
Server: pr-lab.cfna2018.starkandwayne.com
