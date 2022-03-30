// Note: cgo can't handle some Darwin/ARM structures, so this file can't
// be auto generated by cgo yet.
// Created based on output of `cgo -cdefs defs_darwin.go` and Darwin/ARM
// specific header (mainly mcontext and ucontext related stuff)

package runtime

import "unsafe"

const (
	_EINTR     = 0x4
	_EFAULT    = 0xe
	_ETIMEDOUT = 0x3c

	_PROT_NONE  = 0x0
	_PROT_READ  = 0x1
	_PROT_WRITE = 0x2
	_PROT_EXEC  = 0x4

	_MAP_ANON    = 0x1000
	_MAP_PRIVATE = 0x2
	_MAP_FIXED   = 0x10

	_MADV_DONTNEED = 0x4
	_MADV_FREE     = 0x5

	_SA_SIGINFO   = 0x40
	_SA_RESTART   = 0x2
	_SA_ONSTACK   = 0x1
	_SA_USERTRAMP = 0x100
	_SA_64REGSET  = 0x200

	_SIGHUP    = 0x1
	_SIGINT    = 0x2
	_SIGQUIT   = 0x3
	_SIGILL    = 0x4
	_SIGTRAP   = 0x5
	_SIGABRT   = 0x6
	_SIGEMT    = 0x7
	_SIGFPE    = 0x8
	_SIGKILL   = 0x9
	_SIGBUS    = 0xa
	_SIGSEGV   = 0xb
	_SIGSYS    = 0xc
	_SIGPIPE   = 0xd
	_SIGALRM   = 0xe
	_SIGTERM   = 0xf
	_SIGURG    = 0x10
	_SIGSTOP   = 0x11
	_SIGTSTP   = 0x12
	_SIGCONT   = 0x13
	_SIGCHLD   = 0x14
	_SIGTTIN   = 0x15
	_SIGTTOU   = 0x16
	_SIGIO     = 0x17
	_SIGXCPU   = 0x18
	_SIGXFSZ   = 0x19
	_SIGVTALRM = 0x1a
	_SIGPROF   = 0x1b
	_SIGWINCH  = 0x1c
	_SIGINFO   = 0x1d
	_SIGUSR1   = 0x1e
	_SIGUSR2   = 0x1f

	_FPE_INTDIV = 0x7
	_FPE_INTOVF = 0x8
	_FPE_FLTDIV = 0x1
	_FPE_FLTOVF = 0x2
	_FPE_FLTUND = 0x3
	_FPE_FLTRES = 0x4
	_FPE_FLTINV = 0x5
	_FPE_FLTSUB = 0x6

	_BUS_ADRALN = 0x1
	_BUS_ADRERR = 0x2
	_BUS_OBJERR = 0x3

	_SEGV_MAPERR = 0x1
	_SEGV_ACCERR = 0x2

	_ITIMER_REAL    = 0x0
	_ITIMER_VIRTUAL = 0x1
	_ITIMER_PROF    = 0x2

	_EV_ADD       = 0x1
	_EV_DELETE    = 0x2
	_EV_CLEAR     = 0x20
	_EV_RECEIPT   = 0x40
	_EV_ERROR     = 0x4000
	_EV_EOF       = 0x8000
	_EVFILT_READ  = -0x1
	_EVFILT_WRITE = -0x2

	_PTHREAD_CREATE_DETACHED = 0x2

	_F_SETFD    = 0x2
	_FD_CLOEXEC = 0x1
)

type stackt struct {
	ss_sp    *byte
	ss_size  uintptr
	ss_flags int32
}

type sigactiont struct {
	__sigaction_u [4]byte
	sa_tramp      unsafe.Pointer
	sa_mask       uint32
	sa_flags      int32
}

type usigactiont struct {
	__sigaction_u [4]byte
	sa_mask       uint32
	sa_flags      int32
}

type siginfo struct {
	si_signo  int32
	si_errno  int32
	si_code   int32
	si_pid    int32
	si_uid    uint32
	si_status int32
	si_addr   uint32
	si_value  [4]byte
	si_band   int32
	__pad     [7]uint32
}

type timeval struct {
	tv_sec  int32
	tv_usec int32
}

func (tv *timeval) set_usec(x int32) {
	tv.tv_usec = x
}

type itimerval struct {
	it_interval timeval
	it_value    timeval
}

type timespec struct {
	tv_sec  int32
	tv_nsec int32
}

//go:nosplit
func (t *timespec) set_nsec(ns int64) {
	t.tv_sec = int32(ns / 1000000000)
	t.tv_nsec = int32(ns % 1000000000)
}

type floatstate32 struct {
	r     [32]uint32
	fpscr uint32
}

type regs32 struct {
	r    [13]uint32 // r0 to r12
	sp   uint32     // r13
	lr   uint32     // r14
	pc   uint32     // r15
	cpsr uint32
}

type exceptionstate32 struct {
	trapno     uint32 // NOTE: on 386, the trapno field is split into trapno and cpu
	err        uint32
	faultvaddr uint32
}

type mcontext32 struct {
	es exceptionstate32
	ss regs32
	fs floatstate32
}

type ucontext struct {
	uc_onstack  int32
	uc_sigmask  uint32
	uc_stack    stackt
	uc_link     *ucontext
	uc_mcsize   uint32
	uc_mcontext *mcontext32
}

type keventt struct {
	ident  uint32
	filter int16
	flags  uint16
	fflags uint32
	data   int32
	udata  *byte
}

type pthread uintptr
type pthreadattr struct {
	X__sig    int32
	X__opaque [36]int8
}
type pthreadmutex struct {
	X__sig    int32
	X__opaque [40]int8
}
type pthreadmutexattr struct {
	X__sig    int32
	X__opaque [8]int8
}
type pthreadcond struct {
	X__sig    int32
	X__opaque [24]int8
}
type pthreadcondattr struct {
	X__sig    int32
	X__opaque [4]int8
}

type machTimebaseInfo struct {
	numer uint32
	denom uint32
}