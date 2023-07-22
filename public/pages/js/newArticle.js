var element;
var size = document.getElementById("Size")
var b = document.getElementById("Bold")
var i = document.getElementById("Italic")
var title = document.querySelector(".title")

// TOOLS
title.addEventListener("dblclick", (e) => {
  if (title.style.color == "white") {
    title.style.color = "black";
  }
  else {
    title.style.color = "white";
  }
})

function bold() {
  if (element.style["font-weight"] == "700") {
    element.style["font-weight"] = "400"
  }
  else {
    element.style["font-weight"] = "700";
  }
  changeBg(b, element.style["font-weight"] == "700")
} 

function italic() {
  if (element.style["font-style"] == "italic") {
    element.style["font-style"] = "normal"
  }
  else {
    element.style["font-style"] = "italic";
  }
  changeBg(i, element.style["font-style"] == "italic")
}

document.getElementById("Size").addEventListener("input", (e) => {
  element.style["font-size"] = e.target.value + "px"
})
 
function select() {
  element = document.querySelectorAll(':hover')[4];
  if (element.style["font-size"]) {
    size.value = element.style["font-size"].substring(0, element.style["font-size"].length - 2)
  }
  else {
    size.value = "16"
  }
  changeBg(b, element.style["font-weight"] == "700")
  changeBg(i, element.style["font-style"] == "italic")
}

function changeBg(e, condition) {
  if (condition) {
    e.style["background-color"] = "#43ee65";
    e.style["color"] = "white";
  }
  else {
    e.style["background-color"] = "#efefef";
    e.style["color"] = "black";
  }
}

// ----------------------------------------------

const twr = document.querySelector(".title-wrap")
let prevImg;
// preview
preview.addEventListener("change", (e) => {
  prevImg = preview.files[0];
  if (twr.childElementCount > 1) {
    twr.removeChild(twr.lastChild)
  }
  let prev = document.createElement("img")
  prev.classList.add("preview")
  prev.style.width = "100%"
  prev.src = `${URL.createObjectURL(prevImg)}`
  twr.appendChild(prev)
  document.querySelector(".title").style.top = "50%";
})

let images = new Map();
let tempImages = new Map();
let index = 0;
const upload = document.querySelector("#upload");
const output = document.querySelector("#output");

upload.addEventListener("change", (e) => {
  const files = upload.files;
  tempImages.clear();
  Array.from(files).forEach(element => {
    images.set(index, element);
    tempImages.set(index, element);
    index++;
  });
  display();
})

function display() {
  let imagesHTML = ""
  tempImages.forEach((img, i, m) => {
    imagesHTML += `<div class="img" draggable ondragstart="drag(event)"><img width="100%" id="${i}" class="image" src="${URL.createObjectURL(img)}"></div>`
  })
  output.innerHTML = imagesHTML

  var imgs = document.querySelectorAll(".image")
  for (let i = 0; i < imgs.length; i++) {
    imgs[i].addEventListener("dblclick", (e) => {
      e.target.remove();
      images.delete(Number(e.target.id));
      tempImages.delete(Number(e.target.id));
    })
  }
}

function allowDrop(ev) {
  ev.preventDefault();
}

function drag(ev) {
  ev.dataTransfer.setData("text", ev.target.id);
}

function drop(ev) {
  ev.preventDefault();
  var data = ev.dataTransfer.getData("text");
  ev.target.appendChild(document.getElementById(data));
}

async function submit() {
  theme = document.getElementById("theme");
  title = document.getElementById("title");
  text = document.getElementById("text");
  
  let fd = new FormData();
  fd.set("title", title.value);
  fd.set("theme", theme.value);
  fd.set("color", title.style.color)
  let i = 0;
  images.forEach((elem, index) => {
    fd.append("images", elem);
    document.getElementById(`${index}`).src = `img-src-${i}`;
    i++;
  })
  fd.append("images", prevImg);

  fd.set("text", text.innerHTML);
  fetch("/articles/new", {
    method: "POST", 
    body: fd,
  }).then(r => {
    window.location.replace("/articles")
  })
}

