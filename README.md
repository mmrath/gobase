
# A sample application for learning `golang`


#To update go repos
`bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=go_repositories.bzl%go_repositories`

#To build
`bazel run uaa/uaa-server`

