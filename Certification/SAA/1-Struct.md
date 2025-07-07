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

## IAM Policies Structure
* Consists
  - Version : 정책 언어 버전이며 항상 사용되는 데이터는 2012. 10. 17로 기본값으로 설정되어 있다.
  - Id : 정책 고유 식별자
  - Statement : 한개 이상의 개별 진술[?]

* Statements Consists
  - Sid : 진술 식별자
  - Effect : 접근 허용 여부
  - Principal : 정책이 적용되는 계정 혹은 사용자나 역할[?]
  - Action : 접근 허용 권한 범위
  - Resource : 적용된 주소
  - Condition : 정책이 적용되는 조건[?]

* Example S3 Bucket Policy
  ```json
  {
    "Version": "2012-10-17",
    "Id": "S3-Learn",
    "Statement": [
      {
        "Sid": "VisualEditor0",
        "Effect": "Allow",
        "Action": [
            "s3:CreateBucket",
            "s3:ListBucket",
            "s3:DeleteBucket"
        ],
      "Resource": "arn:aws:s3:::*"
      },
      {
        "Sid": "VisualEditor1",
        "Effect": "Allow",
        "Action": [
            "s3:GetObject",
            "s3:DeleteObject"
        ],
        "Resource": "arn:aws:s3:::*/*"
      }
    ]
  }
  ```