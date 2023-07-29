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
   let text = document.getElementById(`text${commentId}`);

   item.innerHTML = (`
    <div class="comment_buttons" id="comment_buttons">
        <button class="applyBtn" id="applyBtn">Apply</button>
        <button class="cancelBtn" id="cancelBtn">Cancel</button>
    </div>`);
    text.contentEditable = true;
    if(document.getElementById(`comment_buttons`)===null) {
        comment.append(item);
    }

   const cancel = document.getElementById(`cancelBtn`);
   const apply = document.getElementById(`applyBtn`);

    apply.addEventListener('click', () => {
        changeComment(new String(commentId), text.innerText);
        text.contentEditable = false;
        document.getElementById(`comment_buttons`).remove();
    });
    cancel.addEventListener('click', () =>{
        text.contentEditable = false;
        document.getElementById(`comment_buttons`).remove();
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
    let item = document.createElement(`change`);
    item.innerHTML = (`<div class="change_article" id="change_article">
    <div class="article_buttons" id="buttons">
        <button class="applyArticleBtn" id="applyArticleBtn">Apply</button>
        <button class="cancelArticleBtn" id="cancelArticleBtn">Cancel</button>
    </div>
    </div>`);
    articletext.contentEditable = true;
    if(document.getElementById(`change_article`)===null) {
        article.append(item);
    }

    const applyBtn = document.getElementById(`applyArticleBtn`);
    const cancelBtn = document.getElementById(`cancelArticleBtn`);
    applyBtn.addEventListener(`click`, () => {
        changeArticle(articleId, articletext.innerHTML);
        articletext.contentEditable = false;
        document.getElementById(`change_article`).remove();
    });
    cancelBtn.addEventListener(`click`, () =>{
        articletext.contentEditable = false;
        document.getElementById(`change_article`).remove();
    });
});

async function changeArticle(articleId, newText) {
    const response = await fetch(`/articles/${articleId}/change_article`, {
        method : "PATCH",
        body : JSON.stringify({articleId: articleId, text: newText})
    })

    const data = await response.json();
    if(data.success){
        item.remove();
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
