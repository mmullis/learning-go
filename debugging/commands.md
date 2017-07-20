References
==========
https://dave.cheney.net/tag/gomaxprocs
https://godoc.org/runtime#hdr-Environment_Variables


GOGC=off will disable garbage collection entirely.


Not sure what versions these all work with

GOTRACEBACK=0 will suppress all tracebacks, you only get the panic message.
GOTRACEBACK=1 is the default behaviour, stack traces for all goroutines are shown, but stack frames related to the runtime are suppressed.
GOTRACEBACK=2 is the same as the previous value, but frames related to the runtime are also shown, this will reveal goroutines started by the runtime itself.
GOTRACEBACK=crash is the same as the previous value, but rather than calling os.Exit, the runtime will cause the process to segfault, triggering a core dump if permitted by the operating system.

GOTRACEBACK=none will suppress all tracebacks, you only get the panic message.
GOTRACEBACK=single is the new default behaviour that prints only the goroutine believed to have caused the panic.
GOTRACEBACK=all causes stack traces for all goroutines to be shown, but stack frames related to the runtime are suppressed.
GOTRACEBACK=system is the same as the previous value, but frames related to the runtime are also shown, this will reveal goroutines started by the runtime itself.
GOTRACEBACK=crash is unchanged from Go 1.5.


GODEBUG=gctrace=1,schedtrace=1000 godoc -http=:8080
GODEBUG=gctrace=1 godoc -http=:8080 -index
// gctrace=1 is the output of the heap scavenger.

GODEBUG=schedtrace=1000 godoc -http=:8080 -index

GODEBUG=scheddetail=1,schedtrace=1000 godoc -http=:8080 -index


