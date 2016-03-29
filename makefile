
# Build the server

SHELL       := /bin/bash
makepath    := $(shell pwd)
buildDate   := $(shell date | sed s.[[:space:]].-.g)
buildVersion = 0
buildLinux   = 0
buildDarwin  = 0

# Verbose options:

#verbose    := echo
#cp         := cp -av
#rm         := rm -Rfv

verbose     := true
cp          := cp -a
rm          := rm -Rf

.SILENT:

export userhost:=blitzhere@blitzhere.com
export installpath:=/home/blitzhere
export GOPATH:=$(makepath)
export GOROOT:=
export PATH:=$(makepath)/bin:$(PATH)


# Compile --


gobuildDarwin= \
    $(verbose) ">>> Darwin."; \
    env GOOS=darwin GOARCH=amd64 go install $(goflags) ; \
    if [[ $$? != 0 ]]; then exit 1; fi; \
    obj=$$(basename $1); \
    $(cp) bin/$$obj Staging/Versions/$$obj.Darwin; \


 gobuildLinux= \
    $(verbose) ">>> Linux."; \
    env GOOS=linux  GOARCH=amd64 go install $(goflags); \
    if [[ $$? != 0 ]]; then exit 1; fi; \
    obj=$$(basename $1); \
    $(cp) ../../bin/linux_amd64/$$obj ../../Staging/Versions/$$obj.Linux; \


gobuild= \
    cd $1; \
    echo ">>> Building `pwd`."; \
    $(eval goflags= -v -ldflags '-X violent.blue/GoKit/ServerUtil.compileTime=$(buildDate) -X violent.blue/GoKit/ServerUtil.compileVersion=$(buildVersion)') \
    $(call gobuildLinux, $1) \
    cd - >/dev/null; \


compile: \
    FORCE \
    updateversion \
    src/BlitzMessage/Conversations.pb.go \
    src/BlitzMessage/Device.pb.go \
    src/BlitzMessage/EntityTags.pb.go \
    src/BlitzMessage/Feed.pb.go \
    src/BlitzMessage/Friends.pb.go \
    src/BlitzMessage/Search.pb.go \
    src/BlitzMessage/Server.pb.go \
    src/BlitzMessage/Types.pb.go \
    src/BlitzMessage/UserEvents.pb.go \
    src/BlitzMessage/UserMessages.pb.go \
    src/BlitzMessage/UserProfiles.pb.go \
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
src/BlitzMessage/Conversations.pb.go \
src/BlitzMessage/Device.pb.go \
src/BlitzMessage/EntityTags.pb.go \
src/BlitzMessage/Feed.pb.go \
src/BlitzMessage/Friends.pb.go \
src/BlitzMessage/Search.pb.go \
src/BlitzMessage/Server.pb.go \
src/BlitzMessage/Types.pb.go \
src/BlitzMessage/UserEvents.pb.go \
src/BlitzMessage/UserMessages.pb.go \
src/BlitzMessage/UserProfiles.pb.go : \
    Protobuf/Source/Conversations.proto \
    Protobuf/Source/Device.proto \
    Protobuf/Source/EntityTags.proto \
    Protobuf/Source/Feed.proto \
    Protobuf/Source/Friends.proto \
    Protobuf/Source/Search.proto \
    Protobuf/Source/Server.proto \
    Protobuf/Source/Types.proto \
    Protobuf/Source/UserEvents.proto \
    Protobuf/Source/UserMessages.proto \
    Protobuf/Source/UserProfiles.proto \
    ; \
        ./Protobuf/make-proto ; \
        if [[ $$? != 0 ]]; then echo $?; exit 1; fi;


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
    FORCE \
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

