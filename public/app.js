// good ol global functions, yay appending to the window object!
function fancyTimeFormat(duration) // https://stackoverflow.com/questions/3733227/javascript-seconds-to-minutes-and-seconds/11486026#11486026
{
    // Hours, minutes and seconds
    var hrs = ~~(duration / 3600);
    var mins = ~~((duration % 3600) / 60);
    var secs = ~~duration % 60;

    // Output like "1:01" or "4:03:59" or "123:03:59"
    var ret = "";

    if (hrs > 0) {
        ret += "" + hrs + ":" + (mins < 10 ? "0" : "");
    }

    ret += "" + mins + ":" + (secs < 10 ? "0" : "");
    ret += "" + secs;
    return ret;
}

// good ol global variables, yay appending to the window object!
var vid = document.getElementById("video-player");
var vidParent = vid.parentNode;
vid.addEventListener('ended', function() {
    vm.videoIsPlaying = false;
});
vid.addEventListener('play', function() {
    vm.videoIsPlaying = true;
    vm.drawerShown = false;
});
vid.addEventListener('pause', function() {
    vm.videoIsPlaying = false;
    vm.drawerShown = true;
});


var bPlay = document.getElementById("button-play");
bPlay.onclick = function() {
    vm.sendPlay();
}
var bPause = document.getElementById("button-pause");
bPause.onclick = function() {
    vm.sendPause();
}
var bAudio = document.getElementById("button-audio");
bAudio.onclick = function() {
    vid.muted = !vid.muted;
}
var bAudio = document.getElementById("button-back");
bAudio.onclick = function() {
    vm.sendBackSeconds();
}

var gniles = document.getElementById("gniles");
var newt = document.getElementById("newt");

setInterval(() => {
    let friendlyTime = fancyTimeFormat(vid.currentTime) + " of " + fancyTimeFormat(vid.duration);
    document.getElementById("playback").textContent = friendlyTime;
    vm.displayCurrentTime = friendlyTime;
}, 1000);

