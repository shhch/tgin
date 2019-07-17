package tool

type Param struct {
	Password string
	Ip string
	Swoole string
	Paltform string
	Web string
	Client string
	Datrixlog string
}

func CheckBuild() (string, error) {
	command := "ps aux|grep /data/git-develop-build/|grep -v grep|awk '{print $2}'"
	output, err := RunCmd(command)
	return output, err
}

func Build() () {

}
