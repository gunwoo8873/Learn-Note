#!/usr/env bash

set -e

function Selected_Version()
{
  options=("Ubuntu" "REHL" "Back")
  select OPTIONS in "${options[@]}" do
    case $OPTIONS in
      Ubuntu);;
      REHL);;
      Back);;
    esac
  done
}

function Update_Version_Ubuntu()
{
  source ./bin/ubuntu_update.sh
}

function Update_Version_REHL()
{
  source ./bin/rehl_update.sh
}


function update_menu()
{
  options=("Update" "Exit")
  select OPTIONS in "${options[@]}" do
    case $OPTIONS in
      Update);;
      Exit);;
    esac
  done
}