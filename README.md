# gradle-demo

This is my first project using `gradle` and `Go`.


## Gradle
  - `build`: basic build command for unix executable, placed in *out/*
  - `my_clean`: clean task to remove the executable
  - `copy`: task copying files from one location to another AND task dependency on `my_clean`
  - `zip`: demonstrates zipping contents of a directory


## Go
`go-gradle.go` contains source code with a few features demonstrated. `struct House` is defined as a type with two members; an object is constructed with both members being set, and then printed.
`struct Transition` has more members, and more functions defined on it.  This was a test for my Push Down Automaton project
