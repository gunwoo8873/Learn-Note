# AWS 글로벌 구조
> [AWS Global Infrastructure URL](https://infrastructure.aws/)

## AWS Regions [지역]

  - **1. Compilance**
  - **2. Proximity**
  - **3. Available Service**
  - **4. Pricing**

## AWS Availability Zones / AZ [가용성 구역]
* AZ는 지역 안의 존재하는 물리적 데이터 센터 그룹으로서 그리고 지역의 태그 형태로 구성 되어 있다.
* AZ 안에 존재하는 a, b, c 의 물리적 위치는 다르다

```md
  Region[지역] - Direction[방향] - Availability Zones[구역]

  * United States [EX : us-east-2a]
  * Asia [EX : ap-south-1a]
  * Canada [EX : ca-central-1a]
  * Europe [EX : eu-west-1a]
  * South America [EX : sa-east-1a]
```

## AWS Data Centers [데이터 센터]
## AWS Edge Locations and Points of Presence [에지 위치와 존재 지점]



# AWS Console
* AWS Global Service
  - Identity Access Menagement [IAM]
  - Route53 [DNS Service]
  - CloudFront [Content Delivery Network]
  - WAF [Web Application]

* AWS Service are Region-Scoped
  - Amazone EC2 [Infrastructure]
  - Elastic Beanstalk [Platform]
  - Lambda [Function]
  - Rekognition [S/W]


# AWS Identity and Access Management [IAM]
* IAM : 정체성과 접근 관리 및 글로벌 서비스에 해당
* Root Account : 기본적으로 생성하여 사용하고 그리고 공유
* Users : 조직내 사용자들을 하나의 그룹화
* Groups : 사용자를 포함하고 다른 그룹에 속하지 않도록 조정

## IAM Permissions
* User와 Group에서의 정책으로 JSON 파일로 작성된 문서를 할당 받을 수 있다. 사용자의 권한을 정의하며 AWS에서는 최소 권한 원칙을 적용하여 필요로 하는 권한에서 보다 많은 권한을 부여하지 않도록 하는 방지기능

