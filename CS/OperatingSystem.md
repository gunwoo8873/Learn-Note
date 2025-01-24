# 운영체제와 커널 [OperatingSystem]
### 운영체제
```md
운영체제는 응용 프로그램(Application Programs, App)과 하드웨어(Hardware, H/W) 사이에 위치하며, 사용자가 응용 프로그램을 통해서
어떤 서비스를 요청하면 응용 프로그램은 운영체제(OS)에게 신호를 전달하고 하드웨어를 조작을 요청한다. 요청 받은 운영체제는 서비스에
맞는 하드웨어를 작동시켜 수행을 완수한다.
```

### 운영체제 정의
```md
* 컴퓨터 위에서 항상 실행되는 하나의 프로그램
* Kernel로도 정의
* H/W를 제어하는 소프트웨어(S/W)
* 응용 프로그램과 하드웨어 사이의 중계자
```

<figure align="center">
  <img src="./img/os-n-kernel-fig1.jpeg" alt="">
  <figcaption align="center">OS Kernel</figcaption>
</figure>

### 운영체제의 문제
```md
메모리가 보호가 되지 않을 경우 세그멘테이션 오류(Segmentation fault), 메모리 침범 등 다양한 문제가 발생할 수 있으며, 
읽기(Read)전용 메모리 영역에 쓰기(Write)를 시도한 경우 H/W에 의해 인지되면 세그멘테이션 오류가 발생하지만 인지가 안됐을 때는 
오염이 발생한채로 프로그램이 계속 수행되어 더욱 심한 문제를 야기할 수 있다.
가상 주소(Virtual address)를 채용한 OS에서는 프로세스 별로 페이지 테이블을 관리하여 원칙적으로는 다른 프로세스의 메모리 영역에
접근할 수 없으나, 취약점이 노출될 경우 임의로 메모리를 갈취하거나 수정할 수도 있게 된다. 결국 보안을 위해 시스템 자체에 영향을
줄 수 있는 작업들은 커널이라는 분리된 공간에서 특별한 권한을 가지고 실핼하도록 설계된다.
```

### Kernel
```md
커널 설계는 OS에서 차지하는 커널의 비중이나 추상화 방법에 따라 분류하여 다음과 같이 부른다.
* 단일형 커널 (Monolithic Kernel)
* 마이크로 커널 (Micro Kernel)
* 하이브리드 커널 (Hybrid Kernel)
* 엑소 커널 (Exo Kernel)
* ...ETC
```

#### Monolithic Kernel
```md
커널에 운영체제의 모든 기능(프로세스 관리, 메모리 관리, 디바이스 드라이버 등)이 포함

# 장점
시스템 호출과 커널 간 통신이 빠르고 드라이버와 서비스가 커널 내에 밀접하게 연결되어 효율적이다
# 단점
커널 크기가 크고 복잡하며 에러 발생시 시스템 전체가 영향을 받을 가능성이 크다
```

#### Micro Kernel
```md
커널의 최소 기능만 포함하고 나머지 운영체제 서비스(파일 시스템, 드라이버 등)는 사용자 공간에서 실행

# 장점
한 모듈의 오류가 전체 시스템에 영향을 미치지 않고 확장성과 유지보수성이 뛰어나다
# 단점
사용자 공간과 커널 간의 빈번한 메시지 전달로 성능에 대한 오버헤드가 발생
```

#### Hybrid Kernel
```md
단일형 커널과 마이크로 커널의 장점을 결합. 핵심 기능은 커널에, 나머지는 모듈로 분리

# 장점
단일형 커널의 성능과 마이크로 커널의 안정성을 동시에 추구하며, 모듈화된 설계로 확장성과 관리가 용이하다
# 단점
설계가 복잡하고 구현이 어려움
```

#### Exo Kernel
```md
하드웨어 리소스를 최대한 직접적으로 사용자 애플리케이션에 노출

# 장점
애플리케이션이 하드웨어를 직접 제어하므로 오버헤드가 거의 없고 특정 애플리케이션에 최적화된 자원 관리 가능
# 단점
복잡한 프로그래밍과 범용성 부족
```

# 컴퓨터 시스템의 구성 [Computer-System Organization]

### Bootstrap
```md
컴퓨터 전력이 들어자마자 최우선 순위로 수행되는 프로그램이다. 디스크에 저장되어 있는 Bootstrap loader를 CPU가 실행하여 
디스크에 저장되어 있는 OS를 메모리에 복사하여 올린다. 즉 Bootstrap 프로그램을 정의하면 OS를 메모리에 적재(load)하는 프로그램이다.
```

### Interrupt
```md
H/W가 자동중에 CPU에게 알려주는 신호(Signal)이다. 시스템 버스(System bus)를 통해 CPU에게 신호를 전송함 으로써
어느 시간이던 H/W는 인터럽트를 발생시킬 수 있다.
```

<figure align="center">
  <img src="./img/Interrupt.png" alt="" width=600>
  <figcaption align="center">Interrupt Signal</figcaption>
</figure>

### Von neumann architecture
```md
폰 노이만 구조(Von neumann architecture)는 폰노이만이 제시한 프로그램 내장 방식이다. 다음과 같이 내장 방식이다.
* 명령어 실행 사이클(Instruction execution cycle)은 첫번째로 메모리에서 명령어를 인출(fetches)하고 명령어
  레지스터(Instruction Register)에 명령어를 저장함
* 
```