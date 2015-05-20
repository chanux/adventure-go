# Go Adventure

Trying to implement [python adventure](https://github.com/thinkcube/python-adventure) with Go.

Build as follows

    go build -o advgo adv.go

Run as follows

    ./advgo

Check it as follows

    wget -qO - localhost:9000/adventure

    curl localhost:9000/adventure

    http --stream localhost:9000/adventure


Static build

    CGO_ENABLED=0 go build -o advent -a -ldflags '-s' adv.go

I build it with static linking to put it in a docker image.

Build docker image as follows

    docker build -t chanux/advgo

Run docker container as follows (This works even if you din't build. I pushed 
[my image](https://registry.hub.docker.com/u/chanux/advgo/) to docker hub)

    docker run -it --rm -p 9000:9000 chanux/advgo

You can check it as same as before.
