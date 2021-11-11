package run

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"

	"github.com/lil5/maakfile/global"
)

// https://github.com/Katsushun89/scripts

func execute(script string, c chan *exec.Cmd, wg *sync.WaitGroup) {
	cmd := exec.Command("sh", "-c", script)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = io.Writer(os.Stdout)
	cmd.Stderr = io.Writer(os.Stderr)
	cmd.Start()
	c <- cmd

	cmd.Wait()
	wg.Done()
}

func execParallel(scripts []string) {
	/// channel that waits until all scripts are finished
	c := make(chan int, 1)
	/// wait group for parallel execution
	var wg sync.WaitGroup
	/// channel that collects all PIDs when they are created
	/// this is for killing the process later on
	cmd := make(chan *exec.Cmd, len(scripts))
	for _, script := range scripts {
		wg.Add(1)
		go execute(script, cmd, &wg)
	}

	cmds := []*exec.Cmd{}

	for range scripts {
		v := <-cmd
		cmds = append(cmds, v)
	}
	close(cmd)

	// wait for work group to finish then add to c channel
	go func() {
		wg.Wait()
		c <- 1
	}()

	// reads terminal signals like ctrl + c or SIGTERM from e.g. htop
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(s)

	// here process waits for either os.signal or when c is closed
	select {
	// ctrl + c signal recieved
	case sig := <-s:
		fmt.Printf("%ssignal: %s%s\n", global.ColorYellow, sig, global.ColorReset)
		for _, cmd := range cmds {
			cmd.Process.Kill()
		}
	// all scripts finished
	case <-c:
	}
}

func execSequential(scripts []string) {
	// reads terminal signals like ctrl + c or SIGTERM from e.g. htop
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(s)

	for i := range scripts {
		script := scripts[i]
		/// channel that waits until all scripts are finished
		c := make(chan int, 1)
		/// wait group for parallel execution
		var wg sync.WaitGroup
		/// channel that collects all PIDs when they are created
		/// this is for killing the process later on
		cCmd := make(chan *exec.Cmd, 1)

		wg.Add(1)
		go execute(script, cCmd, &wg)

		cmd := <-cCmd
		close(cCmd)

		// wait for work group to finish then add to c channel
		go func() {
			wg.Wait()
			c <- 1
		}()

		select {
		// ctrl + c signal recieved
		case sig := <-s:
			fmt.Printf("%ssignal: %s%s\n", global.ColorYellow, sig, global.ColorReset)
			cmd.Process.Kill()
		// script finished
		case <-c:
		}
		close(c)
	}
}
