workflow "Build" {
  on = "push"
  resolves = ["Build"]
}

action "Build" {
  uses = "apex/actions/go@master"
  args = "build -o lmwtfy cmd/main.go"
}
