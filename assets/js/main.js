import { EmojiButton } from '@joeattardi/emoji-button';

window.picker = new EmojiButton();

window.picker.on('emoji', selection => {
  // handle the selected emoji here
  if (window.vm.newMsg.length) {
    window.vm.newMsg += " ";
  }
  window.vm.newMsg += selection.emoji;
});
