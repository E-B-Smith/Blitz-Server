
# Build the server

SHELL       := /bin/bash
makepath    := $(shell pwd)
buildDate   := $(shell date | sed s.[[:space:]].-.g)
buildVersion = 0
buildLinux   = 0
buildDarwin  = 0
rm          := rm -Rf
.SILENT:

export userhost:=blitzhere@shimmering.blue
export installpath:=/home/blitzhere
export GOPATH:=$(makepath)
export GOROOT:=
export PATH:=$(makepath)/bin:$(PATH)


# Compile --


gobuild= \
    cd $1; \
    echo ">>> Building `pwd`."; \
    $(eval goflags= -v -ldflags '-X main.globalCompileTime=$(buildDate) -X main.globalVersion=$(buildVersion)') \
    env GOOS=darwin GOARCH=amd64 go install $(goflags) ; \
    if [[ $$? != 0 ]]; then exit 1; fi; \
    if [[ $(buildLinux) == 1 ]]; then env GOOS=linux  GOARCH=amd64 go install $(goflags); fi; \
    if [[ $$? != 0 ]]; then exit 1; fi; \
    cd - >/dev/null; \
    obj=$$(basename $1); \
    if [[ "$2" == "" ]]; then bob=$$(basename $1); else bob=$$(basename $2); fi; \
    cp -av bin/$$obj Staging/Versions/$$bob.Darwin; \
    if [[ $(buildLinux) == 1 ]]; then cp -av bin/linux_amd64/$$obj Staging/Versions/$$bob.Linux; fi; \
    if [[ $$? != 0 ]]; then exit 1; fi;


compile: FORCE \
    updateversion \
    src/BlitzMessage/Friends.pb.go \
    src/BlitzMessage/UserMessages.pb.go \
    src/BlitzMessage/Types.pb.go \
    src/BlitzMessage/Device.pb.go \
    src/BlitzMessage/Search.pb.go \
    src/BlitzMessage/UserProfiles.pb.go \
    src/BlitzMessage/Feed.pb.go \
    src/BlitzMessage/EntityTags.pb.go \
    src/BlitzMessage/Server.pb.go \
    src/BlitzMessage/UserEvents.pb.go \
    src/ApplePushService/ResourceData.go \
    ; \
        echo ">>> Build version $(buildVersion) $(buildDate)."; \
        $(call gobuild, src/BlitzHere-Server) \
        $(call gobuild, src/Signup-Server) \
        $(call gobuild, src/Status-Server)


updateversion: \
    ; \
    $(eval buildVersion=$(shell Staging/fetch-version -i blitzhere)) \
    if [[ $$? != 0 || "$(buildVersion)" == "" ]]; then exit 1; fi; \
    echo ">>> Updated version to $(buildVersion)."


linux: \
    ;  \
    $(eval buildLinux=1) \
    echo ">>> Building for Linux.";


src/ApplePushService/ResourceData.go : \
    $(shell find src/ApplePushService/Resources) \
    ; \
        echo ">>> Building ApplePush resources"; \
        cd src/ApplePushService; \
        ../Resource/go-makeresource Resources/*


FORCE:


proto \
src/BlitzMessage/Friends.pb.go \
src/BlitzMessage/UserMessages.pb.go \
src/BlitzMessage/Types.pb.go \
src/BlitzMessage/Device.pb.go \
src/BlitzMessage/Search.pb.go \
src/BlitzMessage/UserProfiles.pb.go \
src/BlitzMessage/Feed.pb.go \
src/BlitzMessage/EntityTags.pb.go \
src/BlitzMessage/Server.pb.go \
src/BlitzMessage/UserEvents.pb.go : \
    Protobuf/Source/Friends.proto \
    Protobuf/Source/UserMessages.proto \
    Protobuf/Source/Types.proto \
    Protobuf/Source/Device.proto \
    Protobuf/Source/Search.proto \
    Protobuf/Source/UserProfiles.proto \
    Protobuf/Source/Feed.proto \
    Protobuf/Source/EntityTags.proto \
    Protobuf/Source/Server.proto \
    Protobuf/Source/UserEvents.proto \
    ; \
        ./Protobuf/make-proto ; \
        if [[ $$? != 0 ]]; then echo $?; exit 1; fi;

#        echo ">>> Building proto files."; \



# Clean --


clean: \
    ; \
        echo ">>> Cleaning..."; \
        $(rm) bin/*; \
        $(rm) pkg/*; \
        $(rm) Protobuf/Build; \
        $(rm) src/ApplePushService/ResourceData.go; \
        $(rm) src/BlitzMessage/*.pb.go;


# Testing --


gotest= \
    cd $1; \
    IFS=$$'\n'; \
    echo ">>>"; \
    echo ">>> Linting & testing `pwd`"; \
    echo ">>>"; \
    for dir in `find . -maxdepth 1 -type d  -not -name ".*"`; \
    do  cd "$$dir"; \
        echo ">>> -------------------------------------- $$dir"; \
        go vet || true; \
        go-nyet . || true; \
        go test; \
        cd - ; \
    done


test: \
    ; \
        cd $(makepath); $(call gotest, src); \
        cd $(makepath); $(call gotest, src/violent.blue/GoKit);


#  Deploy --


deploy: \
    linux \
    ; \
        echo ">>> Deploying to $$userhost." ; \
        \
        ssh $$userhost mkdir -p "$$installpath"/bin  "$$installpath"/database; \
        rsync -aP --force  --progress  \
            --exclude '.*' \
            --exclude '*.log' \
            --exclude 'log' \
            --exclude '*.Darwin*' \
            Staging/  \
            $$userhost:"$$installpath"/bin ;\
        if [[ $$? != 0 ]]; then exit 1; fi; \
        \
        rsync -aP --force  --progress \
            --exclude '.*' \
            Database/  \
            $$userhost:"$$installpath"/database; \
        if [[ $$? != 0 ]]; then exit 1; fi; \
        ssh $$userhost  bin/link-versions ; \
        if [[ $$? != 0 ]]; then exit 1; fi;


restart: \
    ; \
        echo ">>> Restarting servers" ; \
        ssh -T $$userhost sc restart all -f ;


all: clean linux compile deploy restart


server: linux compile deploy restart

