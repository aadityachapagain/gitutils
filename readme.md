# gitutils

## Motivation

github cli tool [gh](https://cli.github.com) doesnot have capability to login into multiple github. The real pain is when you have to switch between different user. You have to manually logout and login with github cli , which I feel uncessary and time consuming.

Hence, this small piece of program will keep track of your github users for you when github cli won't, **because I care about you.**

## Installation

Currently It only support installation through **Go**.

> go install github.com/aadityachapagain/gitutils

## Usage

> gitutils depends on github cli tool so you need to initialize github cli tool first [github cli](https://cli.github.com)

* login to your github account using github cli
* Use github_token for authentication. ( for security reasons only use Oauth_token when authentication through github cli)

* Update gitutils node to detect authenticated users
    > gitutils update
    * run `gitutils update` every time you login into new users

* List githubusers 
    > gitutils list

* Switch to specifed user
    > gitutils switch <github_username>

## Improvements
- Add device authorization through cli
- Automatic syncing of github credentials