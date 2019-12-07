
# A sample application for learning `golang`


#To update/regenerate go repos
Remove the following lines from WORKSPACE
```
load("//:go_repositories.bzl", "go_repositories")
go_repositories()
```

Then regenerate the file using the below command

`bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=go_repositories.bzl%go_repositories`

#To build
`bazel run uaa/uaa-server`

