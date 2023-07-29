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

window.addEventListener(`scroll`,trackScroll);
let element = document.getElementById(`container_page`);
let another = document.getElementById(`introduction`);
let inIntroduction = true;

function trackScroll(){
    const YPos = another.scrollTop;
    console.log(this.scrollY, inIntroduction);
    const height = document.documentElement.clientHeight;
    // if(this.scrollY < 700 && !inIntroduction){
    //     another.scrollIntoView({behavior: "smooth"});
    //     inIntroduction = true;
    //     return;
    // }

    // if(this.scrollY > 125 && this.scrollY < 350 && inIntroduction){
    //     element.scrollIntoView({behavior : "smooth"});
    //     inIntroduction = false;
    // }
}