const deleteBtns = document.querySelectorAll('.deleteBtn');
const article = document.getElementById(`article`);

deleteBtns.forEach( deleteBtn=>{
    deleteBtn.addEventListener('click', () => {
        const articleId = article.dataset.articleId;
        const commentId = deleteBtn.dataset.commentId;
        console.log("commentId",commentId, "articleId",articleId);
        deleteComment(articleId, commentId);
    });
});

async function deleteComment(articleId, commentId) {
    result = confirm("Are you sure? If you delete this comment, you cannot recover it");
    if(!result) {
       return;
    }
    const response = await fetch(`/articles/${articleId}/delete-comment/${commentId}`,{
        method : "DELETE"
    })

    const data = await response.json();
    console.log(data);
    if(data.success){
        console.log(commentId)
        document.getElementById(`comments${commentId}`).remove();
    } else {
        alert(data.message);
    }
    console.log("mama", data.success, "papa", data.message);
}

const changeBtns = document.querySelectorAll(`.changeBtn`);

changeBtns.forEach(changeBtn => {
   changeBtn.addEventListener(`click`, () => {
       const commentId = changeBtn.dataset.commentId;
       console.log(commentId);
       replaceCommentWithTextarea(commentId);
   });
});

function replaceCommentWithTextarea(commentId) {
   let comment = document.getElementById(`comments${commentId}`);
   console.log(comment);
   let item = document.createElement(`div`);
   let text = document.getElementById(`text${commentId}`).textContent;
   console.log(text);

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

   const cancel = document.getElementById(`cancelBtn`);
   const apply = document.getElementById(`applyBtn`);

   apply.addEventListener('click', () => {
        console.log("apply clicked");
        let newText = document.getElementById(`change_text`).value;
        let newDate = changeComment(commentId, newText);
        let commentTextElement = comment.querySelector(`#text${commentId}`);
        let commentHeaderElement = comment.querySelector(`.comment_author`);
        commentTextElement.textContent = newText;
        let options = {
            year: 'numeric',
            month: 'long',
            day: 'numeric',
            hour: 'numeric',
            minute: 'numeric',
        }
        //commentHeaderElement.textContent = commentHeaderElement.textContent + ` Edited ${new Date().toLocaleString(options)}`;
        parentNode.replaceChild(comment, item);
    });
    cancel.addEventListener('click', () =>{
        console.log("cancel clicked");
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
        alert(data.message)
    }
    return data
}