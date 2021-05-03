import { EmojiButton } from '@joeattardi/emoji-button';
const confetti = require('canvas-confetti');

let myCanvas = document.createElement('canvas');
document.getElementById("video-app").appendChild(myCanvas);
let myConfetti = confetti.create(myCanvas, {
  resize: true,
  useWorker: true
});
window.confetti = myConfetti;

window.picker = new EmojiButton();

window.picker.on('emoji', selection => {
  // handle the selected emoji here
  if (window.vm.newMsg.length) {
    window.vm.newMsg += " ";
  }
  window.vm.newMsg += selection.emoji;
});
