const sendBtn = document.getElementById(`sendBtn`);
const comment = document.getElementById(`comment_text`);
const all = document.getElementById(`allComments`);

sendBtn.addEventListener('click', () =>{
    const articleId = article.dataset.articleId;
    const text = comment.value;
    const author = comment.dataset.authorName;
    if(!text){
        alert('Field must contain at least 1 character');
    } else {
        const response = postComment(articleId, text, author);
        comment.value = '';
        console.log(comment.value);
    }
});

function commentToHTML(articleID, textComment, author, date, comment_id){
    const item = document.createElement('li');

    item.innerHTML = (`<div class="comments" id="comments${comment_id}">
            <div class="comment_author">Written by ${author} ${date}</div>
            <div class="message">
                <div class="comment_text">${textComment}</div>
            </div>
            <button class="deleteBtn" data-comment-id="${comment_id}">
                <svg viewBox="0 0 15 17.5" height="17.5" width="15" xmlns="http://www.w3.org/2000/svg" class="icon">
                    <path transform="translate(-2.5 -1.25)"
                          d="M15,18.75H5A1.251,1.251,0,0,1,3.75,17.5V5H2.5V3.75h15V5H16.25V17.5A1.251,1.251,0,0,1,15,18.75ZM5,5V17.5H15V5Zm7.5,10H11.25V7.5H12.5V15ZM8.75,15H7.5V7.5H8.75V15ZM12.5,2.5h-5V1.25h5V2.5Z"
                          id="Fill"></path>
                </svg>
            </button>
        </div>`);
    all.appendChild(item);
}

async function postComment(articleId, textComment, author) {
    const response = await fetch(`/articles/${articleId}/new`,{
        method : "POST",
        body : JSON.stringify({article: articleId, text: textComment, author: author})
    })

    const data = await response.json();
    console.log(data);

    if(data.success){
        commentToHTML(articleId, textComment, author, data.date, data.comment_id);
    } else {
        alert(data.message)
    }
}

