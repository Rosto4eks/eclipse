const delete_btns = document.querySelectorAll('.btn');
const search = document.getElementById("search")


async function getArticles(inp) {
    const response = await fetch(`/articles/search?value=${inp}`,{
        method : "GET"
    })
    const name = document.getElementById("username")
    const container = document.getElementById("container")
    const data = await response.json();
    container.innerHTML = ""
    data["articles"].forEach(article => {
        let elem = document.createElement("a");
        let deletebtn = ''
        if (name && article.NameAuthor == name.innerHTML) {
            deletebtn = `
                    <button class="btn" id="delete_btn{{ .ID }}" data-article-id="{{ .ID }}">
                        <svg viewBox="0 0 15 17.5" height="17.5" width="15" xmlns="http://www.w3.org/2000/svg" class="icon">
                            <path transform="translate(-2.5 -1.25)" d="M15,18.75H5A1.251,1.251,0,0,1,3.75,17.5V5H2.5V3.75h15V5H16.25V17.5A1.251,1.251,0,0,1,15,18.75ZM5,5V17.5H15V5Zm7.5,10H11.25V7.5H12.5V15ZM8.75,15H7.5V7.5H8.75V15ZM12.5,2.5h-5V1.25h5V2.5Z" id="Fill"></path>
                        </svg>
                    </button>
                `
        }
        elem.href = `/articles/${article.ID}`;
        elem.classList.add("album");
        elem.style.opacity = 0;
        elem.style.transform = "translate(0, 100px)"
        elem.innerHTML = `
            <img class="image preview" width="100%" src="/public/articles/${article.Date}-${article.Name}/preview.jpeg"></img>
            <div class="name">${article.Name}</div>
            <div class="author-date">
                <div class="author">By ${article.NameAuthor}</div>
                <div class="date">${article.Date}</div>
                <div class="theme">${article.Theme}</div>
                ${deletebtn}
            </div>
            <div class="text-wrap">
                <div class="text" id="text-${article.ID}"></div>
                <div class="more"></div>
            </div>
            <div class="read">Read more</div>
        `
        container.appendChild(elem);
         document.getElementById(`text-${article.ID}`).innerHTML = article.Text;
        setTimeout(() => {
            elem.style.opacity = 1;
            elem.style.transform = "translate(0,0)"
        }, 50);
    })
}

let id
search.addEventListener("input", e => {
    clearTimeout(id)
    id = setTimeout(() => {
        getArticles(e.target.value)
    }, 500);
})

delete_btns.forEach( delete_btn=> {
    delete_btn.addEventListener('click', () => {
        const articleId = delete_btn.dataset.articleId;
        console.log(articleId);
        deleteArticle(articleId);
    });
});

async function deleteArticle(articleId) {
    result = confirm("Are you sure? If you delete this article, you cannot recover it");
    if(!result){
        return;
    }
    const response = await fetch(`/articles/delete-article/${articleId}`,{
        method : "DELETE"
    })

    const data = await response.json();
    console.log(data);
    if(data.success){
        console.log(result);
        document.getElementById(`container${articleId}`).remove();
        document.getElementById(`delete_btn${articleId}`).remove();
    } else {
        alert(data.message);
    }
}


