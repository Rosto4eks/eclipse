var element;
var size = document.getElementById("Size")
var b = document.getElementById("Bold")
var i = document.getElementById("Italic")

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

async function submit() {
  theme = document.getElementById("theme")
  title = document.getElementById("title")
  text = document.getElementById("text")
  data = {
    theme: theme.value,
    title: title.value,
    text: text.innerHTML
  }
  const response = await fetch("/articles/new", {
    method: "POST", 
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  const res = await response.json();
  console.log(res);
}
