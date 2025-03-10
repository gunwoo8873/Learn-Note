# Load Average and System load

# Load Average 계산 과정
```bash
# CMD
uptime

# uptime result
19:02:36 up  5:21,  3 users,  load average: 0.00, 0.00, 0.00

# /proc/loadavg
0.00 0.00 0.00 1/223 3050
```

* ## Load average uptime
  ```bash
  # CMD
  strace -s [PID] -f -t -o uptime_dump uptime

  # ~/uptime_dump
  3082  19:10:01 execve("/bin/uptime", ["uptime"], 0x7fffd08d2bf0 /* 24 vars */) = 0
  3082  19:10:01 brk(NULL)                = 0x55d71b98b000
  3082  19:10:01 access("/etc/ld.so.preload", R_OK) = -1 ENOENT (No such file or directory)
  3082  19:10:01 openat(AT_FDCWD, "/etc/ld.so.cache", O_RDONLY|O_CLOEXEC) = 3
  3082  19:10:01 fstat(3, {st_mode=S_IFREG|0644, st_size=26907, ...}) = 0
  3082  19:10:01 mmap(NULL, 26907, PROT_READ, MAP_PRIVATE, 3, 0) = 0x7fa815db6000
  3082  19:10:01 close(3)                 = 0
  {...}
  ```

* ## vmstat to system load check
  ```bash
  # CMD
  vmstat

  # Interval result
  procs -----------memory---------- ---swap-- -----io---- -system-- -------cpu-------
  r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st gu
  1  0      0 7175400   4224 268808   0    0    13     1  104    0  0  0 100  0  0  0
  0  0      0 7175400   4224 268852   0    0     0     0  274  160  0  0 100  0  0  0
  0  0      0 7175176   4224 268852   0    0     0     0  222  154  0  1 99  0  0  0
  0  0      0 7175176   4224 268852   0    0     0     0  167  113  0  0 100  0  0  0
  0  0      0 7175176   4224 268852   0    0     0     0  115   97  0  0 100  0  0  0
  0  0      0 7175176   4224 268852   0    0     0     0  106   96  0  0 100  0  0  0
  0  0      0 7175176   4224 268852   0    0     0     0  116  107  0  1 99  0  0  0
  0  0      0 7175176   4224 268852   0    0     0     0   91   79  0  0 100  0  0  0
  
  > r / b : r [실행되기를 기다리거나 현재 실행되고 있는 프로세스의 개수], b [I/O를 위해 대기열에 존재하는 프로세스의 개수를 의미]
  > r = nr_running, n = nr_uninterruptible
  ```