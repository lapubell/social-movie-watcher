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

setInterval(() => {
    $("#playback").text(fancyTimeFormat(vid.currentTime) + " of " + fancyTimeFormat(vid.duration));
}, 1000);


var vm = new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        email: null, // Email address used for grabbing an avatar
        username: null, // Our username
        joined: false, // True if email and username have been filled in
        videoIsPlaying: false, // local state var so that if people come and go they can sync up to the global status
        audio: null,
        chatSounds: true // wether or not to make a sound each time a chat happens
    },

    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);

            // check original initialization here
            if (msg.initialMessage) {
                vid.currentTime = msg.video.timestamp;
                if (msg.video.isPlaying) {
                    vid.play();
                }
            }

            // check for various system messages here
            if (msg.video) {
                if (msg.newStatus === "play") {
                    if (!self.videoIsPlaying) {
                        console.log("was told to start playback")
                        vid.play();
                        self.videoIsPlaying = true
                    }
                }
                if (msg.newStatus === "pause") {
                    if (self.videoIsPlaying) {
                        console.log("was told to pause playback")
                        vid.pause();
                        console.log("updating video timestamp to match incoming value");
                        vid.currentTime = msg.video.timestamp;

                        self.videoIsPlaying = false
                    }
                }
            }

            self.chatContent += '<div class="message"><div class="chip">'
                    + '<img src="' + self.gravatarURL(msg.email) + '">' // Avatar
                    + msg.username
                + '</div>'
                + emojione.toImage(msg.message) + '</div>'; // Parse emojis

            setTimeout(() => {
                var element = document.getElementById('chat-messages');
                element.scrollTop = element.scrollHeight+100; // Auto scroll to the bottom

                if (self.chatSounds) {
                    document.getElementById("chatSound").volume = 0.2;
                    document.getElementById("chatSound").play();
                }
            }, 20);

        });
        this.audio = new Audio("/Pling-KevanGC-1485374730.mp3");
    },

    methods: {
        send: function () {
            if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        email: this.email,
                        username: this.username,
                        message: $('<p>').html(this.newMsg).text() // Strip out html
                    }
                ));
                this.newMsg = ''; // Reset newMsg
            }
        },

        sendPlay: function() {
            console.log("sending play signal to all")
            this.ws.send(
                JSON.stringify({
                    message: '!!video-play!!'
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
                Materialize.toast('You must enter an email', 2000);
                return
            }
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.email = $('<p>').html(this.email).text();
            this.username = $('<p>').html(this.username).text();
            this.joined = true;

            this.newMsg = "(just joined)"
            this.send()
        },

        gravatarURL: function(email) {
            return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
        }
    },

    watch: {
        joined() {
            if (this.joined) {
                $("#button-wrapper").show();
            } else {
                $("#button-wrapper").false();
            }
        }
    }
});
