workflow "Build" {
  on = "push"
  resolves = ["Build_action"]
}

action "Build_action" {
  uses = "apex/actions/go@master"
  args = "build -o lmwtfy cmd/main.go"
}
