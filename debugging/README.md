## Debugging with Delve and eBPF

> https://developers.redhat.com/articles/2023/02/13/how-debugging-go-programs-delve-and-ebpf-faster#how_to_trace_go_programs_with_delve

### Trace Go program with Delve

```
dlv trace foo
```

### Trace Go program with Delve and eBPF

> TODO: this is still not working.

```
dlv trace --ebpf foo

# Get the following errors on macOS
# unable to set tracepoint on function main.foo: "eBPF is not supported"
```

- useing docker on macOS

```
docker build -t dlv-debugging-foo .
docker run -it --rm dlv-debugging-foo
docker rmi dlv-debugging-foo
```
