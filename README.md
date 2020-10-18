# Go Chat

This is a social video watching program, forked from a simple chat web app written in Go. Please checkout [this rad demo](https://github.com/scotch-io/go-realtime-chat) to see the origins.

## First, go encode a video

For legal purposes, this should be a video that you own or have the rights to share over the interwebs.

Save that video as `video/video.mp4`.

## Go get a server

Go get a server from Digi Ocean, Vultr, Linode, whoever. If it's a ubuntu server, you won't have any firewall issues to deal with. It's nice and insecure by default.

## Upload everything.

Once it's on the internet ssh into it and run the binary in the `src` directory.

```
scp ./* root@NEWSERVERIP:/video-watcher/
cd /video-watcher/src/
./video-time
```

Then point your browser to http://IPADDRESS:8081

### References

Chat blip came from here:
http://soundbible.com/1645-Pling.html
