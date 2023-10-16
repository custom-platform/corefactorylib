package lib

import (
	"context"
	"fmt"
	"log"
	"os"
)

func GitClone(ctx context.Context, dir, repo string) {
	comando := "cd " + dir + " && git clone " + repo + " ."
	ExecCommand(ctx, comando, true)
}

func GitAdd(ctx context.Context, dir string) {
	comando := "cd " + dir + " && git add . "
	ExecCommand(ctx, comando, true)
}

func GitNewbranch(ctx context.Context, dir, tag string) {
	comando := "cd " + dir + " && git checkout -b " + tag
	ExecCommand(ctx, comando, true)
}

func GitCheckout(ctx context.Context, dir, tag string) {

	comando := "cd " + dir + " && git checkout " + tag
	ExecCommand(ctx, comando, true)
}

func GitCommit(ctx context.Context, dir, message string) {
	comando := "cd " + dir + " && git commit -m '" + message + "'"
	ExecCommand(ctx, comando, true)
}

func GitPush(ctx context.Context, dir, branch string) {
	log.Println(branch)
	comando := "cd " + dir + "; git push -u origin '" + branch + "'"
	ExecCommand(ctx, comando, true)
}

func GitInitRepo(ctx context.Context, nomeRepo string) {
	comando := " curl -X POST -v -u \"laszlo72:2TvddWPjJaSdJFTqhUdD\" -H \"Content-Type: application/json\" "
	comando += " " + os.Getenv("bitbucketHost") + "/repositories/" + os.Getenv("bitbucketProject") + "/" + nomeRepo
	comando += " -d '{\"scm\": \"git\", \"is_private\": \"true\",\"project\": {\"key\": \"MSF\"}, \"name\":\"" + nomeRepo + "\"}'"
	ExecCommand(ctx, comando, true)
}

func GitCreateNewBranchApi(ctx context.Context, repo, branch string) {
	comando := "curl -X POST -is -u \"laszlo72:2TvddWPjJaSdJFTqhUdD\" -H \"Content-Type: application/json\" "
	comando += " " + os.Getenv("bitbucketHost") + "/repositories/" + os.Getenv("bitbucketProject") + "/" + repo + "/refs/branches "
	comando += " -d '{ \"name\": \"" + branch + "\", \"target\": { \"hash\": \"master\" } }'"
	ExecCommand(ctx, comando, true)
}

func GitInit(ctx context.Context, dir, nomeRepo, GitSrcTipo, Namespace string) {

	// // troppo a majale
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// fmt.Println("sudo chown laszlo:laszlo " + dir + " -R")
	// fmt.Println()
	// fmt.Println("sudo chmod 777 " + dir + " -R")
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()

	// comando := "sudo chown laszlo:laszlo " + dir + " -R"
	// ExecCommand(comando)
	// comando = "sudo chmod 777 " + dir + " -R"
	// ExecCommand(comando)

	// git init
	comando := "cd " + dir + "; git init "
	ExecCommand(ctx, comando, true)

	// scrivo il README.md
	f, err := os.Create(dir + "/README.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("##README##\n")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	comando = "git config --global user.email \"p.punzo@custom.it\""
	ExecCommand(ctx, comando, true)
	comando = "git config --global user.name \"devops-operator\""
	ExecCommand(ctx, comando, true)

	// git add .
	comando = "cd " + dir + "; git add ."
	ExecCommand(ctx, comando, true)

	// git commit .
	comando = "cd " + dir + "; git commit -m 'first commit'"
	ExecCommand(ctx, comando, true)

	// punto a bitbucket
	comando = "cd " + dir + "; git remote add origin https://" + os.Getenv("bitbucketUser") + ":" + os.Getenv("bitbucketToken") + "@bitbucket.org/" + os.Getenv("bitbucketProject") + "/" + nomeRepo
	ExecCommand(ctx, comando, true)

	// // pull senno si incazza al secondo giro
	// comando = "cd " + dir + "; git pull origin master --allow-unrelated-histories"
	// ExecCommand(comando)

	// push in remoto
	comando = "cd " + dir + "; git push -u origin master"
	ExecCommand(ctx, comando, true)
}
