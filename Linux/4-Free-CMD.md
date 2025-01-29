# Linux system process amount used check

* ## free memory command
  ```md
  # CMD
  free [-Option]

  # Result
                  total       used        free      shared  buff/cache   available
  Mem:         7829468      623876     7181988        9304      274988     7205592
  Swap:        6291452           0     6291452

  > buff / cache : 
  ```

* ## meminfo
  ```bash
  # CMD
  cat /proc/meninfo

  # Resutl
  MemTotal:        7829468 kB
  MemFree:         7181296 kB
  MemAvailable:    7204924 kB
  Buffers:            4224 kB
  Cached:           228240 kB
  SwapCached:            0 kB
  Active:           233652 kB
  Inactive:         117452 kB
  Active(anon):     127856 kB
  Inactive(anon):        0 kB
  Active(file):     105796 kB
  Inactive(file):   117452 kB
  Unevictable:           0 kB
  Mlocked:               0 kB
  SwapTotal:       6291452 kB
  SwapFree:        6291452 kB
  {...}
  ```