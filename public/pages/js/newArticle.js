function get_selected_text() {
	if (window.getSelection()) {
		var range = window.getSelection().getRangeAt(0);
    var selectionContents = range.extractContents();
    var div = document.createElement("div");
    div.style.color = "red";
    div.style.display = "inline-block";
    div.appendChild(selectionContents);
    range.insertNode(div);
	}
}
