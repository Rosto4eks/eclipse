const gifList = [
  "/public/images/moon-phases-moon.gif",
  "/public/images/7gkY.gif",
  "/public/images/1389643383_1137314990.gif",
  "/public/images/two.gif",
  "/public/images/three.gif"
];

const gifElement = document.getElementById(`gif`);
let curIndex = 0;

function changeGif(){
    gifElement.style.backgroundImage = `url(${gifList[curIndex]})`;
    curIndex = (curIndex + 1) % gifList.length;
}

setInterval(changeGif, 2500);