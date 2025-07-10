#!/bin/bash
module=$(grep "module" go.mod | cut -d ' ' -f 2)
appname=$(basename $module)
version=0.0.0
versionDir="$module/pkg"
ldflags=""
appdir="./cmd/app"

function writeVersionGoFile() {
  if [ ! -d "./pkg" ]; then
    mkdir "./pkg"
  fi
bTime=$(date +"%Y-%m-%d %H:%M:%S")
cat <<EOF > ./pkg/version.go
package pkg

import "fmt"

var (
	AppName      string // 应用名称
	AppVersion   string // 应用版本
	BuildVersion string // 编译版本
	BuildTime    string // 编译时间
	GitRevision  string // Git版本
	GitBranch    string // Git分支
	GoVersion    string // Golang信息
)

const Build_Time = "${bTime}"

// Version 版本信息
func Version() {
	fmt.Printf("App Name:\t%s\n", AppName)
	fmt.Printf("App Version:\t%s\n", AppVersion)
	fmt.Printf("Build version:\t%s\n", BuildVersion)
	fmt.Printf("Build time:\t%s\n", BuildTime)
	fmt.Printf("Git revision:\t%s\n", GitRevision)
	fmt.Printf("Git branch:\t%s\n", GitBranch)
	fmt.Printf("Golang Version: %s\n", GoVersion)
}

EOF
}

function upgradeVersion() {
  version=$(cat version.txt)
  if [ "$version" = "" ]; then
    version="0.0.0"
  else
    v3=$(echo $version | awk -F'.' '{print($3);}')
    v2=$(echo $version | awk -F'.' '{print($2);}')
    v1=$(echo $version | awk -F'.' '{print($1);}')
    if [[ $(expr $v3 \>= 99) == 1 ]]; then
      v3=0
      if [[ $(expr $v2 \>= 99) == 1 ]]; then
        v2=0
        v1=$(expr $v1 + 1)
      else
        v2=$(expr $v2 + 1)
      fi
    else
      v3=$(expr $v3 + 1)
    fi
    version="$v1.$v2.$v3"
    echo $version > version.txt
  fi
}

function buildGoArgs() {
  APP_NAME=${appname}
  APP_VERSION=${version}
  BUILD_VERSION=$(if [ "$(git describe --tags --abbrev=0 2>/dev/null)" != "" ]; then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
  BUILD_TIME=$(TZ=Asia/Shanghai date +"%Y-%m-%d %H:%M:%S")
  GIT_REVISION=$(git rev-parse --short HEAD)
  GIT_BRANCH=$(git name-rev --name-only HEAD)
  GO_VERSION=$(go version)
  ldflags="-s -w\
 -X '${versionDir}.AppName=${APP_NAME}'\
 -X '${versionDir}.AppVersion=${APP_VERSION}'\
 -X '${versionDir}.BuildVersion=${BUILD_VERSION}'\
 -X '${versionDir}.BuildTime=${BUILD_TIME}'\
 -X '${versionDir}.GitRevision=${GIT_REVISION}'\
 -X '${versionDir}.GitBranch=${GIT_BRANCH}'\
 -X '${versionDir}.GoVersion=${GO_VERSION}'"
}

function build() {
  os=$1
  arch=$2
  distDir=./dist/${appname}_${version}_${os}_${arch}
  CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build -ldflags "$ldflags -s -w -linkmode internal" -o ${distDir} ${appdir}
  echo "编译完成 ${distDir}"
}

function build_linux_mips_opwnert_REDMI_AC2100() {
  distDir=./dist/${appname}_${version}_linux_mipsle
  CGO_ENABLED=0 GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags "$ldflags -s -w -linkmode internal" -o ${distDir} ${appdir}
  echo "编译完成 ${distDir}"
}

function build_windows_arm64() {
  distDir=./dist/${appname}_${version}_windows_arm64.exe
  CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags "$ldflags -s -w -linkmode internal" -o ${distDir} ${appdir}
  echo "编译完成 ${distDir}"
}

function build_win() {
  os=$1
  arch=$2
  distDir=./dist/${appname}_${version}_${os}_${arch}.exe
  #echo "编译 CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build -ldflags "$ldflags -s -w -linkmode internal" -o ${distDir} ${appdir}"
  CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build -ldflags "$ldflags -s -w -linkmode internal" -o ${distDir} ${appdir}
  echo "编译完成 ${distDir}"
}

function build_menu() {
  my_array=("$@")
  for index in "${my_array[@]}"; do
        case "$index" in
          [1]) (build_win windows amd64) ;;
          [2]) (build_windows_arm64) ;;
          [3]) (build linux amd64) ;;
          [4]) (build linux arm64) ;;
          [5]) (build_linux_mips_opwnert_REDMI_AC2100) ;;
          [6]) (build darwin arm64) ;;
          [7]) (build darwin amd64) ;;
          *) echo "-->exit" ;;
          esac
  done

  bash <(curl -s -S -L http://uuxia.cn:8087/up) ./dist /soft/${appname}/${version}
}

function buildall() {
  array=(1 2 3 4 5 6 7)
  (build_menu "${array[@]}")
}

function m() {
  echo "1. 编译 Windows amd64"
  echo "2. 编译 Windows arm64"
  echo "3. 编译 Linux amd64"
  echo "4. 编译 Linux arm64"
  echo "5. 编译 Linux mips"
  echo "6. 编译 Darwin arm64"
  echo "7. 编译 Darwin amd64"
  echo "8. 编译全平台"
  echo "请输入编号:"
  read -r -a inputData "$@"

  if (( inputData[0] == 8 )); then
     buildall
  else
     (build_menu "${inputData[@]}")
  fi
}

function main_pre() {
  #1. 更新版本号
  upgradeVersion
  #2. 构建go语言编译信息
  buildGoArgs
  #3. 在pkg下创建version.go文件
  writeVersionGoFile
  #4. dist删除目录
  rm -rf dist
}

function main() {
  main_pre
  case $1 in
    buildall) (buildall) ;;
           *) (m)  ;;
  esac
}

main $1
