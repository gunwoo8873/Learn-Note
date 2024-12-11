#!/bin/bash

# 설정
total=20
progress=0

while [ $progress -le $total ]; do
  # 퍼센트 계산
  percent=$(( progress * 100 / total ))
  
  # 진행바 생성
  bar=$(printf "%0.s=" $(seq 1 $progress))
  spaces=$(printf "%0.s-" $(seq 1 $((total - progress))))
  
  # 출력
  printf "\r%3d%% [%s%s]" "$percent" "$bar" "$spaces"
  
  # 진행
  progress=$((progress + 1))
  sleep 0.1
done

echo -e "\n완료!"
