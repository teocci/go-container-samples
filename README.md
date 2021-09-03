## go-container-samples ![Go Reference][1]

`go-csv-samples` is an open-source implementation of a basic container build from scratch in Go solely for learning
purpose. It uses namespaces and cgroups, mount a tmpfs that's isolated from host filesystem.

### What it does

> **Note:** Before starting, you will have to extract the content of `ubuntu-fs.tar.gz` file into `./os-root-fs/ubuntu` which will be mounted as the container's root directory.

> **Note:** You will have to be inside a Linux box (Ubuntu in my case) to try this.

The example `5-isolate-cgroup-sample` is the full implementation, and it needs root privilege for creating `cgroup`:

### Usage
```bash
sudo su
go run main.go run /bin/bash
``` 

It will:
- Fork itself with `CLONE_NEWUTS`, `CLONE_NEWPID`, `CLONE_NEWNS` flags with isolated hostname, processes and mounts
- The forked process will create `cgroup` to limit memory usage of itself and any child process it creates
- Mount `./ubuntu` directory as root filesystem using `chroot` to limit access to host machine's filesystem
- Mount `/newtemp` directory as tmpfs. Any change made to this directory will not be visible from host.
- Mount proc (where `CLONE_NEWPID` namespace was already set) so that container can run `ps` and see only the processes
  running inside it.
- Execute the supplied argument `/bin/bash` inside the isolated environment

---

## Sources of the inspiration and information

* [Building Containers from Scratch][4] with Go by Liz Rice
If you don't have access to [Building Containers from Scratch][4], Liz Rice gave several talks on the same topic in
other conferences. One of them is ["GOTO 2018 • Containers From Scratch"][5].
* [Building a container from scratch in Go][8]
* [Containers from scratch: The sequel][9]
* [Build Your Own Container Using Less than 100 Lines of Go][10]
* [Understand Container][6] by Bin Chen
---

## Bonus tip: Setting up VS Code for cross-platform development

I have used Windows to develop the container from scratch and have run it inside Ubuntu in a virtualbox by sharing the
development directory. While this setup is fine for running the code inside Linux, the development experience is not
great because a lot of pieces of this application is Linux specific. For example, calls like `syscall.Sethostname` or
the `Cloneflags` field in the `syscall.SysProcAttr{}` struct is not available in intellisense in **VSCode** when the dev
environment is not Linux. **VSCode** will mark those lines as errors, because they are platform specific and declared in
the standard library in Go for Linux only.

Fortunately, there is a workaround, and it is very simple. Search for `"go.toolsEnvVars"` in **VSCode** settings, copy it to User Settings and change it to:

```json
"go.toolsEnvVars": {
    "GOOS": "linux"
}
```
Now restart **VSCode** and after that it will recognize all Linux specific declarations and will not see them as errors.
`Go-to-definition` will work properly too.
---

> **PS**: the contents of `ubuntu_fs.tar.gz` file has been extracted from **Ubuntu docker image** by using `docker export...` command.

[1]: https://pkg.go.dev/badge/github.com/teocci/go-container-samples.svg
[2]: https://pkg.go.dev/github.com/teocci/go-container-samples
[3]: https://github.com/teocci/go-container-samples/releases/latest
[4]: https://www.safaribooksonline.com/videos/building-containers-from/9781491988404
[5]: https://www.youtube.com/watch?v=8fi7uSYlOdc
[6]: https://pierrchen.blogspot.com/2018/08/understand-container-index.html
[7]: https://github.com/lizrice/containers-from-scratch
[8]: https://www.youtube.com/watch?v=Utf-A4rODH8
[9]: https://www.youtube.com/watch?v=_TsSmSu57Zo
[10]: https://www.infoq.com/articles/build-a-container-golang/