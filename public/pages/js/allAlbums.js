const gallery = document.getElementById("gallery")
const end = document.getElementById("end")
let offset = 0, count = 5;

var observer = new IntersectionObserver((entries) => {
    if (entries[0].isIntersecting == true) {
        offset += 5;
        loadAlbums(offset, count)
    }
}, {threshold: [0]});
observer.observe(end);

async function loadAlbums(offset, count) {
    const response = await fetch(`/load-albums?offset=${offset}&count=${count}`,{
        method : "GET"
    })
    const data = await response.json();
    appendAlbums(data["albums"]);
}

function appendAlbums(data) {
    data.forEach(album => {
        let elem = document.createElement("a");
        elem.href = `/albums/${album.Id}`;
        elem.style.opacity = 0;
        elem.style.transform = "translate(0, 100px)"
        elem.innerHTML = `
            <div class="album">
                <div class="image" id="img${album.Id}">
                    <div class="image_text-wrapper">
                        <div class="image_text">Show ${ album.Images_count } images</div>
                    </div>
                </div>
                <div class="name">${ album.Name }</div>
                <div class="wr">
                    <div class="author">By ${ album.Author }</div>
                    <div class="date">Created ${ album.Date }</div>
                </div>
            </div>
        `
        gallery.appendChild(elem);
        console.log(album)
        document.getElementById(`img${album.Id}`).style.backgroundImage = `url(${escape(`/public/albums/${ album.Date }-${ album.Name }/0-compressed.jpeg`)})`
        setTimeout(() => {
            elem.style.opacity = 1;
            elem.style.transform = "translate(0,0)"
        }, 50);
    })
}
