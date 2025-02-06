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

* ## 5. Helm client chart repositoy add
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

* ## 2. AWS Container github repository for helm chart cloning
  ```bash
  git clone https://github.com/aws-containers/eks-app-mesh-polyglot-demo.git
  ```

* ## 3. Helm chart dir in the files the checking list
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

* ## 4. Application in the helm chart packaging
  ```bash
  helm package helm-chart/
  ```

* ## 5. Aplication helm chart created the checking
  ```bash
  ls -ltr
  ```

* ## 6. Packaged app Helm chart to S3 helm repository push
  ```bash
  helm s3 push ./productcatalog_workshop-1.0.0.tgz productcatalog
  ```

* ## 7. The helm chart to S3 bucket saved checking for push
  ```bash
  aws s3 ls $S3_BUCKET_NAME
  ```

# AWS S3 repository to helm chart search and install
* ## 1. The saved S3 bucket in the helm chart version search
  ```bash
  helm search repo
  ```

* ## 2. Don`t delivery and chart install test
  ```bash
  helm install productcatalog s3://$S3_BUCKET_NAME/productcatalog_workshop-1.0.0.tgz --version 1.0.0 --dry-run --debug
  ```
  > need to cmd add --dry-run, --debug flag

  * productcatalog: 이 설치에 부여하는 릴리스 이름입니다.
  * s3://$S3_BUCKET_NAME/productcatalog_workshop-1.0.0.tgz: S3 버킷 리포지토리에서 차트 아카이브의 위치를 지정합니다.
  * --version 1.0.0: 설치할 차트 버전을 1.0.0으로 고정합니다.
  * --dry-run: Helm에게 차트를 검증하고 실제로 배포하지 않고 리소스 매니페스트를 생성하도록 지시합니다.
  * --debug: 설치 프로세스에 대한 자세한 디버깅 정보를 출력합니다.

* ## 3. Helm chart install
  ```bash
  helm install productcatalog s3://$S3_BUCKET_NAME/productcatalog_workshop-1.0.0.tgz --version 1.0.0
  ```

* ## 4. Helm chart delivery resoucre to list output
  ```bash
  kubectl get pod,svc,deploy -n workshop -o name
  ```

* ## 4.2 Frontend service is shell variable save and terminal output
  ```bash
  FRONTEND_URL=http://$(kubectl get svc -n workshop frontend -o jsonpath='{.status.loadBalancer.ingress[0].hostname}')
  echo "$FRONTEND_URL"
  ```

# Application helm chart edit and redelivery
* ## 1. Yaml file edit
  ```bash
  cd /home/ssm-user/eks-app-mesh-polyglot-demo/workshop/helm-chart
  sed -i "s/replicaCount:.*/replicaCount: 3/g" values.yaml
  ```
  > sed -i is string text change feature

* ## 2. Updated helm value using to appication config
  ```bash
  helm upgrade productcatalog /home/ssm-user/eks-app-mesh-polyglot-demo/workshop/helm-chart/
  ```

* ## 3. New pod stating for checking [with. Rollback run]
  ```bash
  kubectl get pod -n workshop
  ```

* ## 4. Rollback run
  ```bash
  helm rollback productcatalog 1
  ```