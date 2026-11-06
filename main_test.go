package main
import("strings";"testing")
func errorsIn(issues []Issue)int{n:=0;for _,i:=range issues{if i.Severity=="error"{n++}};return n}
func TestValidAndDuplicateCommands(t *testing.T){base:="[Unit]\nDescription=Worker\n[Service]\nExecStart=/usr/bin/true\n";if len(check(base))!=0{t.Fatal(check(base))};if errorsIn(check(base+"ExecStart=/usr/bin/false\n"))!=1{t.Fatal("duplicate not detected")};if errorsIn(check(base+"Type=oneshot\nExecStart=/usr/bin/false\n"))!=0{t.Fatal("valid oneshot rejected")}}
func TestResetAndExceptionalStopOnly(t *testing.T){text:="[Unit]\nDescription=Test\n[Service]\nExecStart=/a\nExecStart=\nExecStart=/b\n";if errorsIn(check(text))!=0{t.Fatal(check(text))};text="[Unit]\nDescription=Test\n[Service]\nRemainAfterExit=yes\nExecStop=/usr/bin/true\n";if errorsIn(check(text))!=0{t.Fatal("stop-only service rejected")}}
func TestMalformed(t *testing.T){for _,text:=range []string{"x=y\n", "[Service]\nbroken\n", "[Service]\nType=dbus\nExecStart=/a\n"}{if errorsIn(check(text))==0{t.Fatal("accepted",strings.TrimSpace(text))}}}
