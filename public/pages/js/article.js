const deleteBtns = document.querySelectorAll('.deleteBtn');
const article = document.getElementById(`article`);

deleteBtns.forEach( deleteBtn=>{
    deleteBtn.addEventListener('click', () => {
        const articleId = article.dataset.articleId;
        const commentId = deleteBtn.dataset.commentId;
        deleteComment(articleId, commentId);
    });
});

async function deleteComment(articleId, commentId) {
    result = confirm("Are you sure? If you delete this comment, you cannot recover it");
    if(!result) {
       return;
    }
    const response = await fetch(`/articles/${articleId}/delete-comment/${new String(commentId)}`,{
        method : "DELETE"
    })

    const data = await response.json();
    if(data.success){
        document.getElementById(`comments${commentId}`).remove();
    } else {
        alert(data.message);
    }
}

const changeBtns = document.querySelectorAll(`.changeBtn`);

changeBtns.forEach(changeBtn => {
   changeBtn.addEventListener(`click`, () => {
       const commentId = changeBtn.dataset.commentId;
       replaceCommentWithTextarea(commentId);
   });
});

function replaceCommentWithTextarea(commentId) {
   let comment = document.getElementById(`comments${commentId}`);
   let item = document.createElement(`div`);
   let text = document.getElementById(`text${commentId}`).textContent;

   item.innerHTML = (`
    <div class="change_comment_body">
    <textarea name="input_change" id="change_text">${text}</textarea>
    <div class="buttons">
        <button class="applyBtn" id="applyBtn">Apply</button>
        <button class="cancelBtn" id="cancelBtn">Cancel</button>
    </div>
    </div>`);

   let parentNode = comment.parentNode;
   parentNode.replaceChild(item, comment);

    let com = document.getElementById(`change_text`);
    let height = com.scrollHeight;
    com.style.height = `${height}px`;
    textareaAutoResize();

   const cancel = document.getElementById(`cancelBtn`);
   const apply = document.getElementById(`applyBtn`);

   apply.addEventListener('click', () => {
        let newText = document.getElementById(`change_text`).value;
        let newDate = changeComment(commentId, newText);
        let commentTextElement = comment.querySelector(`#text${commentId}`);
        let commentHeaderElement = comment.querySelector(`.comment_author`);
        commentTextElement.textContent = newText;
        parentNode.replaceChild(comment, item);
    });
    cancel.addEventListener('click', () =>{
        parentNode.replaceChild(comment,item);
    });
}

async function changeComment(commentId, newText) {
    const response = await fetch('/articles/:article_id/change_comment',{
        method : "PATCH",
        body : JSON.stringify({commentId: commentId, text: newText})
    });

    const data = await response.json();
    console.log(data);

    if(data.success){
        console.log("success");
    } else {
        alert(data.message);
    }
    return data;
}

const changeArticleBtn = document.getElementById(`changeArticleBtn`);
const articletext = document.getElementById(`text`);

changeArticleBtn.addEventListener(`click`, () => {
    const articleId = changeArticleBtn.dataset.articleId;
    let item = document.createElement(`div`);
    item.innerHTML = (` <div class="change_article">
    <div class="buttons">
        <button class="applyArticleBtn" id="applyArticleBtn">Apply</button>
        <button class="cancelArticleBtn" id="cancelArticleBtn">Cancel</button>
    </div>
    </div>`);
    articletext.contentEditable = true
    article.append(item);

    const applyBtn = document.getElementById(`applyArticleBtn`);
    applyBtn.addEventListener(`click`, () => {
        changeArticle(articleId, articletext.innerHTML);
        article.contentEditable = false;
    });
    const cancelBtn = document.getElementById(`cancelArticleBtn`);
    cancelBtn.addEventListener(`click`, () =>{
       article.contentEditable = false;
    });
});

async function changeArticle(articleId, newText) {
    const response = await fetch(`/articles/${articleId}/change_article`, {
        method : "PATCH",
        body : JSON.stringify({articleId: articleId, text: newText})
    })

    const data = await response.json();
    if(data.success){

    } else {
        alert(data.message);
    }
}

function textareaAutoResize() {
    textAreas = document.querySelectorAll(`textarea`);
    textAreas.forEach( textArea =>{
        textArea.addEventListener(`keydown`, e => {
            textArea.style.height = `auto`;
            let height = e.target.scrollHeight;
            textArea.style.height = `${height}px`;
        });
    });
}

textareaAutoResize();
