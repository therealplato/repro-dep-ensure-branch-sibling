Start by checking out `feature` branch.

Objective:
`go run parent/main.go` on `feature` branch outputs "updated dependant",
without manually manipulating vendor folder,
and without committing dependant/ changes to master branch.

Ideally there's a way to tell dep "do not vendor github.com/therealplato/repro-dep-ensure-branch-sibling/dependant, just use your sibling in
this working tree".

Symptom:
./parent.go:11:34: cannot use "github.com/therealplato/repro-dep-ensure-branch-sibling/parent/vendor/github.com/sirupsen/logrus".Fields literal (type "github.com/therealplato/repro-dep-ensure-branch-sibling/parent/vendor/github.com/sirupsen/logrus".Fields) as type "github.com/sirupsen/logrus".Fields in argument to dependant.Exported
