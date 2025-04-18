# Network

## 네트워크 자동화

* 특정 사용 사례나 기술에 대한 것이 아닌 사용자 환경과 비즈니스 이익에 도움
* 네트워크 구성을 조정하거나, 트래픽 엔지니어링 정책 적용
* **설계 제약[design constraint]**을 강제하는 등 운영 이벤트에 반응하는 작업 포함

> 이 모든 활동을 통합하는 한가지 공통적인 특징은 원하는 동적이 예상한 결과로 이어지는 구체적인 단계를 순서대로 설명할 수 있다는 능력이며,  
> 다음과 같이 5가지 단계로 구성되어 네트워크 설계의 구성 템플릿 개발 및 테스트를 자동화할 수 있다.
> 
> * 정적 구성 분석[Static configuration analysis]
> * 수학적 네트워크 모델링[Mathhematical network modeling]
> * 테스트 주도 개발[TDD : Test Driven Development]
> * 소프트웨어 세계의 개념을 도입

## 네트워크 자동화의 필요성

### 상향식 관점
|                     |                                                                                                                           |
| ------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| 일관성[Consistency] | 컴퓨터가 이런 변경을 수행하므로 매번 같은 결과를 기대할 수 있으며, 같은 구성이나 템플릿 또는 정책을 여러 요소에 강제 적용 |
| 신뢰성[Reliability] | 지침[Instruction]은 컴퓨터가 명확하게 해석하는 코드고 입력이나 결과를 검증하는 자동 검사도 추가                           |
| 가시성[Visibility]  | 네트워크의 모든 미래와 과거 변경 사항을 팀의 모든 구성원이 볼 수 있으므로 검토를 통해 결함을 쉽게 해결                    |
| 보편성[Ubiquity]    | 같은 도구를 사용하므로 상호작용이 간단해지며 지식 공유에 원할함                                                           |

### 하향식 관점
|                                     |                                                                                                                                                                                               |
| ----------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 비용 관리[Cost management]          | 리소스 최적화[Resource optimization]를 통해 비용을 절감하고 결함 해결 감소와 감사 보고서 작성을 작성하지 않음으로 네트워크 운영 비용 감소                                                     |
| 서비스 제공 속도[Speed of delivery] | 네트워크 변경 사항을 구성하고 검증하는 속도가 빨라져 고객 서비스를 더욱 빠르게 제공하거나 고객의 요구에 따라 맞춤 서비스를 제공할 수 있다                                                     |
| 위험 관리[Risk management]          | 모든 운영에 보안 정책[Security policy]을 일관되게 강제해 위험을 줄이고 서비스에 영향을 미치는 이슈의 수를 줄여 수익을 개선                                                                    |
| 비즈니스 역량[Business caplability] | 조직이 가치를 정의하는 방법에 따라 기회를 발견하는 데 도움을 줄 수 있으며 가시성이 좋아지면 용량 계획[Capacity planning]을 개선하고 사용하지 않은 용량이나 하스폿[Hostspot]을 발견하는데 도움 |