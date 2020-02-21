# go-package-base
A base template to start new packages from

## Getting Started

* Create a new repository using github template
  * Go to create new repository in github
  * Choose a repository templace (promoboxx/go-package-base)
  * Create repository
* Initialize the new package
  * Clone the repository to your local machine
  * Run the init script

```sh
git clone git@github.com:promoboxx/<my_new_package_name>.git
cd new-service
sh init.sh <my_new_package_name>

# govendor init
# govendor update +external
# or if you do not have all the dependencies on your local machine
# govendor fetch +outside

git push -u origin master
```
