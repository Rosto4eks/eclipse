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
   });
});

async function changeComment(commentId, newText, authorName) {
    const response = await fetch('',{
        method : "PUT",
        body : JSON.stringify({commentId: commentId, text: newText, author: authorName})
    });

    const data = response.json();
    console.log(data);

    if(data.success){

    } else {

    }
}