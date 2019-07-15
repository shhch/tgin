package tool

func CheckBuild() bool {
	command := "ps aux|grep /data/git-develop-build/|grep -v grep|awk '{print $2}'"
	arg := ""
	RunCmd(command, arg)
}
