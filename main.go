// A conservative, read-only linter for standalone Linux systemd service files.
package main
import("bufio";"fmt";"os";"strings")
type Issue struct{Line int;Severity,Message string}
type Entry struct{Line int;Value string}
func check(text string)[]Issue{
 issues:=[]Issue{};sections:=map[string]map[string][]Entry{};section:=""
 scanner:=bufio.NewScanner(strings.NewReader(text));scanner.Buffer(make([]byte,4096),1024*1024);line:=0
 add:=func(n int,severity,message string){issues=append(issues,Issue{n,severity,message})}
 for scanner.Scan(){line++;value:=strings.TrimSpace(scanner.Text());if value==""||strings.HasPrefix(value,"#")||strings.HasPrefix(value,";"){continue}
  if strings.HasSuffix(value,`\`){add(line,"warning","continuation lines are outside this linter's supported subset")}
  if strings.HasPrefix(value,"[")&&strings.HasSuffix(value,"]"){
   section=strings.TrimSpace(value[1:len(value)-1]);if sections[section]==nil{sections[section]=map[string][]Entry{}}
   if section!="Unit"&&section!="Service"&&section!="Install"{add(line,"warning","unrecognized service section: "+section)};continue
  }
  key,val,ok:=strings.Cut(value,"=");key=strings.TrimSpace(key);val=strings.TrimSpace(val)
  if !ok||key==""{add(line,"error","expected Key=Value or [Section]");continue}
  if section==""{add(line,"error","setting appears before a section");continue}
  if val==""{sections[section][key]=nil}else{sections[section][key]=append(sections[section][key],Entry{line,val})}
 }
 if err:=scanner.Err();err!=nil{add(line,"error","input line exceeds supported size or cannot be read")}
 get:=func(s,k string)string{v:=sections[s][k];if len(v)==0{return ""};return v[len(v)-1].Value}
 if sections["Service"]==nil{add(0,"error","missing [Service] section")}
 if get("Unit","Description")==""{add(0,"warning","add a human-readable [Unit] Description")}
 kind:=get("Service","Type");if kind==""{kind="simple"}
 allowed:=map[string]bool{"simple":true,"exec":true,"forking":true,"oneshot":true,"dbus":true,"notify":true,"notify-reload":true,"idle":true}
 if !allowed[kind]{add(0,"error","unsupported Type="+kind)}
 starts:=sections["Service"]["ExecStart"]
 if len(starts)==0 && !(get("Service","RemainAfterExit")=="yes"&&get("Service","ExecStop")!=""){add(0,"error","missing ExecStart (or RemainAfterExit=yes plus ExecStop)")}
 if len(starts)>1&&kind!="oneshot"{add(starts[1].Line,"error","multiple ExecStart commands require Type=oneshot")}
 if kind=="dbus"&&get("Service","BusName")==""{add(0,"error","Type=dbus requires BusName")}
 if kind=="oneshot"&&(get("Service","Restart")=="always"||get("Service","Restart")=="on-success"){add(0,"error","oneshot cannot restart always or on-success")}
 for _,s:=range starts{command:=strings.Fields(s.Value);if len(command)>0&&strings.ContainsAny(command[0],"|><"){add(s.Line,"warning","ExecStart is not interpreted by a shell; use an explicit shell only if intended")}}
 if get("Service","User")=="root"{add(0,"warning","service explicitly runs as root; review whether that is required")}
 return issues
}
func main(){
 if len(os.Args)!=2||os.Args[1]=="--help"{fmt.Println("Usage: unit-check FILE.service\nRead-only subset linter. Follow with systemd-analyze verify on your Linux host.");if len(os.Args)!=2{os.Exit(2)};return}
 data,err:=os.ReadFile(os.Args[1]);if err!=nil{fmt.Fprintln(os.Stderr,err);os.Exit(2)}
 issues:=check(string(data));code:=0;for _,issue:=range issues{fmt.Printf("%s:%d: %s: %s\n",os.Args[1],issue.Line,issue.Severity,issue.Message);if issue.Severity=="error"{code=2}else if code==0{code=1}}
 if len(issues)==0{fmt.Println("No issues in supported subset. Full systemd validation is still recommended.")};os.Exit(code)
}
