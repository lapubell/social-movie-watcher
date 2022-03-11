import { EmojiButton } from '@joeattardi/emoji-button';

// globally enabled confetti
const confetti = require('canvas-confetti');
let myCanvas = document.createElement('canvas');
myCanvas.style.position = "fixed";
myCanvas.style.zIndex = 100;
myCanvas.style.top = "100%";
myCanvas.style.left = 0;
myCanvas.style.width = "100%";
myCanvas.style.height = "100%";
myCanvas.id = "confettiCanvas";
document.getElementById("content").appendChild(myCanvas);
let myConfetti = confetti.create(myCanvas, {
  resize: true,
  useWorker: true
});

window.confetti = myConfetti;

window.celebrate = function (seconds) {
    console.log(this);
    var duration = seconds * 1000;
    var animationEnd = Date.now() + duration;
    var defaults = { startVelocity: 30, spread: 360, ticks: 60, zIndex: 0 };

    function randomInRange(min, max) {
        return Math.random() * (max - min) + min;
    }
    var canvasOverlay = document.getElementById("confettiCanvas");
    canvasOverlay.style.display = "block";
    canvasOverlay.style.top = 0;

    var interval = setInterval(function() {
        var timeLeft = animationEnd - Date.now();

        if (timeLeft <= 0) {
            canvasOverlay.style.display = "none";
            canvasOverlay.style.top = "100%";
            return clearInterval(interval);
        }

        var particleCount = 50 * (timeLeft / duration);
        // since particles fall down, start a bit higher than random
        window.confetti(Object.assign({}, defaults, { particleCount, origin: { x: randomInRange(0.1, 0.3), y: Math.random() - 0.2 } }));
        window.confetti(Object.assign({}, defaults, { particleCount, origin: { x: randomInRange(0.7, 0.9), y: Math.random() - 0.2 } }));
    }, 250);
}


window.picker = new EmojiButton();

window.picker.on('emoji', selection => {
  // handle the selected emoji here
  if (window.vm.newMsg.length) {
    window.vm.newMsg += " ";
  }
  window.vm.newMsg += selection.emoji;
});
