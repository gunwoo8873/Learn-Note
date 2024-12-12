## Azure App Service 검사
Azure App Service는 웹 애플리케이션, REST API 및 모바일 백 엔드를 `호스트하는 HTTP 기반 서비스`
`유동적인 프로그래밍 언어 또는 프레임워크로 개발이 가능`하며 `Window or Linux 기반 환경에서 애플리케이션을 쉽게 실행하고 확장`할 수 있다.

## Auto Size Scaling 지원
스케일 업/다운 혹은 아웃/인 기능은 Azure App Service에 포함되어 있다. `스케일 업/다운은 웹앱 사용량에 따라 웹앱을 호스트 중인 머신의 리소스를 늘리거나 줄여 앱을 관리`할 수 있다. 리소스에는 사용 가능한 코어 수 또는 RAM 크기가 포함되어 있기 때문에 `스케일 아웃/인은 웹앱을 실행 중인 머신의 인스턴스 수를 늘리거나 줄이는 기능`이다.

## Container 지원
Azure App Service를 사용하면 `Window와 Linux에서 Container화된 웹앱을 배포하고 실행`할 수 있다.  
Private Azure Container Registry또는 Docker Hub에서 Container Image를 가져올 수 있다.  
Azure App Service는 Container Instance Orchestration을 위해 Container Apps, Window Container 및 Docker Compose도 지원한다.

## CI/CD 지원
Azure Portal은 개발 머신에서 Azure DevOps Service, Github, Bitbucket, FTP 또는 Local Git Repogistory에 대한 연속적인 CI/CD를 기본적으로 제공한다. 웹앱을 위의 원본과 연결하면 APP Service가 코드 및 코드의 이후 변경 내용을 웹앱에 자동 동기화하여 나머지 작업을 수행한다. Azure Container Registry또는 Docker Hub를 사용하여 Container화된 웹앱에 대한 연속 CI/CD도 지원한다.

## CD Slot
웹앱을 배포할 때 표준 App Service Plan Layer 이상에서 실행하면 기본 Production Slot 대신 별도의 CD Slot을 사용할 수 있다. CD Slot은 Self-Host 이름을 가진 Live App이다. 앱 콘텐츠 및 구성 요소는 Production Slot을 포함하여 두 배포 슬롯 간에 교환될 수 있다.