var vm = new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        displayCurrentTime: "", // human friendly readout of the current time
        newMsg: '', // Holds new messages to be sent to the server
        error: '',
        chatContent: '', // A running list of chat messages displayed on the screen
        chatMessages: [],
        email: null, // Email address used for grabbing an avatar
        username: null, // Our username
        oldUsername: null, // Our old username, for use when we switch usernames
        joined: false, // True if email and username have been filled in
        videoIsPlaying: false, // local state var so that if people come and go they can sync up to the global status
        audio: null,
        chatSounds: true, // wether or not to make a sound each time a chat happens
        rooms: [],
        creatingRoom: false,
        chosenVideo: "",
        availableVideos: [],
        lurkers: [],
        drawerShown: false,
        showChangeName: false,
        windowSize: 0,
        appFeatures: "both", // both|video|chat
        autoJoin: false,
        autoJumpToBottom: false,
    },

    created: function() {
        this.email = localStorage.getItem('email');
        this.username = localStorage.getItem('username');
        this.audio = new Audio("/Pling-KevanGC-1485374730.mp3");

        this.wsInitialize();

        this.windowSize = screen.width;

        window.addEventListener("resize", (e) => {
            this.windowSize = e.currentTarget.innerWidth;
        });

        this.autoJoin = localStorage.getItem("auto-join") === "true";
        this.chatSounds = localStorage.getItem("chat-sounds") === "true";
        this.autoJumpToBottom = localStorage.getItem("jump-to-bottom") === "true";
    },

    methods: {
        wsInitialize: function() {
            // reset it to null in case it's something weird
            this.ws = null;
            this.error = "";

            // bind new WS connection
            var self = this;
            var proto = window.parent.location.protocol === "http:" ? "ws": "wss";
            this.ws = new WebSocket(proto + '://' + window.location.host + '/ws');
            this.ws.addEventListener('message', function(e) {
                var msg = JSON.parse(e.data);

                // check original initialization here
                if (msg.initialMessage) {
                    vid.setAttribute("src", msg.video.video)
                    vid.currentTime = msg.video.timestamp;
                    if (msg.video.isPlaying) {
                        vid.play();
                        self.videoIsPlaying = true;
                    }

                    return
                }

                // check for various system messages here
                if (msg.newStatus) {
                    if (msg.newStatus === "play") {
                        if (!self.videoIsPlaying) {
                            vid.play();
                        }
                        vid.currentTime = msg.video.timestamp;
                        self.videoIsPlaying = true
                    }
                    if (msg.newStatus === "pause") {
                        if (self.videoIsPlaying) {
                            vid.pause();
                        }
                        vid.currentTime = msg.video.timestamp;
                        self.videoIsPlaying = false
                    }
                    if (msg.newStatus === "change") {
                        vid.pause()
                        vid.setAttribute("src", msg.changeTo)
                        vid.play()
                        self.videoIsPlaying = true
                    }
                    if (msg.newStatus === "list") {
                        self.availableVideos = msg.videos
                    }
                    if (msg.newStatus === "peeps") {
                        self.lurkers = msg.peeps
                    }

                    return
                }

                var avatar = self.gravatarURL(msg.email);
                self.chatMessages.push({
                    avatarURL: avatar,
                    username: msg.username,
                    message: msg.message
                });

                if (msg.message.toLowerCase().substr(0, 5) === "it me") {
                    let duration = msg.message.substr(5).length;
                    if (duration < 1.5) {
                        duration = 1.5;
                    }
                    if (duration > 10) {
                        duration = 10;
                    }
                    self.itMe(duration);
                }

                if (msg.message.toLowerCase().substr(0,6) === "gniles") {
                    if (!gniles.classList.contains("peekaboo")) {
                        gniles.classList.add("peekaboo");

                        setTimeout(() => {
                            gniles.classList.remove("peekaboo");
                        }, 2000);
                    }
                }

                if (msg.message.toLowerCase().substr(0,8) === "mewstley" || msg.message.toLowerCase().substr(0,7) === "mewstly") {
                    if (!newt.classList.contains("peekaboo")) {
                        newt.classList.add("peekaboo");

                        setTimeout(() => {
                            newt.classList.remove("peekaboo");
                        }, 2000);
                    }
                }

                if (self.autoJumpToBottom) {
                    setTimeout(() => {
                        var element = document.getElementById('chat-messages');
                        if (!element) {
                            return;
                        }
                        element.scrollTop = element.scrollHeight+100; // Auto scroll to the bottom

                    }, 250);
                }

                if (self.chatSounds) {
                    setTimeout(() => {
                        document.getElementById("chatSound").volume = 0.15;
                        document.getElementById("chatSound").play();
                    }, 100);
                }
            });

            this.ws.onclose = function() {
                self.error = "No websocket"
            };

            // check auto join
            this.ws.onopen = function() {
                if (localStorage.getItem("auto-join") === "true" && self.joined === false) {
                    self.join();
                }
            }

            // scroll to bottom of chat, helpful if reconnecting
            setTimeout(() => {
                var element = document.getElementById('chat-messages');
                if (!element) {
                    return;
                }
                element.scrollTop = element.scrollHeight+100; // Auto scroll to the bottom

            }, 250);
        },
        creatingARoom: function() {
            this.creatingRoom = true;
            this.ws.send(
                JSON.stringify({
                    message: "!!available-videos!!"
                })
            );
        },
        createRoom: function() {
            this.we.send(
                JSON.stringify({
                    message: "!!create-room!!",
                    username: this.chosenVideo
                })
            );
        },
        send: function () {
            if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        email: this.email,
                        username: this.username,
                        message: this.newMsg
                    }
                ));
                this.newMsg = ''; // Reset newMsg
            }
        },
        giphySend: function () {
            if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        email: this.email,
                        username: this.username,
                        message: ":giphy " + this.newMsg
                    }
                ));
                this.newMsg = ''; // Reset newMsg
            }
        },

        sendPlay: function() {
            console.log("sending play signal to all")
            this.ws.send(
                JSON.stringify({
                    message: '!!video-play!!',
                    timestamp: Math.round(vid.currentTime)
                }
            ));
            setTimeout(() => {
                this.ws.send(
                    JSON.stringify({
                        email: this.email,
                        username: this.username,
                        message: '(played the video at ' + fancyTimeFormat(vid.currentTime) + ')'
                    }
                ));
            }, 100);
        },

        sendPause: function() {
            this.ws.send(
                JSON.stringify({
                    message: '!!video-pause!!',
                    timestamp: Math.round(vid.currentTime)
                }
            ));
            setTimeout(() => {
                this.ws.send(
                    JSON.stringify({
                        email: this.email,
                        username: this.username,
                        message: '(paused the video at ' + fancyTimeFormat(vid.currentTime) + ')'
                    }
                ));
            }, 100);
        },

        sendBackSeconds: function () {
            vid.pause();
            var newTime = Math.round(vid.currentTime) - 15;
            if (newTime < 0) {
                vid.currentTime = 0;
            } else {
                vid.currentTime = newTime;
            }
            this.sendPause();
            setTimeout(() => {
                this.sendPlay();
            }, 500);
        },

        join: function () {
            if (!this.email) {
                alert('You must enter an email');
                return
            }
            if (!this.username) {
                alert('You must choose a username');
                return
            }
            this.email = this.email;
            this.username = this.username;
            this.oldUsername = this.username;
            this.joined = true;

            this.newMsg = "(just joined)";
            this.send();

            if (this.videoIsPlaying) {
                vid.play();
            }

            localStorage.setItem('email', this.email);
            localStorage.setItem('username', this.username);

            this.newMsg = "!!available-videos!!";
            this.send();
        },

        updateName: function() {
            localStorage.setItem('email', this.email);
            localStorage.setItem('username', this.username);

            this.newMsg = "(\"" + this.oldUsername + "\" is now \"" + this.username + "\")";
            this.send();

            this.newMsg = `!!changeName:NEW::${this.username}!!`;
            this.send();

            this.oldUsername = this.username;

            this.showChangeName = false;
        },

        gravatarURL: function(email) {
            return 'https://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
        },

        showEmoji: function() {
            var trigger = document.getElementById("trigger");
            window.picker.togglePicker(trigger);
        },

        switchMovie(e) {
            console.log("switch!", e.target.value)

            this.newMsg = `!!change:${e.target.value}!!`;
            this.send();
        },
        isCurrentVid(video) {
            if (!vid) {
                return false;
            }
            return "/video/"+video === vid.getAttribute("src")
        },
        itMe(seconds) {
            var duration = seconds * 1000;
            var animationEnd = Date.now() + duration;
            var defaults = { startVelocity: 30, spread: 360, ticks: 60, zIndex: 0 };

            function randomInRange(min, max) {
                return Math.random() * (max - min) + min;
            }
            var canvasOverlay = document.getElementsByTagName("canvas")[0];
            canvasOverlay.style.display = "block";

            var interval = setInterval(function() {
                var timeLeft = animationEnd - Date.now();

                if (timeLeft <= 0) {
                    canvasOverlay.style.display = "none";
                    return clearInterval(interval);
                }

                var particleCount = 50 * (timeLeft / duration);
                // since particles fall down, start a bit higher than random
                confetti(Object.assign({}, defaults, { particleCount, origin: { x: randomInRange(0.1, 0.3), y: Math.random() - 0.2 } }));
                confetti(Object.assign({}, defaults, { particleCount, origin: { x: randomInRange(0.7, 0.9), y: Math.random() - 0.2 } }));
            }, 250);
        },
        removeVideoPlayer() {
            console.log("remove video player!");
            vid.muted = true;
        },
        enableVideoPlayer() {
            console.log("enable video player!");
        }
    },

    watch: {
        autoJoin() {
            if (this.autoJoin) {
                localStorage.setItem("auto-join", "true");
                return;
            }
            localStorage.setItem("auto-join", "false");
        },
        autoJumpToBottom() {
            if (this.autoJumpToBottom) {
                localStorage.setItem("jump-to-bottom", "true");
                return;
            }
            localStorage.setItem("jump-to-bottom", "false");
        },
        chatSounds() {
            if (this.chatSounds) {
                localStorage.setItem("chat-sounds", "true");
                return;
            }
            localStorage.setItem("chat-sounds", "false");
        },
        joined() {
            var btnWrapper = document.getElementById("button-wrapper");

            if (this.joined) {
                btnWrapper.classList.remove("hidden");
            } else {
                btnWrapper.classList.add("hidden");
            }
        },
        appFeatures() {
            if (this.appFeatures === "chat") {
                document.getElementById("video-app").classList.add("disabled");
                document.getElementById("app").classList.remove("disabled");
                this.removeVideoPlayer();
                return;
            }

            if (this.appFeatures === "video") {
                document.getElementById("video-app").classList.remove("disabled");
                document.getElementById("app").classList.add("disabled");
                this.enableVideoPlayer();

                var tag = document.createElement("h1");
                var text = document.createTextNode("YOU'LL BE SORRY");
                tag.appendChild(text);
                var element = document.querySelector("body");
                element.appendChild(tag);

                setTimeout(() => {
                    element.removeChild(tag);
                }, 1000);

                return;
            }

            if (this.appFeatures === "both") {
                document.getElementById("video-app").classList.remove("disabled");
                document.getElementById("app").classList.remove("disabled");
                this.enableVideoPlayer();
                return;
            }

            // unreachable code
            console.log("invalid value for appFeatures");
        },
        windowSize() {
            if (this.windowSize < 500) {
                this.appFeatures = "chat";
            }
        }
    }
});
