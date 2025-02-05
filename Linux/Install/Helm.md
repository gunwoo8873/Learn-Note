# Helm install and study note

* ## 1. Install CMD
  ```bash
  # 1. Helm install for curl request
  curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
  ```
  ```bash
  # 2. Helm user access and shell run file
  chmod 700 get_helm.sh
  ./get_helm.sh
  ```
  ```bash
  # 3. Helm installed to version checking to output
  helm version --short
  ```

* ## 2. Plugin install
  ```bash
  helm plugin install https://github.com/hypnoglow/helm-s3.git
  ```

* ## 3. Hosting helm repository initalized [with. S3]
  ```bash
  helm s3 init s3://$S3_BUCKET_NAME
  ```

* ## 4. In the saved *.html checking [with. S3]
  ```bash
  aws s3 ls $S3_BUCKET_NAME
  ```

* ## 5. Helm client chert repositoy add
  ```bash
  helm repo add productcatalog s3://$S3_BUCKET_NAME
  ```
  > Tip : Helm repo add cmd is getting to index.html in the /home/ssm-user/.cache/helm/repository

# Helm chart packaging for aws s3 to push
* S3 버킷과 같은 Helm 차트 리포지토리에서 차트를 공유
* 단일 Helm instal 명령으로 패키징된 차트를 설치
* 환경 간에 업데이트 및 롤백을 관리

* ## 1. Helm repository list
  ```bash
  helm repo list
  ```

* ## 2. AWS Container github repository for helm chert cloning
  ```bash
  git clone https://github.com/aws-containers/eks-app-mesh-polyglot-demo.git
  ```

* ## 3. Helm chert dir in the files the checking list
  ```bash
  cd eks-app-mesh-polyglot-demo/workshop
  tree helm-chart/
  ```

  > * Chart.yaml: 차트에 대한 정보가 들어 있는 YAML 파일
  > * productcatalog_workshop-1.0.0.tgz: Kubernetes 클러스터에 적용할 수 있는 버전이 지정된 차트 아카이브
  > * security/: Pod 보안 그룹에 대한 매니페스트가 포함된 디렉터리
  > * templates/: 값과 결합되면 유효한 Kubernetes 매니페스트 파일을 생성하는 템플릿의 디렉터리
  > * templates/NOTES.txt: 선택 사항: 간단한 사용 노트가 포함된 일반 텍스트 파일
  > * values.yaml: 기본 구성 값을 정의하는 일련의 YAML 파일

* ## 4. Application in the helm chert packaging
  ```bash
  helm package helm-chart/
  ```

* ## 5. Aplication helm chert created the checking
  ```bash
  ls -ltr
  ```

* ## 6. Packaged