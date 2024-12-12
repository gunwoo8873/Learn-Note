# Software Development Lifecycle Model
## SDLC [S/W Development Life Cycle]
```md
개발 팀이 고품질 S/W를 구축하는 데 사용하는 비용 효율적이고 시간 효율적인 Process이다.
SDLC의 목표는 사전 계획을 통해 프로젝트 위험을 최소화함으로써 S/W가 생산 중일 때 글리고 그 이후에도 고객의 기대치를 충촉하도록 하는 것이며, S/W 개발 플로세스를 할당, 완료 및 측정할 수 있는 작업으로 나누는 일련의 단계를 개략적으로 설명한다.
```

```md
#SDLC의 이점은 대표적으로 4 가지를 가져온다.
1. 개발 프로세스에 대한 모든 관련 이해 관계자의 가시성 향상
2. 효율적인 추정, 계획 및 스케줄링
3. 리스크 관리 및 비용 추정 개선
4. 체계적인 S/W 제공 및 고객 만족도 제고
```

### Sequential Development
```md
Sequential Development는 SDLC의 가장 기본적인 접근 방식으로, 순차적으로 작업을 진행하는 방법이다.
```

```md
1. Analysis
2. Design
3. Implementation
4. Testing
```

### Waterfall Model
```md
Sequential Development의 대표적인 사례로, 단계별로 작업을 수행하고 이전 단계로 되돌아가지 않는 순차적 모델이다.
SDLC에서 초기 모델 중 하나로, 각 단계들의 요구사항이 명확하여 정의되고 완료된 후 다음 단계로 넘어가는 특징을 가지고 있지만 변경사항에 대응이 어렵다.
```

```md
1. Requirements
2. Design
3. Development
4. Testing
5. Deployment
6. Maintenance
```

### V-Model
```md
1. Coding
2. Detailed Design      <-> Unit Testing
3. Architectural Design <-> Integration Testing
4. Specification        <-> System Testing
5. Requirements         <-> Accptance Testing
```

### Agile Software Development [ManiFast]
```md
Plan -> Design -> Develop -> Test -> Deploy -> Review -> Launch
```

### Iterative
```md
Iterative는 전체 프로젝트를 한 번에 끝내는 것이 아닌 반복적으로 개선하며 점진적으로 품질 퀄리티를 높이는 방식이다.  
초기에 전체 요구사항을 완전히 정의하지 않는 상태에서 진행이 가능하며, 반복 과정에서 설계, 구현, 테스트 등 지속적으로 수행한다.  
빠른 초기 버전을 만들어 피드백을 수용하는 과정을 통해 수정과 개선은 반복하기 때문에 점차적을로 완성에 가까워진다.
```

### Incremental
```md
Incremental는 작은 조각들을 완성해 나가면서 점전적으로 전체 제품이나 서비스를 구축하는 방식이다.  
초기부터 전체 시스템을 개발하는 것이 아닌, 작은 기능 단위로 나누어 개발하고, 이를 단계별로 확장한다.  
전체 제품이나 서비스가 완성되지 않더라도 일부 기능 사용이 가능하다.
```

> Interative의 이점
> 1. 신속한 피드백
> 2. 리스크 감소
> 3. 유연성
> 4. 점진적 가치 제공

### Agile Scrum Process [2 - 4 Week Sprint]
```md
Vision -> User Stories -> Sprint Backlog -> Daily Stand up -> Scrum Master -> Exit
                                            |- 24 Hours -|
```